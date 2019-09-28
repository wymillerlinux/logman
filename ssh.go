package main

import (
	"fmt"
	"os"

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
	Port     int
	Hosts    []string
	Logs     []string
}

type SSHClients []*ssh.Client

// TODO: I've got two slices of the same type. To take one slice out, some refactoring is needed
type SSHSessions []*ssh.Session

type SSHTarFile []*ssh.Session

type SSHPush []*ssh.Client

type SSHSFTP []string

type SSHSuccess []error

func initializeConnection(config Configuration) (*SSHConnection, *ssh.ClientConfig) {
	sshConn := &SSHConnection{
		Username: config.Username,
		Password: config.Password,
		Port:     config.Port,
		Hosts:    config.Hosts,
		Logs:     config.Logs,
	}

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

func (s SSHConnection) dialConnection(sshConfig *ssh.ClientConfig) SSHClients {
	clientConn := SSHClients{}

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

func (s SSHConnection) openSession(client SSHClients) SSHSessions {
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

func (s SSHConnection) executeSFTP(execute SSHClients) SSHSFTP {
	// execute order 66 lol
	sftp := SSHSFTP{}
	homedir, _ := os.UserHomeDir()

	for _, j := range execute {
		slashed := s.getFile(j)
		for _, k := range slashed {
			name := osfileToSting(k)
			err := gzipit(homedir+"/"+name, ".")

			if err != nil {
				fmt.Errorf("Cannot gzip file(s)", err)
			}
		}
	}

	return sftp
}
