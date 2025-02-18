package config

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	dt "github.com/docker/docker/api/types"
	dtc "github.com/docker/docker/api/types/container"
	docker "github.com/docker/docker/client"

	"github.com/nodeset-org/hyperdrive-daemon/shared/config"
	"github.com/nodeset-org/hyperdrive-ethereum/adapter/config/utils"
	"github.com/nodeset-org/hyperdrive-ethereum/adapter/config/utils/terminal"
	"github.com/urfave/cli/v2"
)

const dockerImageRegex string = ".*/(?P<image>.*):.*"

// Start the Hyperdrive service
func startService(c *cli.Context, ignoreConfigSuggestion bool) error {
	cfgMgr, err := NewAdapterConfigManager(c)
	if err != nil {
		return fmt.Errorf("Error creating config manager: %w", err)
	}
	cfg := cfgMgr.(*AdapterConfigManager).AdapterConfig

	// TODO: Implement isNew
	if cfg.IsNew {
		return fmt.Errorf("No configuration detected. Please run `hyperdrive service config` to set up Hyperdrive before running it.")
	}

	// Validate config
	// errors := cfg.Validate()
	// if len(errors) > 0 {
	// 	fmt.Printf("%sYour configuration encountered errors. You must correct the following by changing the settings with `hyperdrive service config` in order to start Hyperdrive:\n\n", terminal.ColorRed)
	// 	for _, err := range errors {
	// 		fmt.Printf("%s\n\n", err)
	// 	}
	// 	fmt.Println(terminal.ColorReset)
	// 	return nil
	// }

	// Perform anti-slashing safety check
	if !c.Bool(ignoreSlashTimerFlag.Name) {
		firstRun, err := checkForValidatorChange(cfg)
		if err != nil {
			fmt.Printf("%sWARNING: couldn't verify validator client restart safety:\n\t%s\n", terminal.ColorYellow, err.Error())
			fmt.Println("If changing clients, wait 15 minutes before starting.")
			if !utils.Confirm(fmt.Sprintf("Press y to acknowledge and start Hyperdrive:%s", terminal.ColorReset)) {
				fmt.Println("Cancelled.")
				return nil
			}
		} else if firstRun {
			if !utils.Confirm(fmt.Sprintf("Press y to acknowledge and start Hyperdrive:%s", terminal.ColorReset)) {
				fmt.Println("Cancelled.")
				return nil
			}
		}
	}

	// Start only the Beacon Node and Execution Client
	fmt.Println("Starting Beacon Node and Execution Client...")
	err = hd.StartService([]string{
		"beacon-node",
		"execution-client",
	})
	if err != nil {
		return fmt.Errorf("error starting beacon and execution client: %w", err)
	}

	fmt.Println("Beacon Node and Execution Client started successfully.")
	return nil
}

func checkForValidatorChange(cfg *HyperdriveEthereumConfigSettings) (bool, error) {
	// Get all of the VCs belonging to the project
	prefix := cfg.Hyperdrive.ProjectName.Value
	vcs, err := cfg.GetValidatorContainers(prefix + "_") // Used to be hd
	if err != nil {
		return false, fmt.Errorf("error getting validator client containers: %w", err)
	}

	// Break if there aren't any
	if len(vcs) == 0 {
		return true, nil
	}

	/*
		// TODO: DEBUG
			fmt.Println("Found the following Validator Clients:")
			for _, vc := range vcs {
				fmt.Println(vc)
			}
			fmt.Println()
	*/

	// Get the map of VCs to their new tags in the config
	newTagMap, err := getVcContainerTagParamMap(cfg, vcs)
	if err != nil {
		return false, err
	}

	// Get the list of any VCs that can't be safely started yet
	longestRemainingTime := time.Duration(0)
	for _, vc := range vcs {
		remainingTime, err := checkValidatorClient(vc, newTagMap)
		if err != nil {
			return false, err
		}

		// If this VC has remaining time before it can be safely started, see if it's more than the current max
		if remainingTime > longestRemainingTime {
			longestRemainingTime = remainingTime
		}
	}

	// Show the slashing prevention dialog
	if longestRemainingTime > 0 {
		showSlashingDelay(longestRemainingTime)
	}
	return false, nil
}

