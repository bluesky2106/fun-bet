package daos

import (
	errs "github.com/bluesky2106/fun-bet/errors"
	"github.com/bluesky2106/fun-bet/models"
	"github.com/jinzhu/gorm"
)

const (
	purchaseOrderTable = "purchase_orders"
)

// PurchaseOrder : struct
type PurchaseOrder struct {
}

// NewPurchaseOrder : new purchase order dao
func NewPurchaseOrder() *PurchaseOrder {
	return &PurchaseOrder{}
}

// Create : tx, mod
func (dao *PurchaseOrder) Create(tx *gorm.DB, mod *models.PurchaseOrder) error {
	err := tx.Create(mod).Error
	if err != nil {
		return errs.Wrap(errs.ErrMySQLCreate, err.Error(), "tx.Create")
	}
	return nil
}
