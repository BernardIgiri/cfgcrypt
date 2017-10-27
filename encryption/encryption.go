package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

func Decrypt(key []byte, encodedCipherText string) (plaintext string, err error) {
	// DECODE
	ciphertext, err := base64.StdEncoding.DecodeString(encodedCipherText)
	if err != nil {
		return
	}
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	if len(ciphertext) < aes.BlockSize {
		err = errors.New("ciphertext too short")
		return
	}
	// GET IV
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	if len(ciphertext)%aes.BlockSize != 0 {
		err = errors.New("ciphertext is not a multiple of the block size")
		return
	}
	// DECRYPT
	mode := cipher.NewCBCDecrypter(blockCipher, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	// UNPAD
	plaintextBytes, err := Unpad(ciphertext, aes.BlockSize)
	plaintext = string(plaintextBytes)
	return
}

func Encrypt(key []byte, plaintext string) (encodedCipherText string, err error) {
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	// PAD
	paddedPlaintext, err := Pad([]byte(plaintext), aes.BlockSize)
	if err != nil {
		return
	}
	// MAKE IV
	ciphertext := make([]byte, aes.BlockSize+len(paddedPlaintext))
	iv := ciphertext[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return
	}
	// ENCRYPT
	mode := cipher.NewCBCEncrypter(blockCipher, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], paddedPlaintext)
	// ENCODE
	encodedCipherText = base64.StdEncoding.EncodeToString(ciphertext)
	return
}
