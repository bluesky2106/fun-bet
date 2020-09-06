package api

import (
	"strconv"

	"github.com/bluesky2106/fun-bet/config"
	"github.com/bluesky2106/fun-bet/servers"
	"github.com/gin-gonic/gin"
)

// Server : struct
type Server struct {
	config   *config.Config
	g        *gin.Engine
	wagerSrv *servers.WagerSrv
}

// NewServer : userSvc, walletSvc, assetSvc, config
func NewServer(config *config.Config,
	g *gin.Engine,
	wagerSrv *servers.WagerSrv,
) *Server {
	return &Server{
		config:   config,
		g:        g,
		wagerSrv: wagerSrv,
	}
}

func (s *Server) pagingFromContext(c *gin.Context) (int, int) {
	var (
		pageS  = c.DefaultQuery("page", "1")
		limitS = c.DefaultQuery("limit", "10")
		page   int
		limit  int
		err    error
	)

	page, err = strconv.Atoi(pageS)
	if err != nil {
		page = 1
	}

	limit, err = strconv.Atoi(limitS)
	if err != nil {
		limit = 10
	}

	return page, limit
}
