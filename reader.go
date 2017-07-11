package mutago

import (
	"errors"
	"os"

	"github.com/makebyte/mutago/v2"

	"github.com/makebyte/mutago/v1"
)

type Metadata struct {
	file os.File           // The file containing tags
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
	version, err := ID3Version(f)
	if err != nil {
		return nil, err
	}
	metadata := &Metadata{}
	tags := make(map[string]string)

	switch version {

	case 1:
		tags, err = v1.Parse(f)
		if err != nil {
			return nil, err
		}

		// ** ID3 version 1 implementation pending **
	case 2:
		header := v2.ParseHeader(f)
		pos := int64(v2.HeaderSize)

		for pos <= int64(header.Size) { // Iterate over all ID3 frames
			frame := v2.ParseFrame(f)
			tags[frame.ID] = frame.Info
			pos, _ = f.Seek(0, 1)
		}

		metadata.tags = tags

	}

	return metadata, nil
}
