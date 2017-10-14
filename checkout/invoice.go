package checkout

import (
	"fmt"
)

type invoice struct {
	items       map[string]*item
	taxes       map[string]*tax
	customData  map[string]string
	TotalAmount float64
	Description string
	*actions
}

type item struct {
	name        string
	quantity    int
	unitPrice   string
	totalPrice  string
	description string
}

type tax struct {
	name   string
	amount float64
}

type actions struct {
	cancelURL string
	returnURL string
}

func (i *invoice) AddItem(name string, quantity int, unitPrice string, totalPrice string, description string) {
	i.items[fmt.Sprintf("item_%d", len(i.items))] = &item{
		name:        name,
		quantity:    quantity,
		unitPrice:   unitPrice,
		totalPrice:  totalPrice,
		description: description,
	}
}

func (i *invoice) AddTax(name string, amount float64) {
	i.taxes[fmt.Sprintf("tax_%d", len(i.taxes))] = &tax{
		name:   name,
		amount: amount,
	}
}

func (i *invoice) SetCancelURL(url string) {
	i.actions.cancelURL = url
}

func (i *invoice) SetReturnURL(url string) {
	i.actions.returnURL = url
}

func (i *invoice) AddCustomData(key, value string) {
	i.customData[key] = value
}

func (i *invoice) GetCustomData(key string) string {
	return i.customData[key]
}
