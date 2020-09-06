package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// PurchaseOrder : purchase order info stored in DB
type PurchaseOrder struct {
	gorm.Model
	ID          uint      `gorm:"primary_key"`
	WagerID     uint      `gorm:"index;column:wager_id" json:"wager_id"`
	BuyingPrice uint      `gorm:"column:buying_price" json:"buying_price"`
	BoughtAt    time.Time `gorm:"column:bought_at" json:"bought_at"`
}
