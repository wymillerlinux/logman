package main

import "fmt"

func main() {
	//fmt.Println("Hello!")

	//var config Configuration

	config := initializeConfig("config.yaml")
	sshConn, sshConfig := initializeConnection(config)
	clientConns := sshConn.dialConnection(sshConfig)
	clientSessions := sshConn.openSession(clientConns)
	fmt.Println(clientSessions)
}
