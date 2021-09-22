package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func Encrypt(password string) (string, error) {
	bytes := []byte("ThisIsSecretKey@123456789asdfghb")
	key := hex.EncodeToString(bytes)
	usedKey, _ := hex.DecodeString(key)
	c, err := aes.NewCipher(usedKey)
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

	encryptedStr := gcm.Seal(nonce, nonce, []byte(password), nil)
	return fmt.Sprintf("%x", encryptedStr), nil
}
