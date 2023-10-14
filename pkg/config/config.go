package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type config struct {
	FilePath   string `yaml:"file_path"`
	DriverName string `yaml:"driver_name"`
	Timeout    int    `yaml:"timeout"`
}

func InitConfig() (*config, error) {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	var config config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
