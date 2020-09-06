package main

import (
	"fmt"
	"time"

	"github.com/bluesky2106/fun-bet/api"
	"github.com/bluesky2106/fun-bet/config"
	"github.com/bluesky2106/fun-bet/daos"
	"github.com/bluesky2106/fun-bet/log"
	"github.com/bluesky2106/fun-bet/servers"
	"github.com/bluesky2106/fun-bet/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	conf := config.ParseConfig("config.json", "./config")
	conf.Print()

	// Init logger
	log.InitLogger(conf.Env)

	// Init dao
	initDAO(conf)
	wagerDAO := daos.NewWager()
	purchaseOrderDAO := daos.NewPurchaseOrder()

	// Init services
	wagerSvc := services.NewWagerService(wagerDAO, purchaseOrderDAO)

	// Init servers
	wagerSrv := servers.NewWagerServer(conf, wagerSvc)

	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://*", "https://*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		MaxAge:           12 * time.Hour,
	}))

	svr := api.NewServer(conf, router, wagerSrv)
	svr.Routes()
	if err := router.Run(fmt.Sprintf("%s:%s", conf.Host, conf.Port)); err != nil {
		log.GetLogger().Error("router.Run", zap.Error(err))
	}
}

func initDAO(conf *config.Config) {
	if err := daos.Init(conf); err != nil {
		log.GetLogger().Error("failed to init mysql:", zap.Error(err))
	}
	if err := daos.AutoMigrate(); err != nil {
		log.GetLogger().Error("failed to migrate database:", zap.Error(err))
	}
}
