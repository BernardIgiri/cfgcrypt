package encryption_test

import (
	"testing"

	"github.com/bernardigiri/AESYamlEncryptor/encryption"
	"github.com/stretchr/testify/assert"
)

func TestPad(t *testing.T) {
	data := []byte("Lorem ipsum")
	dataLen := len(data)
	for blockSize := 16; blockSize <= 128; blockSize *= 2 {
		actual, err := encryption.Pad(data, blockSize)
		assert.Nil(t, err)
		assert.Zero(t, len(actual)%blockSize)
		assert.NotZero(t, len(data)%blockSize)
		assert.Equal(t, int(actual[dataLen]), blockSize-dataLen)
	}
}

func TestPad16(t *testing.T) {
	expected := []byte{
		'B', 'B', 'B', 'B',
		'B', 'B', 'B', 'B',
		'B', 'B', 'B', 5,
		5, 5, 5, 5,
	}
	data := []byte("BBBBBBBBBBB")
	actual, err := encryption.Pad(data, 16)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestPad32(t *testing.T) {
	expected := []byte{
		'B', 'B', 'B', 'B',
		'B', 'B', 'B', 'B',
		'B', 'B', 'B', 21,
		21, 21, 21, 21,
		21, 21, 21, 21,
		21, 21, 21, 21,
		21, 21, 21, 21,
		21, 21, 21, 21,
	}
	data := []byte("BBBBBBBBBBB")
	actual, err := encryption.Pad(data, 32)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestUnpad16(t *testing.T) {
	data := []byte{
		'B', 'B', 'B', 'B',
		'B', 'B', 'B', 'B',
		'B', 'B', 'B', 5,
		5, 5, 5, 5,
	}
	expected := []byte("BBBBBBBBBBB")
	actual, err := encryption.Unpad(data, 16)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestUnpad32(t *testing.T) {
	data := []byte{
		'B', 'B', 'B', 'B',
		'B', 'B', 'B', 'B',
		'B', 'B', 'B', 21,
		21, 21, 21, 21,
		21, 21, 21, 21,
		21, 21, 21, 21,
		21, 21, 21, 21,
		21, 21, 21, 21,
	}
	expected := []byte("BBBBBBBBBBB")
	actual, err := encryption.Unpad(data, 32)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
