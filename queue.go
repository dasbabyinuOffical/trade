package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
)

const (
	QueueContract1h string = "queue:contract_1h"
)

// 发送1小时价格波动超过3%的币到1小时队列
func sendMessageTo1hQueue() {
	var cp1h []*ContractPrice1Hour
	err = db.Where("is_sync = ?", 0).Find(&cp1h).Error
	if err == gorm.ErrRecordNotFound || len(cp1h) == 0 {
		fmt.Println("当前1h无新数据", time.Now())
		return
	}
	if err != nil {
		fmt.Println("sendMessageTo1hQueue get database err:", err.Error())
	}

	// 发送到redis队列
	for _, v := range cp1h {
		// 策略在-3% -- 3%之间忽略
		var template string
		if v.PricePercent > -3 && v.PricePercent < 3 {
			continue
		}
		if v.PricePercent > 3 {
			template = templateUp
		}
		if v.PricePercent < -3 {
			template = templateDown
		}
		ts := time.Unix(v.Time, 0).Format("2006-01-02 15:04")
		template = fmt.Sprintf(template, v.Symbol, v.PricePercent, v.PriceBefore, v.PriceNow, ts)
		_, errPush := rdb.LPush(QueueContract1h, template).Result()
		if errPush != nil {
			fmt.Println("sendMessageTo1hQueue failed:", errPush.Error())
		}
	}

	// 设置所有数据为已读
	err = db.Model(&ContractPrice1Hour{}).
		Where("id <= ?", cp1h[len(cp1h)-1].ID).
		UpdateColumn("is_sync", 1).Error
	if err != nil {
		fmt.Println("sendMessageTo1hQueue update database err:", err.Error())
	}
	return
}

func sendQueueMessageToBot() {
	// 发送1小时机器人队列消息到bot中
	for {
		// 设置一个5秒的超时时间
		value, err := rdb.BRPop(5*time.Second, QueueContract1h).Result()
		if err == redis.Nil {
			// 查询不到数据
			time.Sleep(1 * time.Second)
			continue
		}
		if err != nil {
			// 查询出错
			fmt.Println("sendMessageTo1hQueue failed:", err.Error())
			time.Sleep(1 * time.Second)
			continue
		}
		log.Println("消费到数据：", value, "当前时间是：", time.Now())
		time.Sleep(time.Second)

		// 发送数据到bot
		message := strings.Join(value[1:], ",")
		_, errSend := sendMessageToTelegram(message)
		if errSend != nil {
			fmt.Println("sendMessageTo1hQueue failed:", errSend.Error())
		}
		time.Sleep(time.Second)
	}
}
