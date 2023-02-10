package main

import (
	"fmt"
	"gorm.io/gorm"
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
		return
	}
	if err != nil {
		fmt.Println("sendMessageTo1hQueue get database err:", err.Error())
	}

	// 发送到redis队列
	for _, v := range cp1h {
		// 策略在-3% -- 3%之间忽略
		if v.PricePercent < 3 && v.PricePercent > -3 {
			continue
		}
		ts := time.Unix(v.Time, 0).Format("2006-01-02 15:04")
		_, errPush := rdb.LPush(QueueContract1h, fmt.Sprintf("%s:%+f:%s", v.Symbol, v.PricePercent, ts)).Result()
		if errPush != nil {
			fmt.Println("sendMessageTo1hQueue failed:", errPush.Error())
		}
	}

	// 设置所有数据为已读
	err = db.Model(&ContractPrice1Hour{}).
		Where("id <= ?", cp1h[len(cp1h)-1].ID).
		Update("is_sync = ?", 1).Error
	if err != nil {
		fmt.Println("sendMessageTo1hQueue update database err:", err.Error())
	}
	return
}
