package cli

import (
	"fmt"

	"github.com/gentleman-programming/gentle-ai/v2/internal/statecoord"
)

// installStateLockPath preserves the trusted-home lock path contract while
// delegating path resolution to the shared state coordinator.
func installStateLockPath(homeDir string) (string, error) {
	lockPath, err := statecoord.LockPath(homeDir)
	if err != nil {
		return "", fmt.Errorf("resolve install state lock path: %w", err)
	}
	return lockPath, nil
}

func withInstallStateLock(homeDir string, operation func() error) error {
	return statecoord.WithLock(homeDir, operation)
}
