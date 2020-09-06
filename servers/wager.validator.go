package servers

import (
	"math"

	errs "github.com/bluesky2106/fun-bet/errors"
	"github.com/bluesky2106/fun-bet/serializers"
)

func validatePlaceWagerReq(req *serializers.PlaceWagerReq) error {
	if req.SellingPercentage < 1 || req.SellingPercentage > 100 {
		return errs.ErrInvalidSellingPercentage
	}

	if req.SellingPrice != math.Floor(req.SellingPrice*100)/100 {
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
