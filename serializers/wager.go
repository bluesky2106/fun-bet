package serializers

import (
	"encoding/json"

	errs "github.com/bluesky2106/fun-bet/errors"
	"github.com/bluesky2106/fun-bet/models"
)

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

// ConvertToWagerModel : convert request to wager model
func (req *PlaceWagerReq) ConvertToWagerModel() (*models.Wager, error) {
	bs, err := json.Marshal(req)
	if err != nil {
		return nil, errs.Wrap(errs.ErrSystem, err.Error())
	}

	var mod *models.Wager
	if err := json.Unmarshal(bs, &mod); err != nil {
		return nil, errs.Wrap(errs.ErrSystem, err.Error())
	}

	return mod, nil
}
