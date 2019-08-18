package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

type Header struct {
	Comment string
	Extra   []byte
	ModTime time.Time
	Name    string
	OS      byte
}

func gzipit(source string, target string) error {
	reader, err := os.Open(source)

	if err != nil {
		fmt.Errorf("Cannot open source")
	}

	filename := filepath.Base(source)
	target = filepath.Join(target, fmt.Sprintf("%s.tar.gz", filename))
	writer, err := os.Create(target)

	if err != nil {
		fmt.Errorf("Cannot create target archive")
	}

	defer writer.Close()

	archiver := gzip.NewWriter(writer)
	archiver.Name = filename
	defer archiver.Close()

	_, err = io.Copy(archiver, reader)
	return err
}

func ungzipit(source string, target string) {
	reader, err := os.Open(source)

	if err != nil {
		fmt.Errorf("Cannot open source")
	}

	defer reader.Close()

	archive, err := gzip.NewReader(reader)

	if err != nil {
		fmt.Errorf("Cannot open archive")
	}

	defer archive.Close()

	target = filepath.Join(target, archive.Name)
	writer, err := os.Create(target)

	if err != nil {
		fmt.Errorf("Cannot extract archive")
	}

	defer writer.Close()

	_, err = io.Copy(writer, archive)
}
