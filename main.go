package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Hello!")

	//var config Configuration

	config := initializeConfig("config.yaml")
	//fmt.Println(config.Hosts)
	sshConn, sshConfig := initializeConnection(config)
	//fmt.Println(sshConn.Hosts)
	clientConns := sshConn.dialConnection(sshConfig)
	fmt.Println(clientConns)
}
