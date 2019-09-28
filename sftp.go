package main

import (
	"fmt"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func (s SSHConnection) getFile(client *ssh.Client) []*os.File {
	homedir, _ := os.UserHomeDir()
	var dstFiles []*os.File

	for _, j := range s.Logs {
		sftp, err := sftp.NewClient(client)

		if err != nil {
			fmt.Errorf("Error")
		}

		defer sftp.Close()

		srcFile, err := sftp.Open(j)

		if err != nil {
			fmt.Errorf("Error")
		}

		defer srcFile.Close()

		h := slashSeperator(j)

		dstFile, err := os.Create(homedir + "/" + h)

		if err != nil {
			fmt.Errorf("Error")
		}

		defer dstFile.Close()
		srcFile.WriteTo(dstFile)

		dstFiles = append(dstFiles, dstFile)
	}

	return dstFiles
}
