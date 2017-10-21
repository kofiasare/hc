package checkout

import "fmt"

type invoice struct {
	Items       map[string]*item `json:"items,omitempty"`
	Taxes       map[string]*tax  `json:"taxes,omitempty"`
	TotalAmount float64          `json:"total_amount,omitempty"`
	Description string           `json:"description,omitempty"`
}

// AddItem Is used for adding invoice items to the checkout invoice page.
func (i *invoice) AddItem(name string, quantity int, unitPrice string, totalPrice string, description string) {
	i.Items[fmt.Sprintf("item_%d", len(i.Items))] = &item{
		Name:        name,
		Quantity:    quantity,
		UnitPrice:   unitPrice,
		TotalPrice:  totalPrice,
		Description: description,
	}
}

type item struct {
	Name        string `json:"name,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
	UnitPrice   string `json:"unit_price,omitempty"`
	TotalPrice  string `json:"total_price,omitempty"`
	Description string `json:"description,omitempty"`
}

type tax struct {
	Name   string  `json:"name,omitempty"`
	Amount float64 `json:"amount,omitempty"`
}

func (i *invoice) AddTax(name string, amount float64) {
	i.Taxes[fmt.Sprintf("tax_%d", len(i.Taxes))] = &tax{
		Name:   name,
		Amount: amount,
	}
}
