package v2

import (
	"io"
)

type Frame struct {
	ID    string // 4 characters
	Size  uint32 // 4 bytes
	Flags []byte // 2 bytes
	/*
	   id3.org/id3v2.3.0
	   3.3.1 Frame Header Flags
	*/
	//    Encoding byte
	Info string
}

type Albumart struct {
	TextEncoding byte
	Mime         string
	PictureType  byte
	Description  string
	Data         []byte
}

func ParseFrame(reader io.ReadSeeker) *Frame {

	header := make([]byte, FrameHeaderSize)
	io.ReadFull(reader, header)

	frame := &Frame{
		ID:    string(header[:4]),
		Flags: header[8:],
	}

	size, err := BytesToInt(header[4:8], SynchIntLen)
	if err != nil {
		return nil
	}

	frame.Size = size

	info := make([]byte, frame.Size)

	io.ReadFull(reader, info)
	frame.Info = string(info)

	return frame
}

func ParseAlbumart(info string) *Albumart {
	/*
	   Returns the image in byte array, the image type and error
	*/

	var null byte = 0
	data := []byte(info)

	albumart := &Albumart{}
	albumart.TextEncoding = data[0] // First byte signifies text encoding

	for i, val := range data {
		if val == null {
			albumart.Mime = string(data[:i])
			data = data[i+1:]
			break
		}
	}

	albumart.PictureType = data[0]

	for i, val := range data {
		if val == null {
			albumart.Description = string(data[:i])
			data = data[i+1:]
			break
		}
	}

	for data[0] == null {
		data = data[1:]
	}

	albumart.Data = data

	return albumart
}
