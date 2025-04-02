package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"ripe-data-exporter/pkg/types"
)

func SaveToFile(data *types.Response, filename string) error {

	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("JSON marshaling error: %w", err)
	}

	if err := os.WriteFile(filename, jsonData, 0644); err != nil {
		return fmt.Errorf("file write error: %w", err)
	}

	return nil
}
