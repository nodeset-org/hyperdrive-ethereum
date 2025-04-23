package config

import "testing"

func TestBlah(t *testing.T) {
	// Revert to the initial setup
	err := testMgr.RevertSnapshot(initSnapshot)
	if err != nil {
		fail("Error reverting to initial snapshot: %v", err)
	}
	defer handle_panics()
}
