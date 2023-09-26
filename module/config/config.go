package config

import "fmt"

type (
	Config struct {
		Env
	}

	ConfigReader interface {
		ReadConfig() (*Config, error)
	}

	ProductionConfigReader struct{}

	LocalConfigReader struct {
		filePath    string
		searchPaths []string
	}
)

func New(env string) (*Config, error) {
	configReader, err := NewConfigReader(env, "", []string{})
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

func (r *LocalConfigReader) ReadConfig() (*Config, error) {
	// Search for config file in searchPaths and read it
	env, err := NewLocalEnv(r.filePath, r.searchPaths...)
	if err != nil {
		return nil, err
	}

	return &Config{*env}, nil
}

func NewConfigReader(env string, filePath string, searchPaths []string) (ConfigReader, error) {
	switch env {
	case "production":
		return &ProductionConfigReader{}, nil
	case "local":
		return &LocalConfigReader{filePath, searchPaths}, nil
	default:
		return nil, fmt.Errorf("unknown environment: %s", env)
	}
}
