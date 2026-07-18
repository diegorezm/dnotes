package cluster

import (
	"fmt"
	"path/filepath"

	"github.com/diegorezm/dnotes/internals/utils"
)

// IsValidDir returns true if dir contains a .dnotes directory.
func IsValidDir(dir string) (bool, error) {
	ok, err := utils.DirExists(dir)
	if err != nil {
		return false, fmt.Errorf("failed to check directory %s: %w", dir, err)
	}
	if !ok {
		return false, nil
	}

	dnotesPath := filepath.Join(dir, ".dnotes")
	dok, err := utils.DirExists(dnotesPath)
	if err != nil {
		return false, fmt.Errorf("failed to check %s: %w", dnotesPath, err)
	}
	if !dok {
		return false, nil
	}

	return true, nil
}
