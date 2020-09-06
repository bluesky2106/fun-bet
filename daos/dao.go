package daos

import (
	"fmt"
	"strings"

	"github.com/bluesky2106/fun-bet/config"
	errs "github.com/bluesky2106/fun-bet/errors"
	"github.com/bluesky2106/fun-bet/libs/mysql"
	"github.com/bluesky2106/fun-bet/models"
	"github.com/bluesky2106/fun-bet/serializers"
	"github.com/jinzhu/gorm"
)

var (
	tables = []interface{}{(*models.Wager)(nil), (*models.PurchaseOrder)(nil)}
	db     *gorm.DB
)

// GetDB : getter
func GetDB() *gorm.DB {
	return db
}

// AutoMigrate : migrate DB
func AutoMigrate() error {
	if db == nil {
		return errs.ErrMySQLDBEmpty
	}

	if err := db.AutoMigrate(tables...).Error; err != nil {
		return errs.Wrap(errs.ErrMySQLDBAutoMigrate, err.Error())
	}

	return nil
}

// WithTransaction :
func WithTransaction(callback func(*gorm.DB) error) error {
	tx := db.Begin()

	if err := callback(tx); err != nil {
		tx.Rollback()
		return errs.Wrap(err, "callback")
	}

	if err := tx.Commit().Error; err != nil {
		return errs.Wrap(err, "tx.Commit()")
	}

	return nil
}

func where(db *gorm.DB, filters map[string]interface{}) *gorm.DB {
	query := db
	for k, v := range filters {
		if v != nil {
			query = query.Where(k, v)
		} else {
			query = query.Where(k)
		}
	}
	return query
}

// Init : initialize DB
func Init(conf *config.Config) error {
	var (
		err     error
		connURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC", conf.MySQL.Username, conf.MySQL.Password, conf.MySQL.Host, conf.MySQL.Port, conf.MySQL.DBName)
	)

	db, err = mysql.Init(connURL, conf.Env)
	if err != nil {
		return errs.Wrap(err, "mysql.Init")
	}
	return nil
}

func buildOrder(paging *serializers.PaginationReq) string {
	orderBy := strings.ToLower(paging.OrderBy)
	sortBy := strings.ToLower(paging.SortBy)

	if orderBy == "" {
		orderBy = "asc"
	}

	if sortBy == "" {
		sortBy = "id"
	}

	return fmt.Sprintf("%s %s", sortBy, orderBy)
}
