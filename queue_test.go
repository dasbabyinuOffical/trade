package main

import (
	"fmt"
	"testing"
)

func TestSendMessageTo1hQueue(t *testing.T) {
	// init db
	fmt.Println("init db.")
	initDB()
	fmt.Println("init db done.")

	// init redis
	fmt.Println("init redis.")
	initRedis()
	fmt.Println("init redis done.")

	sendMessageTo1hQueue()
}