package config

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

func NewProductionEnv() (*Env, error) {
	return NewLocalEnv("", "./config")
}

// NewLocalEnv reads in config from a local file or from the search paths.
// If a filepath is provided, it will read in the config from that file instead of searching.
// If no filepath is provided, it will search for a file named server-config.yaml in the search paths.
func NewLocalEnv(filePath string, searchPaths ...string) (*Env, error) {
	v := viper.New()

	v.SetConfigName("server-config")

	if filePath != "" {
		fmt.Println(filePath)
		v.SetConfigFile(filePath)
	} else {
		for _, sp := range searchPaths {
			v.AddConfigPath(sp)
		}
	}

	bindEnvsAndDefaults(v, defaultConfig)

	if filePath != "" || len(searchPaths) > 0 {
		if err := v.ReadInConfig(); err != nil {
			if !errors.As(err, &viper.ConfigFileNotFoundError{}) {
				return &Env{}, fmt.Errorf("could not read in config from viper: %w", err)
			}
		}
	}

	var cfg *Env
	if err := v.Unmarshal(&cfg); err != nil {
		return &Env{}, fmt.Errorf("could not unmarshal loaded viper config: %w", err)
	}

	return cfg, nil
}

func bindEnvsAndDefaults(vi *viper.Viper, iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			bindEnvsAndDefaults(vi, v.Interface(), append(parts, tv)...)
		default:
			p := strings.Join(append(parts, tv), ".")
			bindEnv(vi, p)
			setDefault(vi, p, v.Interface())
		}
	}
}

func setDefault(vi *viper.Viper, key string, value interface{}) {
	if !reflect.ValueOf(value).IsZero() {
		vi.SetDefault(key, value)
	}
}

func bindEnv(vi *viper.Viper, key string) {
	envVar := strings.ReplaceAll(key, ".", "_")
	envVar = strings.ToUpper(envVar)

	vi.BindEnv(key, envVar)
}
