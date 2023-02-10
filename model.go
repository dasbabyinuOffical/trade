package main

import (
	"gorm.io/gorm"
)

type ContractPrice struct {
	gorm.Model
	Symbol string  `gorm:"symbol" json:"symbol"`
	Price  float64 `gorm:"price" json:"price"`
	Time   int64   `gorm:"time"  json:"time"`
}

func (p *ContractPrice) TableName() string {
	return "contract_price"
}

type ContractPrice1Day struct {
	gorm.Model
	Symbol       string  `gorm:"symbol" json:"symbol"`
	PriceBefore  float64 `gorm:"price_before" json:"priceBefore"`
	PricePercent float64 `gorm:"price_percent" json:"pricePercent"`
	PriceNow     float64 `gorm:"price_now" json:"priceNow"`
	Time         int64   `gorm:"time"  json:"time"`
}

func (p *ContractPrice1Day) TableName() string {
	return "contract_price_1d"
}

type ContractPrice1Hour struct {
	gorm.Model
	Symbol       string  `gorm:"symbol" json:"symbol"`
	PriceBefore  float64 `gorm:"price_before" json:"priceBefore"`
	PricePercent float64 `gorm:"price_percent" json:"pricePercent"`
	PriceNow     float64 `gorm:"price_now" json:"priceNow"`
	Time         int64   `gorm:"time"  json:"time"`
}

func (p *ContractPrice1Hour) TableName() string {
	return "contract_price_1h"
}

type ContractPrice4Hour struct {
	gorm.Model
	Symbol       string  `gorm:"symbol" json:"symbol"`
	PriceBefore  float64 `gorm:"price_before" json:"priceBefore"`
	PricePercent float64 `gorm:"price_percent" json:"pricePercent"`
	PriceNow     float64 `gorm:"price_now" json:"priceNow"`
	Time         int64   `gorm:"time"  json:"time"`
}

func (p *ContractPrice4Hour) TableName() string {
	return "contract_price_4h"
}

type ContractPrice12Hour struct {
	gorm.Model
	Symbol       string  `gorm:"symbol" json:"symbol"`
	PriceBefore  float64 `gorm:"price_before" json:"priceBefore"`
	PricePercent float64 `gorm:"price_percent" json:"pricePercent"`
	PriceNow     float64 `gorm:"price_now" json:"priceNow"`
	Time         int64   `gorm:"time"  json:"time"`
}

func (p *ContractPrice12Hour) TableName() string {
	return "contract_price_12h"
}

type ContractPrice5Min struct {
	gorm.Model
	Symbol       string  `gorm:"symbol" json:"symbol"`
	PriceBefore  float64 `gorm:"price_before" json:"priceBefore"`
	PricePercent float64 `gorm:"price_percent" json:"pricePercent"`
	PriceNow     float64 `gorm:"price_now" json:"priceNow"`
	Time         int64   `gorm:"time"  json:"time"`
}

func (p *ContractPrice5Min) TableName() string {
	return "contract_price_5m"
}

type ContractPrice15Min struct {
	gorm.Model
	Symbol       string  `gorm:"symbol" json:"symbol"`
	PriceBefore  float64 `gorm:"price_before" json:"priceBefore"`
	PricePercent float64 `gorm:"price_percent" json:"pricePercent"`
	PriceNow     float64 `gorm:"price_now" json:"priceNow"`
	Time         int64   `gorm:"time"  json:"time"`
}

func (p *ContractPrice15Min) TableName() string {
	return "contract_price_15m"
}
