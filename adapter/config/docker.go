package config

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/docker/docker/client"
)

type DockerSettings struct {
	BeaconNodeContainer      string `json:"beaconNodeContainer"`
	ExecutionClientContainer string `json:"executionClientContainer"`
	DockerNetwork            string `json:"dockerNetwork"`
	BeaconDataVolume         string `json:"beaconDataVolume"`
	ExecutionDataVolume      string `json:"executionDataVolume"`
	ContainerTag             string `json:"containerTag"`

	docker     *client.Client
	dockerLock sync.Mutex
}

func (d *DockerSettings) GetDocker() (*client.Client, error) {
	d.dockerLock.Lock()
	defer d.dockerLock.Unlock()

	if d.docker == nil {
		var err error
		d.docker, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			return nil, fmt.Errorf("error creating Docker client: %w", err)
		}
	}

	return d.docker, nil
}

func (d *DockerSettings) StartService(services []string) error {
	serviceList := strings.Join(services, " ")

	cmd := exec.Command("docker-compose", "up", "-d", "--remove-orphans")
	cmd.Args = append(cmd.Args, services...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	fmt.Printf("Running: docker-compose up -d %s\n", serviceList)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to start services %s: %w", serviceList, err)
	}

	fmt.Println("Services started successfully.")
	return nil
}
