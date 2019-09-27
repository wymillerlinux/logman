package main

import (
	"fmt"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

var (
	dstPath  string = "/home/wyatt/"
	srcPath  string = "/var/log/"
	filename string = "dpkg.log"
)

func (s SSHConnection) getFile(client *ssh.Client) {
	sftp, err := sftp.NewClient(client)
	//var srcFile []*os.File
	//var dstFile []*os.File

	if err != nil {
		fmt.Errorf("Error")
	}

	defer sftp.Close()

	for _, j := range s.Logs {
		srcFile, err := sftp.Open(srcPath + j)

		if err != nil {
			fmt.Errorf("Error")
		}

		defer srcFile.Close()

		dstFile, err := os.Create(dstPath + j)

		if err != nil {
			fmt.Errorf("Error")
		}

		defer dstFile.Close()
		srcFile.WriteTo(dstFile)
	}

}
