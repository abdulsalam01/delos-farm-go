package app

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var (
	path        = "./"
	serviceYaml = "app.yaml"
)

func LoadAppConfig() (Config, error) {
	var (
		config Config
		err    error
	)

	fileByte, err := ioutil.ReadFile(path + serviceYaml)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(fileByte, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
