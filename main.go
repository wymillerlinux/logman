package main

import "fmt"

func main() {
	config := initializeConfig("config.yaml")
	sshConn, sshConfig := initializeConnection(config)
	clientConns := sshConn.dialConnection(sshConfig)
	clientSessions := sshConn.openSession(clientConns)
	success := sshConn.executeTarFile(clientSessions)
	fmt.Println(success)
}
