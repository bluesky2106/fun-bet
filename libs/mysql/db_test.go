package mysql

import (
	"fmt"
	"testing"

	"github.com/bluesky2106/fun-bet/config"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

var (
	env     string
	connURL string
)

func init() {
	var (
		username = "prophet"
		password = "prophet"
		dbName   = "fun_bet"
		host     = "localhost"
		port     = "3309"
	)

	connURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC", username, password, host, port, dbName)
	env = config.EnvDevelopment
}

func TestInit(t *testing.T) {
	assert := assert.New(t)

	var (
		err error
		db  *gorm.DB
	)

	db, err = Init(connURL, env)
	assert.Nil(err)
	assert.NotNil(db)
}
