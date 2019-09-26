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
	} else {
		return true
	}
}

func createTemplateConfig(result bool) {
	if result == false {
		os.Create(config)
	} else {
		fmt.Println()
	}
}

func writeTemplateConfig(filename string) {
	file := []byte(filename)
	_ = ioutil.WriteFile("config.yaml", file, 0600)

	config := Configuration{}

	_, err := yaml.Marshal(config)

	if err != nil {
		panic(err)
	}

	fmt.Println("Template for logman's configuration complete!")
	fmt.Println("")
	fmt.Println("Fill in the necessary credentials you need. If")
	fmt.Println("you need assitance, check out the wiki for more.")
	os.Exit(0)
}
