package main

import (
	"github.com/robfig/cron/v3"
)

var (
	crontab *cron.Cron
)

func main() {
	crontab = cron.New()
	defer crontab.Stop()

	// 1分钟获取一次当前价格
	_, errCron := crontab.AddFunc("*/1 * * * *", saveContractPrice)
	if errCron != nil {
		panic(errCron)
	}

	// 启动定时任务
	crontab.Start()

	select {}
}