// Get the map of tags
func getVcContainerTagParamMap(cfg *HyperdriveEthereumConfigSettings, vcs []string) (map[string]string, error) {
	containerTagMap := map[string]string{}

	modCfgs := cfg.GetAllModuleConfigs()
	for _, module := range modCfgs {
		vcInfo := module.GetValidatorContainerTagInfo()
		for name, tag := range vcInfo {
			fullName := cfg.Hyperdrive.GetDockerArtifactName(string(name))
			if _, exists := containerTagMap[fullName]; exists {
				return nil, fmt.Errorf("validator client map already had an entry named [%s]", fullName)
			}
			containerTagMap[fullName] = tag
		}
	}

	// SANITY CHECK
	for _, vc := range vcs {
		_, exists := containerTagMap[vc]
		if !exists {
			return nil, fmt.Errorf("validator client [%s] was missing from the slashing prevention check", vc)
		}
	}

	return containerTagMap, nil
}

func showSlashingDelay(remainingTime time.Duration) {
	fmt.Printf("%s=== WARNING ===\n", terminal.ColorRed)
	fmt.Println("You have changed validator clients. You must wait at least 15 minutes before safely starting them to prevent attesting to the same block twice, which would result in slashing your ETH.")
	fmt.Println("To prevent slashing, Hyperdrive will delay activating the new client until it is safe.")
	fmt.Println("See the documentation for a more detailed explanation: https://docs.nodeset.io")
	fmt.Printf("If you have read the documentation, understand the risks, and want to bypass this cooldown, run `hyperdrive service start --%s`.%s\n\n", ignoreSlashTimerFlag.Name, terminal.ColorReset)

	// Wait for 15 minutes
	safeStartTime := time.Now().Add(remainingTime)
	for remainingTime > 0 {
		fmt.Printf("Remaining time: %s", remainingTime)
		time.Sleep(1 * time.Second)
		remainingTime = time.Until(safeStartTime)
		fmt.Printf("%s\r", terminal.ClearLine)
	}

	fmt.Println(terminal.ColorReset)
	fmt.Println("You may now safely start Hyperdrive without fear of being slashed.")
}

func checkValidatorClient(vcName string, newTagMap map[string]string) (time.Duration, error) {
	// Get the current and pending VC images
	currentTag, err := GetDockerImage(vcName)
	if err != nil {
		return 0, fmt.Errorf("error getting Docker image tag for [%s]: %w", vcName, err)
	}
	currentVcType, err := getDockerImageName(currentTag)
	if err != nil {
		return 0, fmt.Errorf("error parsing current Docker image tag [%s] for [%s]: %w", currentTag, vcName, err)
	}
	pendingTag := newTagMap[vcName]
	pendingVcType, err := getDockerImageName(pendingTag)
	if err != nil {
		return 0, fmt.Errorf("error parsing pending Docker image tag [%s] for [%s]: %w", pendingTag, vcName, err)
	}

	// Compare the clients and warn if necessary
	if currentVcType == pendingVcType {
		fmt.Printf("Validator Client [%s] is still [%s] - no slashing prevention delay necessary.\n", vcName, currentVcType)
		return 0, nil
	} else {
		validatorFinishTime, err := GetDockerContainerShutdownTime(vcName)
		if err != nil {
			return 0, fmt.Errorf("error getting VC [%s] shutdown time: %w", vcName, err)
		}

		// If it hasn't exited yet, shut it down
		zeroTime := time.Time{}
		status, err := GetDockerStatus(vcName)
		if err != nil {
			return 0, fmt.Errorf("error getting VC [%s] status: %w", vcName, err)
		}
		if validatorFinishTime == zeroTime || status == "running" {
			fmt.Printf("%sValidator Client [%s] is currently running, stopping it...%s\n", terminal.ColorYellow, vcName, terminal.ColorReset)
			err := StopContainer(vcName)
			if err != nil {
				return 0, fmt.Errorf("error stopping VC [%s]: %w", vcName, err)
			}
			validatorFinishTime = time.Now()
		}

		// Print the warning and start the time lockout
		safeStartTime := validatorFinishTime.Add(15 * time.Minute)
		remainingTime := time.Until(safeStartTime)
		if remainingTime <= 0 {
			fmt.Printf("Validator Client [%s] has been offline for %s, which is long enough to prevent slashing.\n", vcName, time.Since(validatorFinishTime))
			return 0, nil
		}

		// If this VC has remaining time before it can be safely started, add it to the list
		if remainingTime > 0 {
			fmt.Printf("Validator Client [%s] has changed types from [%s] to [%s].\n", vcName, currentVcType, pendingVcType)
			fmt.Printf("Only %s has elapsed since you stopped it.\n", time.Since(validatorFinishTime))
		}

		// This can't be safely started, return its info
		return remainingTime, nil
	}
}

