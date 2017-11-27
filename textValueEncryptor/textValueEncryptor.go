package textValueEncryptor

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/bernardigiri/cfgcrypt/encryption"
)

const KEY_FILE_EXTENSION = ".key"

const DEFAULT_KEY_SIZE = 16

const MAX_FILE_SIZE = 1E7

func generateKey(textFilePath string, force bool) (key []byte, err error) {
	keyPath := textFilePath + KEY_FILE_EXTENSION
	if _, statErr := os.Stat(keyPath); !os.IsNotExist(statErr) && !force {
		msg := fmt.Sprintf("Key file \"%s\" already exists", keyPath)
		err = errors.New(msg)
		return
	}
	key = make([]byte, DEFAULT_KEY_SIZE)
	if _, err = io.ReadFull(rand.Reader, key); err != nil {
		return
	}
	err = ioutil.WriteFile(keyPath, key, 0600)
	return
}

func isPowerOf2(n int) bool {
	return (n&(n-1)) == 0 && n > 0
}

func EncryptTextFile(textFilePath, prefix, postfix, encodedKey string, force bool) (err error) {
	var key []byte
	if encodedKey == "" {
		key, err = generateKey(textFilePath, force)
	} else {
		key, err = base64.StdEncoding.DecodeString(encodedKey)
	}
	if err != nil {
		return
	}
	stats, err := os.Stat(textFilePath)
	if err != nil {
		return
	}
	if stats.Size() > MAX_FILE_SIZE {
		msg := fmt.Sprintf("File size exceeds the maximum allowed size of %d bytes", MAX_FILE_SIZE)
		err = errors.New(msg)
		return
	}
	data, err := ioutil.ReadFile(textFilePath)
	if err != nil {
		return
	}
	encrypted, err := EncryptSubStrings(string(data), prefix, postfix, key)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(textFilePath, []byte(encrypted), 0)
	return
}

func EncryptSubStrings(plaintext, prefix, postfix string, key []byte) (encrypted string, err error) {
	if !isPowerOf2(len(key)) {
		err = errors.New("Invalid key size")
		return
	}
	if len(postfix) == 0 || len(prefix) == 0 {
		err = errors.New("Invalid prefix/postfix")
		return
	}
	position := 0
	output := []string{}
	for position < len(plaintext) {
		prefixIdx := strings.Index(plaintext[position:], prefix)
		if prefixIdx < 0 {
			break
		} else {
			prefixIdx += position
		}
		postfixIdx := strings.Index(plaintext[prefixIdx:], postfix)
		if postfixIdx < 0 {
			break
		} else {
			postfixIdx += prefixIdx
		}
		var encryptedValue string
		encryptedValue, err = encryption.Encrypt(key, plaintext[prefixIdx+len(prefix):postfixIdx])
		if err != nil {
			return
		}
		output = append(output, plaintext[position:prefixIdx], encryptedValue)
		position = postfixIdx + len(postfix)

	}
	output = append(output, plaintext[position:])
	encrypted = strings.Join(output, "")
	return
}
