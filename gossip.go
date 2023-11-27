package gossip

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Port      int    `yaml:"port"`
	Host      string `yaml:"host"`
	TimeoutMS int    `yaml:"timeout"`
}

func YamlToConfig(path string) (Config, error) {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		return Config{}, err
	}

	fmt.Printf("%+v\n", cfg)
	return cfg, nil
}
