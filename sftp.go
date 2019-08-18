package main

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

var (
	dstPath  string = "/root/"
	srcPath  string = "/var/log/"
	filename string = "dmesg"
)

func getFile(client *ssh.Client) {
	t := time.Now()
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

	timeResult := timeToString(t)
	dstFile, err := os.Create(dstPath + filename + timeResult)

	if err != nil {
		fmt.Errorf("FUCK")
	}

	defer dstFile.Close()

	srcFile.WriteTo(dstFile)
}

func timeToString(currentTime time.Time) string {
	return currentTime.String()
}
