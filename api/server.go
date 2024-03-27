package api

import (
	db "github.com/TTKirito/backend-go/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store db.Store
	route *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	route := gin.Default()

	route.POST("/accounts", server.createAccount)
	route.GET("/accounts/:id", server.getAccount)
	route.GET("/accounts", server.listAccount)

	server.route = route

	return server
}

func (server *Server) Start(address string) error {
	return server.route.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
