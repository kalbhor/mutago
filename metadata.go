package mutago

import (
	"errors"
	"fmt"
	"os"

	v1 "github.com/kalbhor/mutago/v1"
	v2 "github.com/kalbhor/mutago/v2"
)

const (
	id3v1Block = 128
	id3v1      = 1
	id3v2      = 2
	vSize      = 3
)

// Metadata is the main struct that holds the file being parsed and the loaded tags map.
type Metadata struct {
	version int
	tagger  Tagger
}

type Tagger interface {
	Parse() error
	// Set(string, interface{}) error
	Get(string) (string, error)
}

// Open accepts a path to an audio file and returns the loaded metadata struct.
func Open(fPath string) (*Metadata, error) {
	f, err := os.OpenFile(fPath, os.O_RDWR, os.ModeAppend)
	if err != nil {
		return nil, err
	}

	var tg Tagger

	v, err := Version(f)
	if err != nil {
		return nil, err
	}
	switch v {
	case id3v1:
		tg = v1.New(f)
	case id3v2:
		tg = v2.New(f)
	}

	return &Metadata{
		version: v,
		tagger:  tg,
	}, nil
}

func (m *Metadata) Parse() error {
	return m.tagger.Parse()
}

// Version checks the appropriate bytes in the file and tries to parse the ID3 version.
func Version(f *os.File) (int, error) {
	b := make([]byte, vSize)

	// Read the first 3 bytes of the file.
	// If the string representation is "ID3", then the file is encoded using ID3v2.
	_, err := f.ReadAt(b, 0)
	if err == nil && string(b) == "ID3" {
		return id3v2, nil
	}

	// Get file stats for file size
	stat, err := f.Stat()
	if err != nil {
		return 0, fmt.Errorf("could not describe file : %w", err)
	}

	// If the string representation of the first 3 bytes of the block of last 128 bytes
	// is "TAG", then the file is encoded using ID3v1.
	_, err = f.ReadAt(b, stat.Size()-id3v1Block)
	if err == nil && string(b) == "TAG" {
		return id3v1, nil
	}

	return 0, errors.New("ID3 Tags : None found")
}
