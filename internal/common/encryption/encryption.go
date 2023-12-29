package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/ugabiga/falcon/internal/config"
	"io"
)

type Encryption struct {
	encryptionKey string
}

func NewEncryption(cfg *config.Config) *Encryption {
	return &Encryption{
		encryptionKey: cfg.EncryptionKey,
	}
}

func (e Encryption) Encrypt(plainText string) (cipherText string, err error) {
	block, err := aes.NewCipher([]byte(e.encryptionKey))
	if err != nil {
		return "", err
	}

	plainBytes := []byte(plainText)
	cipherBytes := make([]byte, aes.BlockSize+len(plainBytes))
	iv := cipherBytes[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherBytes[aes.BlockSize:], plainBytes)

	return hex.EncodeToString(cipherBytes), nil
}

func (e Encryption) Decrypt(cipherText string) (plainText string, err error) {
	block, err := aes.NewCipher([]byte(e.encryptionKey))
	if err != nil {
		return "", err
	}

	cipherBytes, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	if len(cipherBytes) < aes.BlockSize {
		return "", fmt.Errorf("cipherText too short")
	}

	iv := cipherBytes[:aes.BlockSize]
	cipherBytes = cipherBytes[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherBytes, cipherBytes)

	return string(cipherBytes), nil
}
