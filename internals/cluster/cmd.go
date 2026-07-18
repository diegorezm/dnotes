package cluster

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/diegorezm/dnotes/internals/config"
	"github.com/diegorezm/dnotes/internals/utils"
)

// Use sets the given directory as the active cluster in global config.
// It assumes dir is already validated or will return an error otherwise.
func Use(dir string) error {
	ok, err := IsValidDir(dir)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("%s is not a valid dnotes cluster (missing .dnotes)", dir)
	}

	// Update global config.
	if err := config.SetActiveCluster(dir); err != nil {
		return err
	}

	return nil
}

// Init initializes a new dnotes cluster in the given directory.
func Init(dir string) error {
	abs, err := filepath.Abs(dir)
	if err != nil {
		return fmt.Errorf("failed to resolve directory %s: %w", dir, err)
	}

	// Ensure the cluster root directory exists.
	if err := os.MkdirAll(abs, 0o755); err != nil {
		return fmt.Errorf("failed to create cluster directory %s: %w", abs, err)
	}

	// Ensure .dnotes exists inside the cluster.
	dnotesDir := filepath.Join(abs, ".dnotes")
	if err := os.MkdirAll(dnotesDir, 0o755); err != nil {
		return fmt.Errorf("failed to create .dnotes directory %s: %w", dnotesDir, err)
	}

	// Create a minimal local cluster config.
	localCfgPath := filepath.Join(dnotesDir, "config.json")
	localCfg := LocalConfig{
		Version: 1,
	}
	if err := utils.WriteJSONFile(localCfgPath, &localCfg); err != nil {
		return fmt.Errorf("failed to write local cluster config: %w", err)
	}

	// TODO: initialize SQLite index.db here
	// indexPath := filepath.Join(dnotesDir, "index.db")
	// err = initIndexDB(indexPath)

	return nil
}
