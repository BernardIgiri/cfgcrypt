package encryption_test

import (
	"testing"

	"github.com/bernardigiri/cfgcrypt/encryption"
	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	key := []byte("1234567890123456")
	const expected = 64
	const data = "Hello bob, how are you?"

	actual, err := encryption.Encrypt(key, data)

	assert.Nil(t, err)
	assert.Equal(t, expected, len(actual))
}

func TestDecrypt(t *testing.T) {
	key := []byte("1234567890123456")
	const data = "nDK0gO1ExTPa60VtFjUXjv40VOMBEBtlHtAVqSvLeb9/p92Amla3C8s+6sVMd+Cw"
	const expected = "Hello bob, how are you?"

	actual, err := encryption.Decrypt(key, data)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestEncryptDecrypt(t *testing.T) {
	key := []byte("1234567890123456")
	const expected = "Hello bob, how are you?"

	intermediate, err := encryption.Encrypt(key, expected)
	assert.Nil(t, err)

	actual, err := encryption.Decrypt(key, intermediate)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}
