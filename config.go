package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Username string
	Password string
	Port     int
	Hosts    []string
}

func initializeConfig(filename string) Configuration {
	var config Configuration
	source, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(source, &config)

	if err != nil {
		panic(err)
	}

	return config
}
