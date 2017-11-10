package hc

// CustomData allows saving of custom data on
// the hubtel service which are persisted even after
// successful payment.
type CustomData struct {
	CustomData map[string]interface{} `json:"data,omitempty"`
}

// Add is used for adding custom data to the checkout invoice page.
func (c *CustomData) Add(key string, value interface{}) {
	c.CustomData[key] = value
}

// Get is used for retrieving a custom data.
func (c CustomData) Get(key string) interface{} {
	return c.CustomData[key]
}
