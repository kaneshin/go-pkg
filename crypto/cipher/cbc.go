package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

/*
  mi: message vector
   k: Secret Key
   E: Block cipher

           m1                 m2                    mx
           |                  |                     |
           v                  v                     v
         _____              _____                 _____
        |     |            |     |               |     |
  0 --> |  +  |   +------> |  +  |   +---...---> |  +  |
        |_____|   |        |_____|   |           |_____|
                  |                  |
           |      |           |      |              |
           v      |           v      |              v
         _____    |         _____    |            _____
        |     |   |        |     |   |           |     |
  k --> |  E  |   |  k --> |  E  |   |     k --> |  E  |
        |_____|   |        |_____|   |           |_____|
           |      |           |      |              |
           +------+           +------+              v
                                                 result
*/

// CBC means cipher block chaining
type CBC struct {
	key        []byte
	plainData  []byte
	cipherData []byte
}

// NewCBCEncrypter returns initialized CBC
func NewCBCEncrypter(key, plainData []byte) *CBC {
	c := CBC{
		key:       key,
		plainData: plainData[:],
	}
	c.cipherData = c.Encrypt()
	return &c
}

// NewCBCDecrypter returns initialized CBC
func NewCBCDecrypter(key, cipherData []byte) *CBC {
	c := CBC{
		key:        key,
		cipherData: cipherData[:],
	}
	c.plainData = c.Decrypt()
	return &c
}

// PlainText returns plain text of CBC
func (c CBC) PlainText() string {
	return string(c.plainData)
}

// CipherText returns cipher text of CBC
func (c CBC) CipherText() string {
	return string(c.cipherData)
}

// PlainData returns plain bytes of CBC
func (c CBC) PlainData() []byte {
	return c.plainData
}

// CipherData returns cipher bytes of CBC
func (c CBC) CipherData() []byte {
	return c.cipherData
}

// Encrypt returns encrypted text
func (c CBC) Encrypt() []byte {
	key := c.key
	plainData := c.plainData[:]

	if len(plainData)%aes.BlockSize != 0 {
		panic("plainData is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	cipherData := make([]byte, aes.BlockSize+len(plainData))
	iv := cipherData[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherData[aes.BlockSize:], plainData)

	return cipherData
}

// Encrypt returns decrypted text
func (c CBC) Decrypt() []byte {
	key := c.key
	cipherData := c.cipherData[:]

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(cipherData) < aes.BlockSize {
		panic("cipherData too short")
	}
	iv := cipherData[:aes.BlockSize]
	cipherData = cipherData[aes.BlockSize:]

	if len(cipherData)%aes.BlockSize != 0 {
		panic("cipherData is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(cipherData, cipherData)

	return cipherData
}
