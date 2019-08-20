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

func getFile(client *ssh.Client) {
	sftp, err := sftp.NewClient(client)

	if err != nil {
		fmt.Errorf("FUCK")
	}

	defer sftp.Close()

	srcFile, err := sftp.Open(srcPath + filename)

	if err != nil {
		fmt.Errorf("FUCK")
	}

	defer srcFile.Close()

	dstFile, err := os.Create(dstPath + filename)

	if err != nil {
		fmt.Errorf("FUCK")
	}

	defer dstFile.Close()

	srcFile.WriteTo(dstFile)
}
