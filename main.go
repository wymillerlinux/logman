package main

import "log"

func main() {
	config := initializeConfig("config.yaml")
	sshConn, sshConfig := initializeConnection(config)
	clientConns := sshConn.dialConnection(sshConfig)
	_ = sshConn.openSession(clientConns)
	success := sshConn.executeSFTP(clientConns)
	log.Println(success)
}
