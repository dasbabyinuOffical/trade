package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ContractItem struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
	Time   int64  `json:"time"`
}

func FetchContractPrice() (contractItems []*ContractItem, err error) {
	url := "https://fapi.binance.com/fapi/v1/ticker/price"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return
	}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &contractItems)
	if err != nil {
		return
	}
	return
}
