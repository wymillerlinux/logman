package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// this file is for generating the config if it doesn't exist, mainly for the idiots like me :)
// this is not really a priority right now so it sitting in the backburner as you read this
// signed Wyatt J. Miller, the awesome guy who doesn't generate configs *thumbs up*

var config string = "config.yaml"

func doesConfigExist() bool {
	_, err := os.Stat(config)

	if err != nil {
		return false
	}

	return true
}

func createTemplateConfig(result bool) {
	if result == false {
		os.Create(config)
	} else {
		fmt.Println("Configuration file already created")
	}
}

func writeTemplateConfig(filename string) {
	hosts := make([]string, 1)
	logs := make([]string, 1)

	config := Configuration{
		"jbob",
		"jbobisawesome",
		22,
		hosts,
		logs,
	}

	data, err := yaml.Marshal(&config)

	fmt.Printf(string(data))
	_ = ioutil.WriteFile(filename, data, 0600)

	if err != nil {
		fmt.Errorf("Couldn't create the configuration file", err)
	}

	fmt.Println("Template for logman's configuration complete!")
	fmt.Println("")
	fmt.Println("Fill in the necessary credentials you need. If")
	fmt.Println("you need assitance, check out the wiki for more.")
	os.Exit(0)
}
