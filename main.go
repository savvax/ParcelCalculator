package main

import (
	"ParcelCalculator/cdeklib/core"
	"ParcelCalculator/cdeklib/core/types"
	"fmt"
)

func main() {
	// Account information
	account := "EMscd6r9JnFiQ3bLoyjJY6eM78JrJceI"
	securePassword := "PjLZkKBHEiLK3YsjtNrt3TGNG0ahs3kG"
	// Set up API URLs
	apiURL := "https://api.edu.cdek.ru/v2/oauth/token?parameters"
	apiUrlTariffList := "https://api.edu.cdek.ru/v2/calculator/tarifflist"
	//apiURLloc := "https://api.edu.cdek.ru/v2/location/cities"
	// Set up sending and delivery locations
	sendingAddress := "Россия, г. Москва, Cлавянский бульвар д.1"
	deliveryAddress := "Россия, Воронежская обл., г. Воронеж, ул. Ленина д.43"

	// Sending address
	fromLocation := types.LocationCalc{
		Address: sendingAddress,
	}

	// Delivery address
	toLocation := types.LocationCalc{
		Address: deliveryAddress,
	}

	// Parcel size
	size := types.Size{
		Weight: 1000,
		Length: 20,
		Width:  20,
		Height: 20,
	}

	var client = core.NewClient(true, apiURL, apiUrlTariffList, account, securePassword)

	tariffs, err := client.Calculate(fromLocation, toLocation, size)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(tariffs)
}
