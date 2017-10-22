// Package hc is an api wrapper to hubtel's merchant account checkout API
// allows merchants to accept online payment for goods and services using
// mobile money and credit/debit cards
package hc

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

const (
	// CreateInvoiceURL endpoit for creating checkout invoice
	CreateInvoiceURL = "https://api.hubtel.com/v1/merchantaccount/onlinecheckout/invoice/create"
)

// Checkout represents hubtel online
// checkout page
type Checkout struct {
	auth       string
	Invoice    *Invoice
	Store      *Store
	Actions    *Actions
	CustomData *CustomData
}

type checkoutRequest struct {
	Invoice    *Invoice    `json:"invoice,omitempty"`
	Store      *Store      `json:"store,omitempty"`
	Actions    *Actions    `json:"actions,omitempty"`
	CustomData *CustomData `json:"custom_data,omitempty"`
}

type checkoutResponse struct {
	ResponseCode string `json:"response_code,omitempty"`
	ResponseText string `json:"response_text,omitempty"`
	Description  string `json:"description,omitempty"`
	Token        string `json:"token,omitempty"`
}

// Setup checkout with hubtel checkout
// api account credentials
func Setup(clientID, clientSecret string) (*Checkout, error) {
	var ak string
	var err error

	if ak, err = genAuthKey(clientID, clientSecret); err != nil {
		return nil, err
	}

	c := &Checkout{
		auth:    ak,
		Store:   new(Store),
		Actions: new(Actions),
		Invoice: &Invoice{
			Items: make(map[string]*Item),
			Taxes: make(map[string]*Tax),
		},
		CustomData: &CustomData{
			CustomData: make(map[string]interface{}),
		},
	}
	return c, nil
}

// Create online checkout invoice
func (c *Checkout) Create() (*checkoutResponse, error) {

	body, err := c.genRequestBody()
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest(http.MethodPost, CreateInvoiceURL, bytes.NewBuffer(body))
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

	cr := new(checkoutResponse)
	err = json.Unmarshal(body, cr)
	if err != nil {
		return nil, err
	}

	return cr, nil
}

func (c *Checkout) genRequestBody() ([]byte, error) {
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
