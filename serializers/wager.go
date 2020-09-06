package serializers

// PlaceWagerReq : create wager request
type PlaceWagerReq struct {
	TotalWagerValue   uint    `json:"total_wager_value"`
	Odds              uint    `json:"odds"`
	SellingPercentage uint    `json:"selling_percentage"`
	SellingPrice      float64 `json:"selling_price"`
}

// BuyWagerReq : buy wager request
type BuyWagerReq struct {
	BuyingPrice uint `json:"buying_price"`
}
