# Hubtel's Checkout

API wrapper in Golang

## Installation

```go get -u github.com/kofiasare/hubtel/checkout```

## Usage

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kofiasare/hubtel/checkout"
)

func main() {

  // get credentials
  clientID := os.Getenv("HUBTEL_ID")
  clientSecret := os.Getenv("HUBTEL_SECRET")

  // setup checkout with credentials
  c, err := checkout.Setup(clientID, clientSecret)
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
  c.Invoice.AddItem("T Shirt", 2, "35.0", "70.0", "Order of 2 Shirts")
  c.Invoice.AddItem("Polo Shirt", 1, "35.0", "35.0", "Order of  Polo Shirt")
  c.Invoice.AddItem("Old Navy Jeans", 1, "25.0", "25.0", "Order of 1 Old Navy Jeans")
  c.Invoice.AddTax("Tax on T Shirt", 0.50)
  c.Invoice.TotalAmount = 130.00

  // checkout custom data
  c.CustomData.Add("email", "kofi@gmail.com")

  // checkout actions
  c.Actions.CancelURL = "http://company.com"
  c.Actions.ReturnURL = "http://company.com"

  // create checkout
  r, err := c.Create()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(string(r))
}

```

# Todo

3. Tests all methods and functions

4. Think through after create response handling

5. write demo app with material design bootstrap