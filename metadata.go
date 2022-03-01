package mutago

import (
	"errors"
	"fmt"
	"os"
)

// Metadata is the main struct that holds the file being parsed and the loaded tags map.
type Metadata struct {
	file *os.File          // The file containing tags
	tags map[string]string // tags and their values
}

// Open accepts a path to an audio file and returns the loaded metadata struct.
func Open(fPath string) (*Metadata, error) {
	f, err := os.Open(fPath)
	if err != nil {
		return nil, fmt.Errorf("could not open file %s : %w", fPath, err)
	}

	return &Metadata{
		file: f,
		tags: make(map[string]string),
	}, nil
}

// Version checks the appropriate bytes in the file and tries to parse the ID3 version.
func (m *Metadata) Version() (int, error) {
	b := make([]byte, vSize)

	// Read the first 3 bytes of the file.
	// If the string representation is "ID3", then the file is encoded using ID3v2.
	_, err := m.file.ReadAt(b, 0)
	if err == nil && string(b) == "ID3" {
		return id3v2, nil
	}

	// Get file stats for file size
	stat, err := m.file.Stat()
	if err != nil {
		return 0, fmt.Errorf("could not describe file : %w", err)
	}

	// If the string representation of the first 3 bytes of the block of last 128 bytes
	// is "TAG", then the file is encoded using ID3v1.
	_, err = m.file.ReadAt(b, stat.Size()-id3v1Block)
	if err == nil && string(b) == "TAG" {
		return id3v1, nil
	}

	return 0, errors.New("ID3 Tags : None found")
}
