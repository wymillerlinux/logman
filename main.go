package main

import (
	"fmt"

	"golang.org/x/crypto/ssh"
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

	sshSuccess, err := ssh.Dial("tcp", "crash.local:22", sshConfig)

	if err != nil {
		panic(err)
	}

	fmt.Println(sshSuccess)
}
