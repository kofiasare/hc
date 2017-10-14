package checkout

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

type checkout struct {
	auth    string
	Store   *store
	Invoice *invoice
}

// New access credentials for checkout endpoints
func New(clientID string, clientSecret string) (*checkout, error) {
	var ak string
	var err error

	if ak, err = genAuthKey(clientID, clientSecret); err != nil {
		return nil, err
	}

	c := &checkout{
		auth:  ak,
		Store: new(store),
		Invoice: &invoice{
			items:      make(map[string]*item),
			taxes:      make(map[string]*tax),
			customData: make(map[string]string),
			actions:    new(actions),
		},
	}

	return c, nil
}

func genAuthKey(i, s string) (string, error) {

	switch {
	case i == "" && s == "":
		return "", errors.New("clientID and ClientSecret are required to setup hubtel's checkout")
	case i == "":
		return "", errors.New("error: clientID is required to setup hubtel's checkout")
	case s == "":
		return "", errors.New("error: clientSecret is required to setup hubtel's checkout")
	}

	m := strings.Join([]string{i, s}, ":")
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(m))), nil
}

func genRequestBody() {

}
