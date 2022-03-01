package v1

import (
	"fmt"
	"io"
	"strconv"
)

func (m *Metadata) Set(tag string, data interface{}) error {
	switch tag {
	case "TIT2":
		if d, ok := data.(string); ok {
			return m.SetTitle(d)
		}
		return fmt.Errorf("invalid data type (%T) for tag (%s)", data, tag)
	case "TPE1":
		if d, ok := data.(string); ok {
			return m.SetArtist(d)
		}
		return fmt.Errorf("invalid data type (%T) for tag (%s)", data, tag)
	case "TALB":
		if d, ok := data.(string); ok {
			return m.SetAlbum(d)
		}
		return fmt.Errorf("invalid data type (%T) for tag (%s)", data, tag)
	case "TYER":
		if d, ok := data.(string); ok {
			return m.SetYear(d)
		}
		return fmt.Errorf("invalid data type (%T) for tag (%s)", data, tag)
	case "COMM":
		if d, ok := data.(string); ok {
			return m.SetComments(d)
		}
		return fmt.Errorf("invalid data type (%T) for tag (%s)", data, tag)
	case "TCON":
		if d, ok := data.(int); ok {
			return m.SetGenre(d)
		}
		return fmt.Errorf("invalid data type (%T) for tag (%s)", data, tag)
	default:
		return fmt.Errorf("could not set %s tag, ID3v1 allows setting limited tags.", tag)
	}
}

func (m *Metadata) SetTitle(str string) error {
	if len(str) > tagSize {
		return fmt.Errorf("cannot set tag with length greather than %d", tagSize)
	}

	padding := make([]byte, tagSize-len(str))
	data := append([]byte(str), padding...)

	return m.writeAt(3, data)
}

func (m *Metadata) SetArtist(str string) error {
	if len(str) > tagSize {
		return fmt.Errorf("cannot set tag with length greather than %d", tagSize)
	}

	padding := make([]byte, tagSize-len(str))
	data := append([]byte(str), padding...)

	return m.writeAt(33, data)
}

func (m *Metadata) SetAlbum(str string) error {
	if len(str) > tagSize {
		return fmt.Errorf("cannot set tag with length greather than %d", tagSize)
	}

	padding := make([]byte, tagSize-len(str))
	data := append([]byte(str), padding...)

	return m.writeAt(63, data)
}

func (m *Metadata) SetYear(str string) error {
	if len(str) > 4 {
		return fmt.Errorf("cannot set year tag with length greather than %d", tagSize)
	}

	return m.writeAt(97, []byte(str))
}

func (m *Metadata) SetComments(str string) error {
	if len(str) > tagSize {
		return fmt.Errorf("cannot set tag with length greather than %d", tagSize)
	}

	padding := make([]byte, tagSize-len(str))
	data := append([]byte(str), padding...)

	return m.writeAt(93, data)
}

func (m *Metadata) SetGenre(n int) error {
	if n > len(genres) {
		return fmt.Errorf("invalid index value for genre")
	}

	return m.writeAt(127, []byte(strconv.Itoa(n)))
}

func (m *Metadata) writeAt(n int64, data []byte) error {
	if _, err := m.file.Seek(-id3v1Block+n, io.SeekEnd); err != nil {
		return fmt.Errorf("could not seek last %d bytes : %w", id3v1Block, err)
	}

	if _, err := m.file.Write(data); err != nil {
		return fmt.Errorf("could not write title : %w", err)
	}

	return nil
}
