package app

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

var (
	path        = "./"
	serviceYaml = "./config/app.yaml"
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
