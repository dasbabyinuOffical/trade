package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

var (
	crontab *cron.Cron
)

func main() {
	crontab = cron.New()
	defer crontab.Stop()

	// 1分钟获取一次当前价格
	_, errCron := crontab.AddFunc("*/1 * * * *", func() {
		fmt.Println("start saveContractPrice cron.", time.Now())
		//每1分钟执行一次获取当前的价格
		saveContractPrice()

		fmt.Println("start  saveContractPricePercent.", time.Now())
		// 每分钟判断一次是否需要分析当前价格
		saveContractPricePercent()

		fmt.Println("start  sendMessageTo1hQueue.", time.Now())
		// 每分钟发送数据到消息队列
		sendMessageTo1hQueue()

	})
	if errCron != nil {
		panic(errCron)
	}

	// 启动定时任务
	crontab.Start()

	select {}
}
