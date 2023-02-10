package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	URL = "https://api.telegram.org/bot5854505388:AAEnQbGKzjh_19gMr-B7aW-beDT3ZaYnplY/sendMessage?chat_id=-1001666206862&text=%s"
)

func sendMessage(message string) (res string, err error) {
	url := fmt.Sprintf(URL, message)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	res = string(data)
	return
}
