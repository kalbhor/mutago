package mutago

import (
	"errors"
	"os"

	"github.com/makebyte/mutago/v2"

	"github.com/makebyte/mutago/v1"
)

type Metadata struct {
	file *os.File          // The file containing tags
	tags map[string]string // tags and their values
}

func (m Metadata) Title() (string, error) {
	if val, check := m.tags["TIT2"]; check {
		return val, nil
	}
	return "", errors.New("Tag Error : Could not find title (TIT2)")
}

func (m Metadata) Album() (string, error) {
	if val, check := m.tags["TALB"]; check {
		return val, nil
	}
	return "", errors.New("Tag Error : Could not find album (TALB)")
}

func (m Metadata) Artist() (string, error) {
	if val, check := m.tags["TPE1"]; check {
		return val, nil
	}
	return "", errors.New("Tag Error : Could not find artist (TPE1)")
}

func (m Metadata) Albumart() (*v2.Albumart, error) {
	/*
	   Returns the image in byte array
	*/
	albumart, err := v2.ParseAlbumart(m.tags["APIC"])
	if err != nil {
		return albumart, err
	}
	return albumart, nil
}

func (m Metadata) Get(tag string) (string, error) {
	/*
		Returns the value of a tag
		eg : TALB -> Album name
	*/
	if val, check := m.tags[tag]; check {
		return val, nil
	}
	return "", errors.New("Tag Error : Could not find tag")
}

func (m Metadata) List() (list []string) {
	/*
		Provides a list of available tags
		in a file
	*/
	for key := range m.tags {
		list = append(list, key)
	}
	return
}

func (m Metadata) Close() {
	/*
		Close the file
	*/

	// Pending : Cleanup after writing to file
	defer m.file.Close()
}

func Open(file string) (*Metadata, error) {
	/*
		Open file and parse tags
	*/
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	metadata := &Metadata{}
	metadata.file = f
	tags := make(map[string]string)

	version, err := ID3Version(metadata.file)
	if err != nil {
		return nil, err
	}
	switch version {

	case 0: // No metadata (or invalid metadata)
		return nil, nil

	case 1: // ID3v1
		tags, err = v1.Parse(metadata.file)
		if err != nil {
			return nil, err
		}

	case 2: // ID3v2
		header := v2.ParseHeader(metadata.file)
		pos := int64(v2.HeaderSize)

		for pos <= int64(header.Size) { // Iterate over all ID3 frames
			frame, err := v2.ParseFrame(metadata.file)
			if frame.Size == 0 || err != nil {
				break
			}
			tags[frame.ID] = frame.Info
			pos, _ = metadata.file.Seek(0, 1)
		}

		metadata.tags = tags

	}

	return metadata, nil
}
