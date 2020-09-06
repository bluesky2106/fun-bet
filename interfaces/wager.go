package interfaces

import (
	"github.com/bluesky2106/fun-bet/models"
	"github.com/bluesky2106/fun-bet/serializers"
)

// IWagerService : interface
type IWagerService interface {
	PlaceWager(req *serializers.PlaceWagerReq) (*models.Wager, error)
	BuyWager(req *serializers.BuyWagerReq) (*models.PurchaseOrder, error)
	ListWager(paging *serializers.PaginationReq) ([]*models.Wager, error)
}
