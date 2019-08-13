package main

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

type SSHConnection struct {
	Username string
	Password string
	Port     string
	Hosts    []string
}

func initializeConnection(config Configuration) (SSHConnection, *ssh.ClientConfig) {
	var sshConn SSHConnection
	sshConn.Username = config.Username
	sshConn.Password = config.Password
	sshConn.Port = config.Port
	sshConn.Hosts = config.Hosts

	sshConfig := &ssh.ClientConfig{
		User: sshConn.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshConn.Password),
		},
	}

	return sshConn, sshConfig
}

func (sshConn SSHConnection) DialConnection(sshConnec *ssh.ClientConfig) {
	for _, j := range sshConn.Hosts {
		fmt.Println(j)
	}
}
