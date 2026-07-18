package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/diegorezm/dnotes/internals/utils"
)

// ConfigDir returns the user config directory for dnotes on Linux.
func ConfigDir() (string, error) {
	if xdg := os.Getenv("XDG_CONFIG_HOME"); xdg != "" {
		return filepath.Join(xdg, "dnotes"), nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not determine home directory: %w", err)
	}

	return filepath.Join(home, ".config", "dnotes"), nil
}

// ConfigPath returns the full path to the global config file.
func ConfigPath() (string, error) {
	dir, err := ConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "config.json"), nil
}

// LoadOrInitGlobal loads the global config or creates a default one if missing.
func LoadOrInitGlobal() (*Config, error) {
	path, err := ConfigPath()
	if err != nil {
		return nil, err
	}

	f, err := os.Open(path)
	if err == nil {
		defer f.Close()

		var cfg Config
		if err = json.NewDecoder(f).Decode(&cfg); err != nil {
			return nil, fmt.Errorf("failed to parse global config: %w", err)
		}
		return &cfg, nil
	}

	if !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("failed to open global config: %w", err)
	}

	if _, err := ensureConfigDir(); err != nil {
		return nil, err
	}

	defaultCfg := &Config{
		Version: 1,
		Editor: Editor{
			Command: defaultEditorCommand(),
		},
		Cluster: Cluster{
			Active: "",
		},
	}

	if err := utils.WriteJSONFile(path, defaultCfg); err != nil {
		return nil, fmt.Errorf("failed to write default global config: %w", err)
	}

	return defaultCfg, nil
}

// SaveGlobal writes the given config to the global config file.
func SaveGlobal(cfg *Config) error {
	path, err := ConfigPath()
	if err != nil {
		return err
	}

	if _, err := ensureConfigDir(); err != nil {
		return err
	}

	if err := utils.WriteJSONFile(path, cfg); err != nil {
		return err
	}

	return nil
}

// SetEditorCommand updates the editor.command in global config.
func SetEditorCommand(command string) error {
	cfg, err := LoadOrInitGlobal()
	if err != nil {
		return err
	}

	cfg.Editor.Command = command

	return SaveGlobal(cfg)
}

// SetActiveCluster updates cluster.active in global config.
func SetActiveCluster(path string) error {
	cfg, err := LoadOrInitGlobal()
	if err != nil {
		return err
	}

	cfg.Cluster.Active = path

	return SaveGlobal(cfg)
}

// defaultEditorCommand returns $EDITOR or falls back to nvim.
func defaultEditorCommand() string {
	if env := os.Getenv("EDITOR"); env != "" {
		return env
	}
	return "nvim"
}

// ensureConfigDir ensures the global config directory exists.
func ensureConfigDir() (string, error) {
	dir, err := ConfigDir()
	if err != nil {
		return "", err
	}
	if mkErr := os.MkdirAll(dir, 0o755); mkErr != nil {
		return "", fmt.Errorf("failed to create config directory %s: %w", dir, mkErr)
	}
	return dir, nil
}
