package services

import (
	"time"

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
		mod.PlacedAt = time.Now().UTC()
		mod.CurrentSellingPrice = mod.SellingPrice

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
	var (
		mod *models.PurchaseOrder
		err error
	)

	err = daos.WithTransaction(func(tx *gorm.DB) error {
		mod = &models.PurchaseOrder{
			WagerID:     wagerID,
			BuyingPrice: req.BuyingPrice,
			BoughtAt:    time.Now().UTC(),
		}

		err = svc.pod.Create(tx, mod)
		if err != nil {
			return errs.Wrap(err, "svc.pod.Create")
		}

		return nil
	})

	if err != nil {
		return nil, errs.Wrap(err, "daos.WithTransaction")
	}

	return mod, nil
}

// ListWager : list all wagers
func (svc *WagerSvc) ListWager(paging *serializers.PaginationReq) ([]*models.Wager, error) {
	mods, err := svc.wd.FindAllByQuery(nil, paging)
	if err != nil {
		return nil, errs.Wrap(err, "svc.wd.FindAllByQuery")
	}

	return mods, nil
}
