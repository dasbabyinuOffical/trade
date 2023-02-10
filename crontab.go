package main

import (
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Class int64

const (
	FiveMinutes    Class = 5 * 60
	FifteenMinutes Class = 15 * 60
	OneHour        Class = 3600
	FourHour       Class = 4 * 3600
	TwelveHour     Class = 12 * 3600
	OneDay         Class = 24 * 3600
)

func saveContractPrice() {
	contractItems, err := FetchContractPrice()
	if err != nil {
		fmt.Println("FetchContractPrice err:", err.Error())
		return
	}
	// 毫秒->分钟
	seconds := time.Now().Unix()
	now := seconds - seconds%60
	for _, item := range contractItems {
		price, errPrice := strconv.ParseFloat(item.Price, 64)
		if errPrice != nil {
			fmt.Println("price error:", item.Symbol, item.Time, errPrice.Error())
			continue
		}
		cp := &ContractPrice{
			Symbol: item.Symbol,
			Price:  price,
			Time:   now,
		}

		errDB := db.Create(cp).Error
		if errDB != nil {
			fmt.Println("insert contract  price error:", errDB.Error())
			continue
		}
	}
}

func getModelByClass(class Class, symbol string, time int64, priceBefore float64, priceNow float64) interface{} {
	percent := (priceNow - priceBefore) * 100 / priceBefore
	switch class {
	case FiveMinutes:
		return &ContractPrice5Min{
			Symbol:       symbol,
			PriceBefore:  priceBefore,
			PricePercent: percent,
			PriceNow:     priceNow,
			Time:         time,
		}
	case FifteenMinutes:
		return &ContractPrice15Min{
			Symbol:       symbol,
			PriceBefore:  priceBefore,
			PricePercent: percent,
			PriceNow:     priceNow,
			Time:         time,
		}
	case OneHour:
		return &ContractPrice1Hour{
			Symbol:       symbol,
			PriceBefore:  priceBefore,
			PricePercent: percent,
			PriceNow:     priceNow,
			Time:         time,
		}
	case FourHour:
		return &ContractPrice4Hour{
			Symbol:       symbol,
			PriceBefore:  priceBefore,
			PricePercent: percent,
			PriceNow:     priceNow,
			Time:         time,
		}
	case TwelveHour:
		return &ContractPrice12Hour{
			Symbol:       symbol,
			PriceBefore:  priceBefore,
			PricePercent: percent,
			PriceNow:     priceNow,
			Time:         time,
		}
	case OneDay:
		return &ContractPrice1Day{
			Symbol:       symbol,
			PriceBefore:  priceBefore,
			PricePercent: percent,
			PriceNow:     priceNow,
			Time:         time,
		}
	}
	return nil
}

func analysisPricePercent(seconds int64, class Class) {
	// 获取所有symbol分类
	var symbols []string
	err := db.Model(new(ContractPrice)).
		Select("distinct symbol as symbols").
		Pluck("symbols", &symbols).Error
	if err != nil {
		fmt.Println("analysisPricePercent get symbols failed.")
	}

	for _, symbol := range symbols {
		var (
			cp1 ContractPrice
			cp2 ContractPrice
		)
		// 获取之前的价格
		errPrice := db.Model(new(ContractPrice)).
			Where("symbol = ? and time = ?", symbol, seconds-int64(class)).
			First(&cp1).Error
		if errPrice == gorm.ErrRecordNotFound {
			continue
		}

		if errPrice != nil {
			fmt.Println("err get price history in loop", errPrice.Error())
			continue
		}

		// 获取之后的价格
		errNow := db.Model(new(ContractPrice)).
			Where("symbol = ? and time = ?", symbol, seconds).
			First(&cp2).Error
		if errNow == gorm.ErrRecordNotFound {
			continue
		}

		if errNow != nil {
			fmt.Println("err get price now in loop", errNow.Error())
			continue
		}

		// 获取价格变化范围
		m := getModelByClass(class, symbol, seconds, cp1.Price, cp2.Price)
		errPercent := db.Create(m).Error
		if errPercent != nil {
			fmt.Println("create percent err:", errPercent)
			continue
		}
	}
}

func saveContractPricePercent() {
	seconds := time.Now().Unix()
	seconds = seconds - seconds%60
	// 每5分钟执行一次
	if seconds%(5*60) == 0 {
		analysisPricePercent(seconds, FiveMinutes)
	}

	// 每15分钟执行一次
	if seconds%(15*60) == 0 {
		analysisPricePercent(seconds, FifteenMinutes)
	}

	// 每1小时执行一次
	if seconds%3600 == 0 {
		analysisPricePercent(seconds, OneHour)
	}

	// 每4小时执行一次
	if seconds%(4*3600) == 0 {
		analysisPricePercent(seconds, FourHour)
	}

	// 每12小时执行一次
	if seconds%(12*3600) == 0 {
		analysisPricePercent(seconds, TwelveHour)
	}

	// 每1天执行一次
	if seconds%(24*3600) == 0 {
		analysisPricePercent(seconds, OneDay)
	}
}
