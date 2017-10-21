package checkout

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type checkout struct {
	auth       string
	Store      *store
	Invoice    *invoice
	Actions    *actions
	CustomData *customData
}

type checkoutRequest struct {
	Invoice    *invoice    `json:"invoice,omitempty"`
	Store      *store      `json:"store,omitempty"`
	Actions    *actions    `json:"actions,omitempty"`
	CustomData *customData `json:"custom_data,omitempty"`
}

// Setup initialize checkout with hubtel checkout
//  api account credeclient
func Setup(clientID string, clientSecret string) (*checkout, error) {
	var ak string
	var err error

	if ak, err = genAuthKey(clientID, clientSecret); err != nil {
		return nil, err
	}

	c := &checkout{
		auth:    ak,
		Store:   new(store),
		Actions: new(actions),
		Invoice: &invoice{
			Items: make(map[string]*item),
			Taxes: make(map[string]*tax),
		},
		CustomData: &customData{
			CustomData: make(map[string]interface{}),
		},
	}
	return c, nil
}

const (
	url = "https://api.hubtel.com/v1/merchantaccount/onlinecheckout/invoice/create"
)

func (c *checkout) Create() ([]byte, error) {

	body, err := c.genRequestBody()
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Add("Authorization", c.auth)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *checkout) genRequestBody() ([]byte, error) {
	body := &checkoutRequest{
		Invoice:    c.Invoice,
		Store:      c.Store,
		Actions:    c.Actions,
		CustomData: c.CustomData,
	}

	return json.Marshal(body)
}

func genAuthKey(clientID, clientSecret string) (string, error) {

	switch {
	case clientID == "" && clientSecret == "":
		return "", errors.New("clientID and clientSecret are required to setup hubtel's checkout")
	case clientID == "":
		return "", errors.New("error: clientID is required to setup hubtel's checkout")
	case clientSecret == "":
		return "", errors.New("error: clientSecret is required to setup hubtel's checkout")
	}

	m := strings.Join([]string{clientID, clientSecret}, ":")
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(m))), nil
}
