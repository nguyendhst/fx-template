package config

import (
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
	configReader, err := NewConfigReader("local", configFile, []string{})
	if err != nil {
		t.Fatal(err)
	}

	config, err := configReader.ReadConfig()
	if err != nil {
		t.Fatal(nil)
	}

	// fmt.Printf("%+v\n", config)

	// Check that the config was read
	assert.NotEmpty(t, config)
}
