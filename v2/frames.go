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
