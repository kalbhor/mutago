package mutago

import (
	"errors"
	"fmt"

	v1 "github.com/kalbhor/mutago/v1"
	v2 "github.com/kalbhor/mutago/v2"
)

const (
	id3v1Block = 128
	id3v1      = 1
	id3v2      = 2
	vSize      = 3
)

// Parse() loads the metadats struct with the id3 tags.
// It loads the tags based on the version.
func (m *Metadata) Parse() error {
	defer m.file.Close()
	// Get the id3 version
	v, err := m.Version()
	if err != nil {
		return fmt.Errorf("could not find appropriate version : %w", err)
	}

	// Based on the id3 version, parse the tags and load them inside the tags map.
	switch v {
	case id3v1:
		tags, err := v1.Parse(m.file)
		if err != nil {
			return fmt.Errorf("could not load id3v1 tags : %w", err)
		}
		m.tags = tags
	case id3v2:
		header := v2.ParseHeader(m.file)
		pos := int64(v2.HeaderSize)

		for pos <= int64(header.Size) {
			// Iterate over all ID3 frames
			frame, err := v2.ParseFrame(m.file)
			if frame.Size == 0 || err != nil {
				break
			}
			m.tags[frame.ID] = frame.Info
			pos, _ = m.file.Seek(0, 1)
		}
	}

	return nil
}

// Title() returns the "TIT2" tag from the loaded tags.
func (m *Metadata) Title() (string, error) {
	if val, check := m.tags["TIT2"]; check {
		return val, nil
	}
	return "", errors.New("Tag Error : Could not find title (TIT2)")
}

// Album() returns the "TALB" tag from the loaded tags.
func (m *Metadata) Album() (string, error) {
	if val, check := m.tags["TALB"]; check {
		return val, nil
	}
	return "", errors.New("Tag Error : Could not find album (TALB)")
}

// Artist() returns the "TPE1" tag from the loaded tags.
func (m *Metadata) Artist() (string, error) {
	if val, check := m.tags["TPE1"]; check {
		return val, nil
	}
	return "", errors.New("Tag Error : Could not find artist (TPE1)")
}

// Albumart() returns the "APIC" tag from the loaded tags.
func (m *Metadata) Albumart() (*v2.Albumart, error) {
	albumart, err := v2.ParseAlbumart(m.tags["APIC"])
	if err != nil {
		return albumart, err
	}
	return albumart, nil
}

// Get() returns an arbitrary tag from the loaded tags.
func (m *Metadata) Get(tag string) (string, error) {
	if val, check := m.tags[tag]; check {
		return val, nil
	}
	return "", errors.New("Tag Error : Could not find tag")
}

// List() returns a list of all loaded tags.
func (m *Metadata) List() (list []string) {
	for key := range m.tags {
		list = append(list, key)
	}
	return
}
