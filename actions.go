package hc

// Actions that a customer can take on a checkout page.
type Actions struct {

	// The URL customers should be redirected to after checkout.
	// Has to be specified to enable a redirect to your website
	// after checkout.
	CancelURL string `json:"cancel_url,omitempty"`

	// The URL a customer should be redirected to when an invoice
	// is cancelled.
	ReturnURL string `json:"return_url,omitempty"`
}
