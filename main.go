package main

import "fmt"

func main() {
	//fmt.Println("Hello!")

	//var config Configuration

	config := initializeConfig("config.yaml")
	sshConn, sshConfig := initializeConnection(config)
	clientConns := sshConn.dialConnection(sshConfig)
	fmt.Println(clientConns)
	clientSessions := sshConn.openSession(clientConns)
	fmt.Println(clientSessions)
	success := sshConn.executeTarFile(clientSessions)
	fmt.Println(success)
}
