package main

import (
	"ParcelCalculator/cdeklib"
	"fmt"
)

func main() {
	// Account information
	account := "EMscd6r9JnFiQ3bLoyjJY6eM78JrJceI"
	securePassword := "PjLZkKBHEiLK3YsjtNrt3TGNG0ahs3kG"
	// Set up API URLs
	apiURL := "https://api.edu.cdek.ru/v2/oauth/token?parameters"
	apiUrlTariffList := "https://api.edu.cdek.ru/v2/calculator/tarifflist"
	// Set up sending and delivery locations
	sendingAddress := "Россия, г. Москва, Cлавянский бульвар д.1"
	deliveryAddress := "Россия, Воронежская обл., г. Воронеж, ул. Ленина д.43"

	// Sending address
	fromLocation := cdeklib.Location{
		Address: sendingAddress,
	}

	// Delivery address
	toLocation := cdeklib.Location{
		Address: deliveryAddress,
	}

	// Parcel size
	size := cdeklib.Size{
		Weight: 1000,
		Length: 20,
		Width:  20,
		Height: 20,
	}

	var accessToken, err = cdeklib.GetAccessToken(apiURL, account, securePassword)
	if err != nil {
		fmt.Println("Error getting access token:", err)
		return
	}

	var client = cdeklib.NewClient(accessToken, true, apiUrlTariffList)

	tariffs, err := client.Calculate(fromLocation, toLocation, size)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(tariffs)
}
