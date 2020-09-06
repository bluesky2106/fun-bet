package daos

import (
	errs "github.com/bluesky2106/fun-bet/errors"
	"github.com/bluesky2106/fun-bet/models"
	"github.com/bluesky2106/fun-bet/serializers"
	"github.com/jinzhu/gorm"
)

const (
	wagerTable = "wagers"
)

// Wager : struct
type Wager struct {
}

// NewWager : new wager dao
func NewWager() *Wager {
	return &Wager{}
}

// Create : tx, mod
func (dao *Wager) Create(tx *gorm.DB, mod *models.Wager) error {
	err := tx.Create(mod).Error
	if err != nil {
		return errs.Wrap(errs.ErrMySQLCreate, err.Error(), "tx.Create")
	}
	return nil
}

// Update : tx, mod
func (dao *Wager) Update(tx *gorm.DB, mod *models.Wager) error {
	err := tx.Save(mod).Error
	if err != nil {
		return errs.Wrap(errs.ErrMySQLUpdate, err.Error(), "tx.Save")
	}
	return nil
}

// FindByID : id
func (dao *Wager) FindByID(id uint64) (*models.Wager, error) {
	var mod models.Wager
	if err := db.First(&mod, id).Error; err != nil {
		return nil, errs.Wrap(errs.ErrMySQLRead, err.Error(), "db.First")
	}
	return &mod, nil
}

// FindAllByQuery :
func (dao *Wager) FindAllByQuery(filters map[string]interface{}, paging *serializers.PaginationReq) ([]*models.Wager, error) {
	if paging.Page == 0 {
		paging.Page = 1
	}

	var (
		models []*models.Wager
		offset = paging.Page*paging.Limit - paging.Limit
		order  = buildOrder(paging)
	)

	query := GetDB().Table(wagerTable)
	query = query.Limit(paging.Limit).Offset(offset).Order(order)
	query = where(query, filters)

	if err := query.Find(&models).Error; err != nil {
		return nil, errs.Wrap(errs.ErrMySQLRead, err.Error(), "query.Find")
	}
	return models, nil
}

// CountByQuery :
func (dao *Wager) CountByQuery(filters map[string]interface{}) (uint, error) {
	var count uint

	query := db.Table(wagerTable)
	query = where(query, filters)

	if err := query.Count(&count).Error; err != nil {
		return 0, errs.Wrap(errs.ErrMySQLRead, err.Error(), "query.Count")
	}
	return count, nil
}
