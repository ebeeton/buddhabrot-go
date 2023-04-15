// Package files provides functions for reading and writing files.
package files

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const plotsDirectory = "/plots"

// Read reads the filename from the plots directory and returns its contents as
// a slice of bytes.
func Read(filename string) ([]byte, error) {
	path := filepath.Join(plotsDirectory, filename)

	if b, err := os.ReadFile(path); err != nil {
		return nil, err
	} else {
		return b, nil
	}
}

// Write writes a slice of bytes to the plots directory. The filename returned
// is the SHA-256 sum of the bytes expressed as a hexadecimal string.
func Write(file []byte) (string, error) {
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
