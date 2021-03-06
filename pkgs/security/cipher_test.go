package security_test

import (
	"testing"

	"github.com/mrzack99s/mrz-sso-provider/pkgs/security"
	"github.com/mrzack99s/mrz-sso-provider/pkgs/system"
	"github.com/stretchr/testify/assert"
)

func TestEncryption(t *testing.T) {
	system.ParseSystemConfig("../../config.yaml")
	assert.Equal(t, "0d25ad0a9885624feeff524008607b24", system.SystemConfigVar.MRZ_SSO.Security.Salt)
	plainText := "test"
	cipherText := security.Encrypt(plainText)
	plainTextFromCipherText := security.Decrypt([]byte(cipherText))
	assert.Equal(t, plainText, plainTextFromCipherText)
}
