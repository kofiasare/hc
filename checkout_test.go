package hc

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	clientID     = os.Getenv("HUBTEL_ID")
	clientSecret = os.Getenv("HUBTEL_SECRET")
)

func TestGenAuthKey(t *testing.T) {
	assert := assert.New(t)

	ak, err := genAuthKey(clientID, clientSecret)
	if assert.NoError(err) {
		assert.Equal("", ak)
	}
}
