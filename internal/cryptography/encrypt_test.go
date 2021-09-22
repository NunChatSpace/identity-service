package cryptography_test

import (
	"testing"

	"github.com/NunChatSpace/identity-service/internal/cryptography"
	"github.com/stretchr/testify/assert"
)

func Test_EncryptDecrypt(t *testing.T) {
	t.Run("test encrypt", func(t *testing.T) {
		password := "1234"
		ep, _ := cryptography.Encrypt(password)
		dp, err := cryptography.Decrypt(ep)

		assert.NoError(t, err)
		assert.Equal(t, password, dp)
	})
}
