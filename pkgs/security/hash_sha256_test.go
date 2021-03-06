package security_test

import (
	"testing"

	"github.com/mrzack99s/mrz-sso-provider/pkgs/security"
	"github.com/stretchr/testify/assert"
)

func TestSumSha256(t *testing.T) {
	sha256 := "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"
	hash := security.SumSha256("test")
	assert.Equal(t, sha256, hash)
}
