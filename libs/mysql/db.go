package mysql

import (

	// mysql driver
	"github.com/bluesky2106/fun-bet/config"
	errs "github.com/bluesky2106/fun-bet/errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Init : connect mysql server
func Init(connURL string, env string) (*gorm.DB, error) {
	databaseConn, err := gorm.Open("mysql", connURL)
	if err != nil {
		return nil, errs.Wrap(errs.ErrMySQLConnErrtion, err.Error())
	}
	if env != config.EnvProduction {
		databaseConn.LogMode(true)
	}

	databaseConn = databaseConn.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci auto_increment=1")
	// skip save associations of gorm -> manual save by code
	databaseConn = databaseConn.Set("gorm:save_associations", false)
	databaseConn = databaseConn.Set("gorm:association_save_reference", true)
	databaseConn.DB().SetMaxOpenConns(20)
	databaseConn.DB().SetMaxIdleConns(10)

	return databaseConn, nil
}
