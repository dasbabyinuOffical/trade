package main

import (
	"fmt"
	"strconv"
)

func saveContractPrice() {
	contractItems, err := FetchContractPrice()
	if err != nil {
		fmt.Println("FetchContractPrice err:", err.Error())
		return
	}
	for _, item := range contractItems {
		price, errPrice := strconv.ParseFloat(item.Price, 64)
		if errPrice != nil {
			fmt.Println("price error:", item.Symbol, item.Time, errPrice.Error())
			continue
		}
		cp := &ContractPrice{
			Symbol: item.Symbol,
			Price:  price,
			Time:   item.Time / 1000,
		}

		errDB := db.Create(cp).Error
		if errDB != nil {
			fmt.Println("insert contract  price error:", errDB.Error())
			continue
		}
	}
}
