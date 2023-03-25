package cdeklib

type Client struct {
	Token      string
	IsTestMode bool
	ApiURL     string
}

func NewClient(token string, isTestMode bool, apiURL string) *Client {
	return &Client{Token: token, IsTestMode: isTestMode, ApiURL: apiURL}
}
