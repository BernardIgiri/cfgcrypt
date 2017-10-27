package textValueEncryptor_test

import (
	"strings"
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
	assert.Zero(t, strings.Index(encrypted, "Hello "))
	assert.Contains(t, encrypted, " is your daughter ")
	assert.Contains(t, encrypted, " home?")
	assert.Equal(t, strings.Index(encrypted, " home?"), len(encrypted)-len(" home?"))
	assert.NotContains(t, encrypted, prefix)
	assert.NotContains(t, encrypted, postfix)
	assert.True(t, len(plaintext) < len(encrypted), "Encypted text must be longer")
}

func TestEncryptSubStringsWithExtraPostfixsA(t *testing.T) {
	plaintext := "Hello #{{bob}}# is your daughter #{{sue}}# home?}}#"
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
	assert.Contains(t, encrypted, postfix)
	assert.True(t, len(plaintext) < len(encrypted), "Encypted text must be longer")
}

func TestEncryptSubStringsWithExtraPostfixsB(t *testing.T) {
	plaintext := "Hello #{{bob}}# }}# is your daughter #{{sue}}# home?"
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
	assert.Contains(t, encrypted, postfix)
	assert.True(t, len(plaintext) < len(encrypted), "Encypted text must be longer")
}

func TestEncryptSubStringsWithExtraPostfixsC(t *testing.T) {
	plaintext := "}}#Hello #{{bob}}# }}# is your daughter #{{sue}}# home?"
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
	assert.Contains(t, encrypted, postfix)
	assert.True(t, len(plaintext) < len(encrypted), "Encypted text must be longer")
}

func TestEncryptSubStringsWithExtraPrefix(t *testing.T) {
	plaintext := "Hello #{{bob}}# is your daughter #{{sue}}# home?#{{"
	prefix := "#{{"
	postfix := "}}#"
	key := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6}

	encrypted, err := textValueEncryptor.EncryptSubStrings(plaintext, prefix, postfix, key)

	assert.Nil(t, err)
	assert.NotEqual(t, plaintext, encrypted)
	assert.Contains(t, encrypted, "Hello ")
	assert.Contains(t, encrypted, " is your daughter ")
	assert.Contains(t, encrypted, " home?")
	assert.Contains(t, encrypted, prefix)
	assert.NotContains(t, encrypted, postfix)
	assert.True(t, len(plaintext) < len(encrypted), "Encypted text must be longer")
}

func TestEncryptSubStringsAtEnds(t *testing.T) {
	plaintext := "#{{secret1}}#Hello #{{bob}}# is your daughter #{{sue}}# home?#{{secret2}}#"
	prefix := "#{{"
	postfix := "}}#"
	key := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6}

	encrypted, err := textValueEncryptor.EncryptSubStrings(plaintext, prefix, postfix, key)

	assert.Nil(t, err)
	assert.NotEqual(t, plaintext, encrypted)
	assert.Contains(t, encrypted, "Hello ")
	assert.NotZero(t, encrypted, strings.Index(encrypted, "Hello "))
	assert.Contains(t, encrypted, " is your daughter ")
	assert.Contains(t, encrypted, " home?")
	assert.NotContains(t, encrypted, prefix)
	assert.NotContains(t, encrypted, postfix)
	assert.True(t, len(plaintext) < len(encrypted), "Encypted text must be longer")
}
