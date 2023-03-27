package cdeklib

import "fmt"

type Client struct {
	Token            string
	IsTestMode       bool
	ApiURL           string
	ApiUrlTariffList string
}

func NewClient(isTestMode bool, apiURL string, apiUrlTariffList string, account string, securePassword string) *Client {
	var token, err = GetAccessToken(apiURL, account, securePassword)
	if err != nil {
		fmt.Println("Error getting access token:", err)
	}
	return &Client{Token: token, IsTestMode: isTestMode, ApiURL: apiURL, ApiUrlTariffList: apiUrlTariffList}
}
