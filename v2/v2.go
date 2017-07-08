package v2

import (
	"io"

	"github.com/makebyte/mutago"
)

const (
	HeaderSize  = 10
	SynchIntLen = 7
)

type Header struct {
	Identifier   []byte // 3 bytes indicating ID3
	Version      []byte // 2 bytes
	Flags        byte   // 1 byte
	Unsynch      bool   // bit 7 of flags
	Experimental bool   // bit 5 of flags
	Extended     bool   // bit 6 of flags
	Size         uint32 // Size of Tag (Excluding header)
}

type ExtendedHeader struct {
	Size    uint32 // Size of extended tag (excluding header)
	Flags   []byte // 2 bytes
	Padding []byte // 4 bytes
	CRC     bool   // first bit of flags
}

func ParseHeader(reader io.ReadSeeker) *Header {

	data := make([]byte, HeaderSize)
	io.ReadFull(reader, data)

	header := &Header{
		Identifier: data[:3],
		Version:    data[3:5],
		Flags:      data[5],
	}

	size, err := mutago.ByteInt(data[6:], SynchIntLen)
	if err != nil {
		return nil
	}

	header.Size = size
	header.Unsynch = mutago.BitSet(header.Flags, 7)
	header.Extended = mutago.BitSet(header.Flags, 6)
	header.Experimental = mutago.BitSet(header.Flags, 5)

	return header
}
