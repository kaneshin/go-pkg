package cipher

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	assert := assert.New(t)

	key := []byte("example key 1234")
	plaintext := "0000000012345678"
	ciphertext := "ebea9d1f37bf1d6f34af02e9e2a552a03a8026b8597aa185baf52f3a432eef4f"

	for i := 0; i < 100; i++ {
		// Encrypt -> Decrypt
		plaindata := []byte(plaintext)
		encrypter := NewCBCEncrypter(key, plaindata)
		assert.Equal(plaindata, encrypter.PlainData())

		// Decrypt -> Encrypt
		cipherdata := encrypter.CipherData()
		decrypter := NewCBCDecrypter(key, cipherdata)
		assert.Equal(plaindata, decrypter.PlainData())
	}

	for i := 0; i < 100; i++ {
		// Decrypt -> Encrypt
		plaindata := []byte(plaintext)
		cipherdata, _ := hex.DecodeString(ciphertext)
		decrypter := NewCBCDecrypter(key, cipherdata)
		assert.Equal(plaindata, decrypter.PlainData())

		// Encrypt -> Decrypt
		encrypter := NewCBCEncrypter(key, plaindata)
		assert.Equal(plaindata, encrypter.PlainData())

		// Decrypt -> Encrypt
		cipherdata = encrypter.CipherData()
		decrypter = NewCBCDecrypter(key, cipherdata)
		assert.Equal(plaindata, decrypter.PlainData())
	}

}
