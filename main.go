package main

import (
	"ParcelCalculator/cdeklib/core"
	"ParcelCalculator/cdeklib/core/types"
	"fmt"
)

func main() {
	account := "EMscd6r9JnFiQ3bLoyjJY6eM78JrJceI"
	securePassword := "PjLZkKBHEiLK3YsjtNrt3TGNG0ahs3kG"
	apiURL := "https://api.edu.cdek.ru/"

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

	var client = core.NewClient(true, apiURL, account, securePassword)

	tariffs, err := client.Calculate(fromLocation, toLocation, size)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(tariffs)

	//CreateOrder
	recipient := types.Recipient{
		Name: "Name",
		Phones: []types.Phone{
			{
				Number: "+79991112233",
			},
		},
	}
	packages := []types.Package{
		{
			Length: 15,
			Width:  25,
			Height: 30,
			Weight: 1000,
			Number: "TestNumber",
			Items: []types.Item{
				{
					Name:    "TestItem",
					WareKey: "TestWareKey",
					Payment: types.Money{Value: 0},
					Value:   0,
					Cost:    1000,
					Weight:  1000,
					Amount:  1,
				},
			},
		},
	}

	orderID, err := client.CreateOrder(sendingAddress, deliveryAddress, recipient, packages, 233)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(orderID)

	//CheckOrder
	status, err := client.GetStatus(orderID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(status)

}
