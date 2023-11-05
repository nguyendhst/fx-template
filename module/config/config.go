package config

import (
	"fmt"
	"path/filepath"
)

type (
	Config struct {
		Env
	}

	ConfigReader interface {
		ReadConfig() (*Config, error)
		SetFilePath(filePath string)
	}

	ProductionConfigReader struct{}

	LocalConfigReader struct {
		filePath    string
		searchPaths []string
	}
)

func New(env string) (*Config, error) {
	configReader, err := NewConfigReader(env)
	if err != nil {
		panic(err)
	}

	cfg, err := configReader.ReadConfig()
	if err != nil {
		panic(err)
	}

	return cfg, nil
}

func (r *ProductionConfigReader) ReadConfig() (*Config, error) {
	// Read production config file
	return &Config{}, nil
}

func (r *ProductionConfigReader) SetFilePath(filePath string) {
	// Do nothing
}

func (r *LocalConfigReader) ReadConfig() (*Config, error) {
	// Search for config file in searchPaths and read it
	env, err := NewLocalEnv(r.filePath, r.searchPaths...)
	if err != nil {
		return nil, err
	}

	return &Config{*env}, nil
}

func (r *LocalConfigReader) SetFilePath(filePath string) {
	r.filePath = filePath
}

func NewConfigReader(env string) (ConfigReader, error) {
	switch env {
	case "production":
		return &ProductionConfigReader{}, nil
	case "development":
		rootPath, err := filepath.Abs(".")
		if err != nil {
			return nil, err
		}
		filePath, err := filepath.Abs(filepath.Join(rootPath, "server-config.dev.yaml"))
		if err != nil {
			return nil, err
		}
		searchPaths := []string{".", "../.."}
		return &LocalConfigReader{filePath, searchPaths}, nil
	default:
		return nil, fmt.Errorf("unknown environment: %s", env)
	}
}
