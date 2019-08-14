package main

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

// modes := ssh.TerminalModes{
// 	ssh.ECHO: 0,
// 	ssh.TTY_OP_ISPEED: 14400,
// 	ssh.TTY_OP_OSPEED: 14400,
// }

type SSHConnection struct {
	Username string
	Password string
	Port     string
	Hosts    []string
}

type SSHCleints []*ssh.Client

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
		HostKeyCallback: sshConn.getHostKeys(),
	}

	return sshConn, sshConfig
}

func (s SSHConnection) getHostKeys() ssh.HostKeyCallback {
	return ssh.InsecureIgnoreHostKey()
}

func (s SSHConnection) dialConnection(sshConfig *ssh.ClientConfig) SSHCleints {
	ClientConn := SSHCleints{}

	for _, j := range s.Hosts {
		hostPort := fmt.Sprintf("%s:22", j)
		connection, err := ssh.Dial("tcp", hostPort, sshConfig)

		if err != nil {
			fmt.Errorf("Can't connect", err)
		}

		ClientConn = append(ClientConn, connection)
	}

	return ClientConn
}
