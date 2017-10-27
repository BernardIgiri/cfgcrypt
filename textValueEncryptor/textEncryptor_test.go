package textValueEncryptor_test

import (
	"testing"

	"github.com/bernardigiri/cfgcrypt/textValueEncryptor"
	"github.com/stretchr/testify/assert"
)

func TestEncryptSubStrings(t *testing.T) {
	plaintext := "Hello #{{bob}}# is your daughter #{{sue}}# home?"
	prefix := "#{{"
	postfix := "}}#"
	key := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6}

	encrypted, err := textValueEncryptor.EncryptSubStrings(plaintext, prefix, postfix, key)

	assert.Nil(t, err)
	assert.NotEqual(t, plaintext, encrypted)
	assert.Contains(t, encrypted, "Hello ")
	assert.Contains(t, encrypted, " is your daughter ")
	assert.Contains(t, encrypted, " home?")
	assert.NotContains(t, encrypted, prefix)
	assert.NotContains(t, encrypted, postfix)
	assert.True(t, len(plaintext) < len(encrypted), "Encypted text must be longer")
}
