package models

import "time"

// Wager : wager info stored in DB
type Wager struct {
	ID                  uint      `gorm:"primary_key;column:id" json:"id"`
	TotalWagerValue     uint      `gorm:"column:total_wager_value" json:"total_wager_value"`
	Odds                uint      `gorm:"column:odds" json:"odds"`
	SellingPercentage   uint      `gorm:"column:selling_percentage" json:"selling_percentage"`
	SellingPrice        float64   `gorm:"column:selling_price" json:"selling_price"`
	CurrentSellingPrice float64   `gorm:"column:current_selling_price" json:"current_selling_price"`
	PercentageSold      uint      `gorm:"default:null;column:percentage_sold" json:"percentage_sold"`
	AmountSold          float64   `gorm:"default:null;column:amount_sold" json:"amount_sold"`
	PlacedAt            time.Time `gorm:"column:placed_at" json:"placed_at"`
}
