package v2

import (
	"fmt"
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
	header := parseHeader(m.file)
	pos := int64(HeaderSize)

	for pos <= int64(header.Size) {
		// Iterate over all ID3 frames
		frame, err := parseFrame(m.file)
		if frame.Size == 0 || err != nil {
			break
		}
		m.tags[frame.ID] = frame.Info
		pos, _ = m.file.Seek(0, 1)
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
