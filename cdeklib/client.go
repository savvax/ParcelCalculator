package cdeklib

type Client struct {
	Token      string
	IsTestMode bool
	ApiUrl     string
}

func NewClient(token string, isTestMode bool, apiUrl string) *Client {
	return &Client{Token: token, IsTestMode: isTestMode, ApiUrl: apiUrl}
}
