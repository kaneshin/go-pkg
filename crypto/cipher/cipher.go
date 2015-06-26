package cipher

type Cipher interface {
	PlainText() string
	CipherText() string
	PlainData() []byte
	CipherData() []byte
}
