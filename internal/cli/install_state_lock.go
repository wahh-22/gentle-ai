package cli

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/gentleman-programming/gentle-ai/v2/internal/reviewtransaction"
	"github.com/gentleman-programming/gentle-ai/v2/internal/state"
)

func installStateLockPath(homeDir string) string {
	// Home is the trusted provider root. Resolve platform aliases such as
	// macOS /var -> /private/var before the lock's no-follow directory walk;
	// descendants such as .gentle-ai remain subject to that hardened walk.
	if resolved, err := filepath.EvalSymlinks(homeDir); err == nil {
		homeDir = resolved
	}
	return state.Path(homeDir) + ".lock"
}

func withInstallStateLock(homeDir string, operation func() error) (err error) {
	lock, err := reviewtransaction.AcquireAuthorityFileLock(installStateLockPath(homeDir))
	if err != nil {
		return fmt.Errorf("acquire install state lock: %w", err)
	}
	defer func() {
		if releaseErr := lock.Release(); releaseErr != nil {
			err = errors.Join(err, fmt.Errorf("release install state lock: %w", releaseErr))
		}
	}()
	return operation()
}
