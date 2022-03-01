package v2

import (
	"io"
)

type header struct {
	Identifier   []byte // 3 bytes indicating ID3
	Version      []byte // 2 bytes
	Flags        byte   // 1 byte
	Unsynch      bool   // bit 7 of flags
	Experimental bool   // bit 5 of flags
	Extended     bool   // bit 6 of flags
	Size         uint32 // Size of Tag (Excluding header)
}

type extendedHeader struct {
	Size    uint32 // Size of extended tag (excluding header)
	Flags   []byte // 2 bytes
	Padding []byte // 4 bytes
	CRC     bool   // first bit of flags
}

func parseHeader(reader io.ReadSeeker) *header {
	data := make([]byte, HeaderSize)
	io.ReadFull(reader, data)

	header := &header{
		Identifier: data[:3],
		Version:    data[3:5],
		Flags:      data[5],
	}

	size, err := BytesToInt(data[6:], SynchIntLen)
	if err != nil {
		return nil
	}

	header.Size = size
	header.Unsynch = BitSet(header.Flags, 7)
	header.Extended = BitSet(header.Flags, 6)
	header.Experimental = BitSet(header.Flags, 5)

	return header
}
