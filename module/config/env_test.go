package config

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadYAMLConfig(t *testing.T) {
	// Create a temporary directory for the test
	dir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	// Find YAML config file in root of project
	configFile, err := filepath.Abs("../../server-config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	// Read the YAML config file
	config, err := New(configFile)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", config)

	// Check that the config was read
	assert.NotEmpty(t, config)
}