func StopContainer(containerName string) error {
	d, err := c.GetDocker()
	if err != nil {
		return err
	}
	return d.ContainerStop(context.Background(), containerName, dtc.StopOptions{})
}

// Get the current Docker image used by the given container
func GetDockerStatus(containerName string) (string, error) {
	ci, err := inspectContainer(c, containerName)
	if err != nil {
		return "", err
	}
	return ci.State.Status, nil
}

// Extract the image name from a Docker image string
func getDockerImageName(image string) (string, error) {
	// Return the empty string if the validator didn't exist (probably because this is the first time starting it up)
	if image == "" {
		return "", nil
	}

	reg := regexp.MustCompile(dockerImageRegex)
	matches := reg.FindStringSubmatch(image)
	if matches == nil {
		return "", fmt.Errorf("error parsing the Docker image string [%s]", image)
	}
	imageIndex := reg.SubexpIndex("image")
	if imageIndex == -1 {
		return "", fmt.Errorf("image name not found in Docker image [%s]", image)
	}

	imageName := matches[imageIndex]
	return imageName, nil
}

func GetDockerContainerShutdownTime(containerName string) (time.Time, error) {
	ci, err := inspectContainer(containerName)
	if err != nil {
		return time.Time{}, err
	}

	// Parse the time
	finishTime, err := time.Parse(time.RFC3339, strings.TrimSpace(ci.State.FinishedAt))
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing container [%s] exit time [%s]: %w", containerName, ci.State.FinishedAt, err)
	}
	return finishTime, nil
}

func GetDockerImage(containerName string) (string, error) {
	ci, err := inspectContainer(containerName)
	if err != nil {
		return "", err
	}
	return ci.Config.Image, nil
}

func inspectContainer(container string) (dt.ContainerJSON, error) {
	d, err := GetDocker()
	if err != nil {
		return dt.ContainerJSON{}, err
	}
	ci, err := d.ContainerInspect(context.Background(), container)
	if err != nil {
		return dt.ContainerJSON{}, fmt.Errorf("error inspecting container [%s]: %w", container, err)
	}
	return ci, nil
}

// Get the Docker client
func GetDocker() (*docker.Client, error) {
	if c.docker == nil {
		var err error
		c.docker, err = docker.NewClientWithOpts(docker.WithAPIVersionNegotiation())
		if err != nil {
			return nil, fmt.Errorf("error creating Docker client: %w", err)
		}
	}

	return c.docker, nil
}

func (c *HyperdriveEthereumConfigSettings) GetValidatorContainers(projectName string) ([]string, error) {
	d, err := c.GetDocker()
	if err != nil {
		return nil, err
	}
	cl, err := d.ContainerList(context.Background(), dtc.ListOptions{All: true})
	if err != nil {
		return nil, fmt.Errorf("error getting container list: %w", err)
	}

	// Find all of them that belong to the project
	containers := []string{}
	for _, container := range cl {
		isProjectContainer := false
		for _, name := range container.Names {
			name = strings.TrimPrefix(name, "/") // Docker throws a leading / on names
			if strings.HasPrefix(name, projectName) {
				isProjectContainer = true
			}
		}

		// This container belongs to the project
		if isProjectContainer && strings.Contains(container.Command, config.VcStartScript) {
			name := strings.TrimPrefix(container.Names[0], "/")
			containers = append(containers, name)
		}
	}
	return containers, nil
}
