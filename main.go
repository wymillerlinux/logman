package main

func main() {
	//fmt.Println("Hello!")

	//var config Configuration

	config := initializeConfig("config.yaml")
	sshConn, sshConfig := initializeConnection(config)
	sshConn.DialConnection(sshConfig)
}
