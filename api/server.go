package api

import (
	db "github.com/TTKirito/backend-go/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	store db.Store
	route *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	route := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("status", validStatus)
		v.RegisterValidation("position", validPosition)
		v.RegisterValidation("gender", validGender)
		v.RegisterValidation("eventType", validEventType)
		v.RegisterValidation("visitType", validVisitType)
	}

	route.POST("/accounts", server.createAccount)
	route.GET("/accounts/:id", server.getAccount)
	route.GET("/accounts", server.listAccount)

	route.POST("/events", server.createEvent)
	route.GET("/events/:id", server.getEvent)
	route.GET("/events", server.listEvent)
	server.route = route

	return server
}

func (server *Server) Start(address string) error {
	return server.route.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
