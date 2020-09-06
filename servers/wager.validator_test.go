package servers

import (
	"testing"

	errs "github.com/bluesky2106/fun-bet/errors"
	"github.com/bluesky2106/fun-bet/serializers"
	"github.com/stretchr/testify/assert"
)

func TestValidatePlaceWagerReq(t *testing.T) {
	assert := assert.New(t)

	srv := NewWagerServer(nil, nil)

	var (
		req *serializers.PlaceWagerReq
	)

	// wrong selling percentage
	req = &serializers.PlaceWagerReq{
		TotalWagerValue:   100,
		Odds:              2,
		SellingPercentage: 101,
		SellingPrice:      25.12,
	}
	err := srv.validatePlaceWagerReq(req)
	assert.NotNil(err)
	assert.Equal(errs.ErrInvalidSellingPercentage, err)

	// wrong selling price format
	req = &serializers.PlaceWagerReq{
		TotalWagerValue:   100,
		Odds:              2,
		SellingPercentage: 10,
		SellingPrice:      25.567,
	}
	err = srv.validatePlaceWagerReq(req)
	assert.NotNil(err)
	assert.Equal(errs.ErrInvalidSellingPriceFormat, err)

	// invalid selling price
	req = &serializers.PlaceWagerReq{
		TotalWagerValue:   100,
		Odds:              2,
		SellingPercentage: 20,
		SellingPrice:      10.12,
	}
	err = srv.validatePlaceWagerReq(req)
	assert.NotNil(err)
	assert.Equal(errs.ErrInvalidSellingPrice, err)
}
