package api

import (
	"github.com/bluesky2106/fun-bet/serializers"
	"github.com/gin-gonic/gin"
)

func respondError(c *gin.Context, status int, err error) {
	errResp := serializers.ErrorResp{
		Error: err.Error(),
	}
	c.JSON(status, errResp)
}
