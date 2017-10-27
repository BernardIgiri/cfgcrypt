package encryption

import (
	"bytes"
	"errors"
)

func Pad(buf []byte, size int) ([]byte, error) {
	bufLen := len(buf)
	padLen := size - bufLen%size
	padText := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(buf, padText...), nil
}

func Unpad(padded []byte, size int) ([]byte, error) {
	if len(padded)%size != 0 {
		return nil, errors.New("Padded value wasn't in correct size.")
	}
	paddedLen := len(padded)
	padLen := int(padded[paddedLen-1])
	bufLen := paddedLen - padLen
	return padded[:bufLen], nil
}
