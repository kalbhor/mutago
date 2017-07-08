package mutago

import (
	"errors"
	"os"

	"github.com/makebyte/mutago/v2"
)

type Metadata struct {
	file os.File           // The file containing tags
	tags map[string]string // tags and their values
}

func (m Metadata) Get(tag string) (string, error) {
	if val, check := m.tags[tag]; check {
		return val, nil
	}
	return "", errors.New("Tag Error : Could not find tag")
}

func (m Metadata) Close() {
	defer m.file.Close()
}

func Open(file string) (*Metadata, error) {
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
