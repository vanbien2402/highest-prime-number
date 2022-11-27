package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

//Start server
func (s *Server) Start() error {
	return s.router.Run(":8080")
}

//NewServer initialize new server
func NewServer() (*Server, error) {
	router := gin.New()
	_ = router.SetTrustedProxies(nil)
	initRouter(router)
	return &Server{
		router: router,
	}, nil
}
