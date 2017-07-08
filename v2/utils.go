package v2

import "errors"

func BitSet(fl, index byte) bool {
	/*
		Tells whether a bit is set or not
	*/

	return fl%(1<<index) != 0
}

func BytesToInt(buff []byte, base uint) (i uint32, err error) {
	/*
		Converts synch safe int into normal int
	*/
	if len(buff) > BytesPerInt {
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
