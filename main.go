package main

import "fmt"

func main() {
	config := initializeConfig("config.yaml")
	sshConn, sshConfig := initializeConnection(config)
	clientConns := sshConn.dialConnection(sshConfig)
	_ = sshConn.openSession(clientConns)
	success := sshConn.executeSFTP(clientConns)
	fmt.Println(success)
}
