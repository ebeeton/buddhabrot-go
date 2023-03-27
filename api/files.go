package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const plotsDirectory = "/plots"

func writePng(file []byte) (string, error) {
	h := sha256.New()
	if _, err := h.Write(file); err != nil {
		return "", err
	}
	filename := fmt.Sprintf("%x", h.Sum(nil))
	path := filepath.Join(plotsDirectory, filename)

	if err := os.WriteFile(path, file, os.ModeAppend); err != nil {
		return "", err
	}
	log.Printf("Wrote file %s successfully.", filename)
	return filename, nil
}
