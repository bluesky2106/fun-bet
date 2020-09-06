package servers

import (
	"github.com/bluesky2106/fun-bet/config"
	"github.com/bluesky2106/fun-bet/interfaces"
	"github.com/bluesky2106/fun-bet/models"
	"github.com/bluesky2106/fun-bet/serializers"
)

// WagerSrv : wager server
type WagerSrv struct {
	conf     *config.Config
	wagerSvc interfaces.IWagerService
}

// NewWagerServer : config, wager message
func NewWagerServer(conf *config.Config, wagerSvc interfaces.IWagerService) *WagerSrv {
	return &WagerSrv{
		conf:     conf,
		wagerSvc: wagerSvc,
	}
}

// PlaceWager : create a wager
func (svr *WagerSrv) PlaceWager(req *serializers.PlaceWagerReq) (*models.Wager, error) {
	return svr.wagerSvc.PlaceWager(req)
}

// BuyWager : place a wager
func (svr *WagerSrv) BuyWager(req *serializers.BuyWagerReq) (*models.PurchaseOrder, error) {
	return svr.wagerSvc.BuyWager(req)
}

// ListWager : list all wagers
func (svr *WagerSrv) ListWager(paging *serializers.PaginationReq) ([]*models.Wager, error) {
	return svr.wagerSvc.ListWager(paging)
}
