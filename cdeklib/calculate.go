package cdeklib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Calculate returns an array of tariffs based on the given fromLocation, toLocation, and size.
func (c *Client) Calculate(fromLocation, toLocation Location, size Size) (string, error) {
	requestBody := fmt.Sprintf(`
	{	
		"type": "%d",
    	"date": "%s",
    	"currency": "%d",
    	"lang": "%s",
		"from_location": {
			"code": "%s",
			"postal_code": "%s",
			"country_code": "%s",
			"city": "%s",
			"address": "%s"
		},
		"to_location": {
			"code": "%s",
			"postal_code": "%s",
			"country_code": "%s",
			"city": "%s",
			"address": "%s"
		},
		"packages": [{"weight": %d, "length": %d, "width": %d, "height": %d}]
	}
	`, CdekType, CdekDate, CdekCurrency, CdekLang, fromLocation.Code, fromLocation.PostalCode, fromLocation.CountryCode, fromLocation.City, fromLocation.Address, toLocation.Code, toLocation.PostalCode, toLocation.CountryCode, toLocation.City, toLocation.Address, size.Weight, size.Length, size.Width, size.Height)

	// Create a new HTTP request with the built request body.
	request, err := http.NewRequest("POST", c.ApiUrl, strings.NewReader(requestBody))
	if err != nil {
		return "", err
	}

	// Set the request headers.
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+c.Token)

	// Send the request to the CDEK API.
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check if the response status code is not OK.
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get a response from CDEK API: %d", resp.StatusCode)
	}

	// Read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		panic(err)
	}

	// выводим весь JSON
	jsonDataBytes, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		panic(err)
	}

	result := string(jsonDataBytes)
	return result, nil
}
