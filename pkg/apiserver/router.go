package apiserver

func RouterGroupV1(s *Server) {
	s.Router.Group("v1").POST("/records", s.CreateRecord)
	s.Router.Group("v1").DELETE("/records/:name", s.DeleteRecord)
	s.Router.Group("v1").GET("/records/:name", s.GetRecord)
	s.Router.Group("v1").Group("/records", s.ListRecord)
}
