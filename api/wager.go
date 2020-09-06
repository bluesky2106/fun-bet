package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/bluesky2106/fun-bet/serializers"
	"github.com/gin-gonic/gin"
)

// PlaceWager : create a wager
func (s *Server) PlaceWager(c *gin.Context) {
	var req serializers.PlaceWagerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		respondError(c, http.StatusBadRequest, err)
		return
	}

	wager, err := s.wagerSrv.PlaceWager(&req)
	if err != nil {
		respondError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, wager)
}

// BuyWager : buy a wager
func (s *Server) BuyWager(c *gin.Context) {
	var req serializers.BuyWagerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		respondError(c, http.StatusBadRequest, err)
		return
	}

	idStr := c.Param("wager_id")
	wagerID, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(c, http.StatusBadRequest, err)
		return
	}

	po, err := s.wagerSrv.BuyWager(uint(wagerID), &req)
	if err != nil {
		respondError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, po)
}

// ListWager : list all wagers
func (s *Server) ListWager(c *gin.Context) {
	page, limit := s.pagingFromContext(c)
	var (
		sortBy  = strings.ToLower(c.DefaultQuery("sort_by", "id"))
		orderBy = strings.ToLower(c.DefaultQuery("order_by", "asc"))
	)

	if orderBy != "asc" && orderBy != "desc" {
		orderBy = "asc"
	}

	paging := &serializers.PaginationReq{
		Page:    uint(page),
		Limit:   uint(limit),
		SortBy:  sortBy,
		OrderBy: orderBy,
	}

	wagers, err := s.wagerSrv.ListWager(paging)
	if err != nil {
		respondError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, wagers)
}
