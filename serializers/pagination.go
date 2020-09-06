package serializers

// PaginationReq : pagination request
type PaginationReq struct {
	Page    uint   `json:"page"`
	Limit   uint   `json:"limit"`
	SortBy  string `json:"sort_by"`
	OrderBy string `json:"order_by"`
}
