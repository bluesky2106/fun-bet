package services

import (
	"github.com/bluesky2106/fun-bet/daos"
	errs "github.com/bluesky2106/fun-bet/errors"
	"github.com/bluesky2106/fun-bet/models"
	"github.com/bluesky2106/fun-bet/serializers"
	"github.com/jinzhu/gorm"
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
	var (
		mod *models.Wager
		err error
	)

	err = daos.WithTransaction(func(tx *gorm.DB) error {
		mod, err = req.ConvertToWagerModel()
		if err != nil {
			return errs.Wrap(err, "req.ConvertToWagerModel")
		}

		err = svc.wd.Create(tx, mod)
		if err != nil {
			return errs.Wrap(err, "svc.wd.Create")
		}

		return nil
	})

	if err != nil {
		return nil, errs.Wrap(err, "daos.WithTransaction")
	}

	return mod, nil
}

// BuyWager : place a wager
func (svc *WagerSvc) BuyWager(wagerID uint, req *serializers.BuyWagerReq) (*models.PurchaseOrder, error) {
	return nil, nil
}

// ListWager : list all wagers
func (svc *WagerSvc) ListWager(paging *serializers.PaginationReq) ([]*models.Wager, error) {
	return nil, nil
}
