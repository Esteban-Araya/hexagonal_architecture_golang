package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"os"
)

func Encrypt(plaintext []byte) (string, error) {
	c, err := aes.NewCipher([]byte(os.Getenv("ENCRYPTION_KEY")))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	bb := gcm.Seal(nonce, nonce, plaintext, nil)
	return toBase64(bb), nil
}

func Decrypt(bb string) ([]byte, error) {
	ciphertext, err := FromBase64(bb)
	if err != nil {
		return nil, err
	}

	c, err := aes.NewCipher([]byte(os.Getenv("ENCRYPTION_KEY")))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

func toBase64(plaintext []byte) string {
	return base64.RawStdEncoding.EncodeToString(plaintext)
}

func FromBase64(ciphertext string) ([]byte, error) {
	return base64.RawStdEncoding.DecodeString(ciphertext)
}
