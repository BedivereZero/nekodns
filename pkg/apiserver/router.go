package apiserver

func RouterGroupV1(s *Server) {
	s.Router.Group("v1").POST("/records", s.CreateRecord)
	s.Router.Group("v1").GET("/records/:name", s.GetRecord)
}
