package v1

import (
	"io"
	"os"
)

func Parse(reader io.ReadSeeker) (map[string]string, error) {

	reader.Seek(-128, os.SEEK_END)
	/*
	   v1 tags are in the last 128 bytes
	*/
	data := make([]byte, 128)

	_, err := io.ReadFull(reader, data)
	if err != nil {
		return nil, err
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
