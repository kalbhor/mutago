package mutago

import (
	"errors"
	"os"
)

const (
	SynchByteLen  = 7
	NormalByteLen = 8
	BytePerInt    = 4
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

func BitSet(fl, index byte) bool {
	/*
		Tells whether a bit is set or not
	*/

	return fl%(1<<index) != 0
}

func ByteInt(buff []byte, base uint) (i uint32, err error) {
	/*
		Converts synch safe int into normal int
	*/
	if len(buff) > BytePerInt {
		err = errors.New("Bytes : Invalid []byte length")
		return
	}

	for _, b := range buff {
		if base < NormalByteLen && b >= (1<<base) {
			err = errors.New("Int : Exceeded max bit")
			return
		}

		i = (i << base) | uint32(b)
	}
	return
}
