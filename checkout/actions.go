package checkout

type actions struct {
	CancelURL string `json:"cancel_url,omitempty"`
	ReturnURL string `json:"return_url,omitempty"`
}
