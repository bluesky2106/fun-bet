package services

import (
	"github.com/bluesky2106/fun-bet/daos"
	"github.com/bluesky2106/fun-bet/models"
	"github.com/bluesky2106/fun-bet/serializers"
)

// WagerSvc : struct
type WagerSvc struct {
	wd  *daos.Wager
	pod *daos.PurchaseOrder
}

// NewWagerService : config
func NewWagerService(wd *daos.Wager, pod *daos.PurchaseOrder) *WagerSvc {
	return &WagerSvc{
		wd:  wd,
		pod: pod,
	}
}

// PlaceWager : create a wager
func (svc *WagerSvc) PlaceWager(req *serializers.PlaceWagerReq) (*models.Wager, error) {
	return nil, nil
}

// BuyWager : place a wager
func (svc *WagerSvc) BuyWager(req *serializers.BuyWagerReq) (*models.PurchaseOrder, error) {
	return nil, nil
}

// ListWager : list all wagers
func (svc *WagerSvc) ListWager(paging *serializers.PaginationReq) ([]*models.Wager, error) {
	return nil, nil
}
