package main

type sshConnection struct {
	Username string
	Password string
	Port     string
	Hosts    []string
}
