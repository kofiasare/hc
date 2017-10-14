# Hubtel-Checkout

Hubtel Checkout Api wrapper in Golang

```go
package main

import (
  "log"
  "os"

  "github.com/kofiasare/hubtel/checkout"
)

func main() {

  // get credentials
  clientID := os.Getenv("HUBTEL_ID")
  clientSecret := os.Getenv("HUBTEL_SECRET")

  // setup checkout with credentials
  c, err := checkout.New(clientID, clientSecret)

  if err != nil {
    log.Fatal(err)
  }

  // store
  c.Store.Name = "T Shirt Company"
  c.Store.Tagline = "Tagline of the online store"
  c.Store.Phone = "233244124660"
  c.Store.PostAddress = "Box 10770 Accra - Ghana"
  c.Store.LogoURL = "https://company-logo-final.png"
  c.Store.WebsiteURL = "https://company.com"

  // invoice
  i := c.Invoice
  i.Description = "Description of the invoice here"

  // add invoice items
  i.AddItem("T Shirt", 2, "35.0", "70.0", "Order of 2 Shirts")
  i.AddItem("Polo Shirt", 1, "35.0", "70.0", "Order of 1 Polo Shirt")
  i.AddItem("Old Navy Jeans", 1, "25.0", "70.0", "Order of 1 Old Navy Jeans")

  // invoice taxes
  i.AddTax("Tax on T Shirt", 0.50)
  i.TotalAmount = 120.0

  // invoice actions
  i.SetCancelURL("http://company.com")
  i.SetReturnURL("http://company.com")

  // custom data
  i.AddCustomData("email", "kofi@gmail.com")
  _ = i.GetCustomData("email")

}

```