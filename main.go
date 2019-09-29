package main

import "log"

func main() {
	isCreated := doesConfigExist()

	if isCreated == false {
		createTemplateConfig(isCreated)
		writeTemplateConfig("config.yaml")
	}

	config := initializeConfig("config.yaml")
	sshConn, sshConfig := initializeConnection(config)
	clientConns := sshConn.dialConnection(sshConfig)
	_ = sshConn.openSession(clientConns)
	success := sshConn.executeSFTP(clientConns)
	log.Println(success)
}
