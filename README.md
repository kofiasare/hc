# Hubtel Merchant Checkout

API wrapper in Golang.

The [hubtel merchant account checkout API](https://developers.hubtel.com/documentations/online-checkout-api) allows merchants to accept online payment for goods and services using mobile money and credit/debit cards.

see the [docs](https://godoc.org/github.com/kofiasare/hc) for details

## Download

```$ go get -u github.com/kofiasare/hc```

## Usage

```go

package main

import (
  "fmt"
  "log"
  "os"

  "github.com/kofiasare/hc"
)

func main() {

  // get credentials
  clientID := os.Getenv("HUBTEL_ID")
  clientSecret := os.Getenv("HUBTEL_SECRET")

  // setup checkout with credentials
  c, err := hc.Setup(clientID, clientSecret)
  if err != nil {
    log.Fatal(err)
  }

  // checkout store
  c.Store.Name = "T Shirt Company"
  c.Store.Tagline = "Tagline of the online store"
  c.Store.Phone = "233244124660"
  c.Store.PostAddress = "Box 10770 Accra - Ghana"
  c.Store.LogoURL = "https://company-logo-final.png"
  c.Store.WebsiteURL = "https://company.com"

  // checkout invoice
  c.Invoice.Description = "Invoice Description"
  c.Invoice.AddItem("T Shirt", 2, 35.0, 70.0, "Order of 2 Shirts")
  c.Invoice.AddItem("Polo Shirt", 1, 35.0, 35.0, "Order of  Polo Shirt")
  c.Invoice.AddItem("Old Navy Jeans", 1, 25.0, 25.0, "Order of 1 Old Navy Jeans")
  c.Invoice.AddTax("Tax on T Shirt", 0.50)
  c.Invoice.TotalAmount = 130.50

  // checkout custom data
  c.CustomData.Add("email", "kofi@gmail.com")

  // checkout actions
  c.Actions.CancelURL = "http://company.com"
  c.Actions.ReturnURL = "http://company.com"

  // create checkout invoice
  r, err := c.Create()
  if err != nil {
    log.Fatal(err)
  }

  if r.ResponseCode == "00" {

    fmt.Printf("Redirect URL: %s\n", r.ResponseText)
    fmt.Printf("Token: %s\n", r.Token)

    // retrieve checkout invoice status
    r, err = c.Status(r.Token)
    if err != nil {
      log.Fatal(err)
    }

    if r.ResponseCode == "00" {
      fmt.Printf("Checkout Invoice Status is: %s\n", r.Status)
    }

  } else {
      fmt.Println(r.ResponseText)
  }
}

```
