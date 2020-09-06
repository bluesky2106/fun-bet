package api

// Routes : ...
func (s *Server) Routes() {
	s.g.GET("/", s.DefaultWelcome)
	api := s.g.Group("/api")
	{
		api.GET("/", s.Welcome)

		// place wager
		api.POST("/wagers", s.PlaceWager)

		// buy wager
		api.POST("buy/:wager_id", s.BuyWager)

		// list wagers
		api.GET("/wagers", s.ListWager)
	}
}
