package security_test

import (
	"testing"

	"github.com/mrzack99s/mrz-sso-provider/pkgs/security"
	"github.com/mrzack99s/mrz-sso-provider/pkgs/system"
	"github.com/stretchr/testify/assert"
)

func TestEncryption(t *testing.T) {
	system.ParseSystemConfig("../../config.yaml")
	plainText := "test"
	cipherText := security.Encrypt(plainText)
	plainTextFromCipherText := security.Decrypt([]byte(cipherText))
	assert.Equal(t, plainText, plainTextFromCipherText)
}
