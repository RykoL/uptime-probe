package config

import (
	"github.com/stretchr/testify/assert/yaml"
	"os"
	"time"
)

type Config struct {
	Monitors []*MonitorConfig `yaml:"monitors"`
}

type MonitorConfig struct {
	Name     string        `yaml:"name"`
	Url      string        `yaml:"url"`
	interval time.Duration `yaml:"interval"`
}

func LoadFromFile(fileName string) (*Config, error) {
	content, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = yaml.Unmarshal([]byte(content), cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
