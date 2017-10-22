package hc

import "fmt"

// Invoice object specifies the list of items to display
// on the invoice page. When the invoice is created, your
// application receives an invoice token and a checkout URL
// for a customer to make payment.
type Invoice struct {
	Items       map[string]*Item `json:"items,omitempty"`
	Taxes       map[string]*Tax  `json:"taxes,omitempty"`
	TotalAmount float64          `json:"total_amount,omitempty"`
	Description string           `json:"description,omitempty"`
}

// AddItem Is used for adding invoice items to the checkout invoice page.
func (i *Invoice) AddItem(name string, quantity int, unitPrice string, totalPrice string, description string) {
	i.Items[fmt.Sprintf("item_%d", len(i.Items))] = &Item{
		Name:        name,
		Quantity:    quantity,
		UnitPrice:   unitPrice,
		TotalPrice:  totalPrice,
		Description: description,
	}
}

// Item specifies a single item on the checkout page.
type Item struct {
	Name        string `json:"name,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
	UnitPrice   string `json:"unit_price,omitempty"`
	TotalPrice  string `json:"total_price,omitempty"`
	Description string `json:"description,omitempty"`
}

// Tax specifies the tax items that apply to the invoice.
type Tax struct {
	Name   string  `json:"name,omitempty"`
	Amount float64 `json:"amount,omitempty"`
}

// AddTax is used for adding tax items to the checkout invoice page.
func (i *Invoice) AddTax(name string, amount float64) {
	i.Taxes[fmt.Sprintf("tax_%d", len(i.Taxes))] = &Tax{
		Name:   name,
		Amount: amount,
	}
}
