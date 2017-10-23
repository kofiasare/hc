package hc

import (
	"encoding/json"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnconfiguredClientID(t *testing.T) {
	assert := assert.New(t)
	_, err := Setup("", "efgh")
	assert.Error(err)
	assert.EqualError(err, "configure client_id")

}

func TestUnconfiguredClientSecret(t *testing.T) {
	assert := assert.New(t)
	_, err := Setup("abcd", "")
	assert.Error(err)
	assert.EqualError(err, "configure client_secret")
}

func TestUnconfiguredClientIDAndSecret(t *testing.T) {
	assert := assert.New(t)
	_, err := Setup("", "")
	assert.Error(err)
	assert.EqualError(err, "configure client_id and client_secret")
}

func TestAuthKeyGenerator(t *testing.T) {
	whiteSpace := " "
	assert := assert.New(t)
	ak, _ := genAuthKey("abcd", "efgh")
	assert.True(strings.HasPrefix(ak, "Basic"))
	assert.True(strings.Count(ak, whiteSpace) == 1, "contains < 1 or > 1 whitespace")
	assert.True(ak != "BASIC yWJjZDplZmdo")
}

func TestRequestBodyGenerator(t *testing.T) {
	assert := assert.New(t)
	c, err := Setup("abcd", "efgh")

	assert.NoError(err)
	body, err := c.genRequestBody()

	if assert.NoError(err) {
		assert.IsType([]byte{}, body)

		eb, _ := json.Marshal(&checkoutRequest{
			Invoice:    c.Invoice,
			Store:      c.Store,
			Actions:    c.Actions,
			CustomData: c.CustomData,
		})
		assert.Exactly(string(eb), string(body))
	}

}

func TestCheckoutApiWrapper(t *testing.T) {

	clientID := os.Getenv("HUBTEL_ID")
	clientSecret := os.Getenv("HUBTEL_SECRET")
	assert := assert.New(t)

	c, err := Setup(clientID, clientSecret)
	assert.NoError(err)
	if assert.IsType(&Checkout{}, c) {

		// checkout store
		checkoutStore := &Store{
			Name:        "T Shirt Company",
			Tagline:     "Tagline of the online store",
			Phone:       "233244124660",
			PostAddress: "Box 10770 Accra - Ghana",
			LogoURL:     "https://company-logo-final.png",
			WebsiteURL:  "https://company.com",
		}

		c.Store = checkoutStore

		// checkout actions
		c.Actions.CancelURL = "http://company.com"
		c.Actions.ReturnURL = "http://company.com"

		// checkout invoice
		c.Invoice.Description = "Invoice Description"
		c.Invoice.AddItem("T-Shirt", 2, 35.0, 70.0, "Order of 2 T-Shirts")
		c.Invoice.AddItem("Snikers", 1, 50.0, 50.0, "Order of 1 Old Navy Jeans")
		c.Invoice.AddTax("Tax on purchase", 0.50)
		c.Invoice.TotalAmount = 120.50

		// checkout custom data
		c.CustomData.Add("email", "kofi@gmail.com")

		r, err := c.Create()
		if assert.NoError(err) {

			if r.ResponseCode == "00" {
				open(r.ResponseText)

				r, err = c.Status(r.Token)
				if assert.NoError(err) {
					if r.ResponseCode == "00" {
						assert.Equal(c.Invoice, r.Invoice)
						assert.Equal(c.CustomData, r.CustomData)
					}
				}
			}

		}
	}
}

func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
