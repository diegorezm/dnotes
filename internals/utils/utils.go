package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// DirExists returns true if path exists and is a directory.
func DirExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	if !info.IsDir() {
		return false, nil
	}

	return true, nil
}

// WriteJSONFile writes v as pretty‑printed JSON to the given path,
// creating or truncating the file.
func WriteJSONFile(path string, v any) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		return fmt.Errorf("failed to open %s for writing: %w", path, err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(v); err != nil {
		return fmt.Errorf("failed to encode JSON to %s: %w", path, err)
	}

	return nil
}
