package v1

import (
	"fmt"
	"io"
	"os"
)

// Metadata is the main struct that holds the file being parsed and the tags map.
type Metadata struct {
	file *os.File          // The file containing tags
	tags map[string]string // tags and their values
}

// New returns a new struct of metadata.
func New(f *os.File) *Metadata {
	return &Metadata{
		file: f,
		tags: make(map[string]string),
	}
}

// Parse loads all the available id3 tags into the metadata tags map.
func (m *Metadata) Parse() error {
	if _, err := m.file.Seek(-id3v1Block, io.SeekEnd); err != nil {
		return fmt.Errorf("could not seek last %d bytes : %w", id3v1Block, err)
	}

	data := make([]byte, id3v1Block)
	_, err := io.ReadFull(m.file, data)
	if err != nil {
		return fmt.Errorf("could not read last %d bytes : %w", id3v1Block, err)
	}

	m.tags["TIT2"] = string(data[3:33])   //TITLE
	m.tags["TPE1"] = string(data[33:63])  //ARTIST
	m.tags["TALB"] = string(data[63:93])  //ALBUM
	m.tags["TYER"] = string(data[93:97])  //YEAR
	m.tags["COMM"] = string(data[97:127]) //COMMENTS
	if int(data[127]) < len(genres) {
		m.tags["TCON"] = genres[int(data[127])] //GENRE
	}

	return nil
}

// Get looks up the tags map for an arbitrary id3 tag.
func (m *Metadata) Get(tag string) (string, error) {
	if val, ok := m.tags[tag]; ok {
		return val, nil
	}

	return "", fmt.Errorf("tag not '%s' not found in tag map", tag)
}
