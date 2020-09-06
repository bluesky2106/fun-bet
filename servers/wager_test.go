package servers

import (
	"testing"
	"time"

	"github.com/bluesky2106/fun-bet/config"
	"github.com/bluesky2106/fun-bet/mocks"
	"github.com/bluesky2106/fun-bet/models"
	"github.com/bluesky2106/fun-bet/serializers"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type WagerTestSuite struct {
	suite.Suite

	wagerSrv     *WagerSrv
	mockWagerSvc *mocks.MockIWagerService
}

func (suite *WagerTestSuite) SetupTest() {
	conf := &config.Config{}
	mockCtrl := gomock.NewController(suite.T())
	defer mockCtrl.Finish()
	suite.mockWagerSvc = mocks.NewMockIWagerService(mockCtrl)
	suite.wagerSrv = NewWagerServer(conf, suite.mockWagerSvc)
}

func (suite *WagerTestSuite) TestWagerServerInitSuccessfully() {
	suite.NotNil(suite.wagerSrv)
}

func (suite *WagerTestSuite) TestPlaceWager() {
	// Test case 1: successful
	req := &serializers.PlaceWagerReq{
		TotalWagerValue:   100,
		Odds:              2,
		SellingPercentage: 20,
		SellingPrice:      25.12,
	}
	suite.mockWagerSvc.EXPECT().PlaceWager(req).Return(&models.Wager{
		ID:                  1,
		TotalWagerValue:     100,
		Odds:                2,
		SellingPercentage:   20,
		SellingPrice:        25.12,
		CurrentSellingPrice: 25.12,
		PercentageSold:      0,
		AmountSold:          0,
		PlacedAt:            time.Now(),
	}, nil).Times(1)
	wager, err := suite.wagerSrv.PlaceWager(req)
	suite.Nil(err)
	suite.NotNil(wager)

	// Test case 2: failed (wrong selling price format)
	req = &serializers.PlaceWagerReq{
		TotalWagerValue:   100,
		Odds:              2,
		SellingPercentage: 20,
		SellingPrice:      25.123,
	}
	wager, err = suite.wagerSrv.PlaceWager(req)
	suite.NotNil(err)
	suite.Nil(wager)
}

func TestWagerTestSuite(t *testing.T) {
	suite.Run(t, new(WagerTestSuite))
}
