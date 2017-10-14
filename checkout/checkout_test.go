package checkout

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ci = os.Getenv("HUBTEL_ID")
var cs = os.Getenv("HUBTEL_SECRET")

func TestSetup(t *testing.T) {
	assert := assert.New(t)

	c, err := Setup(ci, cs)

	if assert.NoError(err) {

		assert.Equal("Basic a3lhYnhpd206YnFyaXhkZmg=", c.authKey)

	}
}
