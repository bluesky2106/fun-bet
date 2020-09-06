package models

import (
	"time"
)

// PurchaseOrder : purchase order info stored in DB
type PurchaseOrder struct {
	ID          uint      `gorm:"primary_key;column:id" json:"id"`
	WagerID     uint      `gorm:"index;column:wager_id" json:"wager_id"`
	BuyingPrice float64   `gorm:"column:buying_price" json:"buying_price"`
	BoughtAt    time.Time `gorm:"column:bought_at" json:"bought_at"`
}
