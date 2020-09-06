package servers

import (
	"math"

	errs "github.com/bluesky2106/fun-bet/errors"
	"github.com/bluesky2106/fun-bet/serializers"
)

const eps = 1E-9

func (svr *WagerSrv) validatePlaceWagerReq(req *serializers.PlaceWagerReq) error {
	if req.SellingPercentage < 1 || req.SellingPercentage > 100 {
		return errs.ErrInvalidSellingPercentage
	}

	if math.Abs(req.SellingPrice*100-math.Round(req.SellingPrice*100)) > eps {
		return errs.ErrInvalidSellingPriceFormat
	}
	if req.SellingPrice <= 0 {
		return errs.ErrInvalidSellingPrice
	}
	if req.SellingPrice <= float64(req.TotalWagerValue)*(float64(req.SellingPercentage)/100) {
		return errs.ErrInvalidSellingPrice
	}

	return nil
}

func (svr *WagerSrv) validateBuyWagerReq(wagerID uint, req *serializers.BuyWagerReq) error {
	if req.BuyingPrice <= 0 {
		return errs.ErrInvalidBuyingPrice
	}
	wager, err := svr.wagerSvc.ReadWager(wagerID)
	if err != nil {
		return err
	}
	if req.BuyingPrice > wager.CurrentSellingPrice {
		return errs.ErrInvalidBuyingPrice
	}

	return nil
}
