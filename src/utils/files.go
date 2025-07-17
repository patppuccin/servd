package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"servd/src/models"
)

func resolveSpecPath() string {
	execPath, err := os.Executable()
	if err != nil {
		return ""
	}
	execDir := filepath.Dir(execPath)
	return filepath.Join(execDir, "services.spec.json")
}

func LoadSpec(path string) ([]models.ServiceSpec, error) {
	if path == "" {
		path = resolveSpecPath()
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("spec file not found at %q", path)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read spec file: %w", err)
	}

	var specs []models.ServiceSpec
	if err := json.Unmarshal(data, &specs); err != nil {
		return nil, fmt.Errorf("invalid JSON in spec file: %w", err)
	}

	return specs, nil
}

func SaveSpec(path string, spec *models.ServiceSpec) error {
	if spec == nil {
		return fmt.Errorf("spec cannot be nil")
	}
	if path == "" {
		path = resolveSpecPath()
	}

	specs, err := LoadSpec(path)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to load existing specs: %w", err)
	}

	for _, s := range specs {
		if s.Name == spec.Name {
			return fmt.Errorf("a service with name %q already exists", spec.Name)
		}
	}

	specs = append(specs, *spec)

	data, err := json.MarshalIndent(specs, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal spec list: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write spec file: %w", err)
	}

	return nil
}
