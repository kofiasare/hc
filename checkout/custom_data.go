package checkout

type customData struct {
	CustomData map[string]interface{} `json:"custom_data,omitempty"`
}

func (cd *customData) Add(key string, value interface{}) {
	cd.CustomData[key] = value
}

func (cd *customData) Get(key string) interface{} {
	return cd.CustomData[key]
}
