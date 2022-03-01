package v1

import (
	"fmt"
	"io"
)

func Parse(reader io.ReadSeeker) (map[string]string, error) {
	if _, err := reader.Seek(-id3v1Block, io.SeekEnd); err != nil {
		return nil, fmt.Errorf("could not seek last %d bytes : %w", id3v1Block, err)
	}

	data := make([]byte, id3v1Block)
	_, err := io.ReadFull(reader, data)
	if err != nil {
		return nil, fmt.Errorf("could not read last %d bytes : %w", id3v1Block, err)
	}

	tags := make(map[string]string)

	tags["TIT2"] = string(data[3:33])   //TITLE
	tags["TPE1"] = string(data[33:63])  //ARTIST
	tags["TALB"] = string(data[63:93])  //ALBUM
	tags["TYER"] = string(data[93:97])  //YEAR
	tags["COMM"] = string(data[97:127]) //COMMENTS
	if int(data[127]) < len(genres) {
		tags["TCON"] = genres[int(data[127])] //GENRE
	}

	return tags, nil
}
