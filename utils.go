package mutago

import (
	"errors"
	"os"
)

func ID3Version(f *os.File) (int, error) {
	/*
		Returns ID3 version.

		ID3v1 = 1
		ID3v2 = 2
		No ID3 tag found = error
	*/

	b := make([]byte, 3)
	stat, _ := f.Stat() // Get file stats for file size

	if f.ReadAt(b, stat.Size()-128); string(b) == "TAG" {
		return 1, nil // ID3v1 has 'TAG' in first 3 bytes in last 128 bytes
	} else if f.ReadAt(b, 0); string(b) == "ID3" {
		return 2, nil // ID3v2 has 'ID3' in first 3 bytes
	}
	return 0, errors.New("ID3 Tags : None found")
}
