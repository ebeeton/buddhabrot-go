package main

import (
	"os"
	"path/filepath"
)

const plotsDirectory = "/plots"

func readPng(filename string) ([]byte, error) {
	path := filepath.Join(plotsDirectory, filename)

	if b, err := os.ReadFile(path); err != nil {
		return nil, err
	} else {
		return b, nil
	}
}
