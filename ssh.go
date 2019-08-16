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

// TODO: I've got two slices of the same type. To take one slice out, some refactoring is needed
type SSHSessions []*ssh.Session

type SSHTarFile []*ssh.Session

type SSHPush []*ssh.Client

type SSHSuccess []error

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
	clientConn := SSHCleints{}

	for _, j := range s.Hosts {
		hostPort := fmt.Sprintf("%s:22", j)
		connection, err := ssh.Dial("tcp", hostPort, sshConfig)

		if err != nil {
			fmt.Errorf("Can't connect", err)
		}

		clientConn = append(clientConn, connection)
	}

	return clientConn
}

func (s SSHConnection) openSession(client SSHCleints) SSHSessions {
	clientSessions := SSHSessions{}

	for _, j := range client {
		session, err := j.NewSession()

		if err != nil {
			fmt.Errorf("Can't open session", err)
		}

		clientSessions = append(clientSessions, session)
	}

	return clientSessions
}

func (s SSHConnection) executeTarFile(execute SSHSessions) SSHSuccess {
	// execute order 66 lol
	success := SSHSuccess{}

	for _, j := range execute {
		// this is just a placeholder, change to the actual tarring executable
		err := j.Run("echo 'hello world'! > test.txt")

		if err != nil {
			fmt.Errorf("Can't execute program", err)
		}

		success = append(success, err)
	}

	return success
}
