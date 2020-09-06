package api

// Routes : ...
func (s *Server) Routes() {
	// place wager
	s.g.POST("/wagers", s.PlaceWager)

	// buy wager
	s.g.POST("buy/:wager_id", s.BuyWager)

	// list wagers
	s.g.GET("/wagers", s.ListWager)
}
