package api

import (
	"fmt"

	db "github.com/TTKirito/backend-go/db/sqlc"
	"github.com/TTKirito/backend-go/token"
	"github.com/TTKirito/backend-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     utils.Config
	store      db.Store
	route      *gin.Engine
	tokenMaker token.Maker
}

func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker :%v", err)
	}

	server := &Server{config: config, store: store, tokenMaker: tokenMaker}
	route := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("status", validStatus)
		v.RegisterValidation("position", validPosition)
		v.RegisterValidation("gender", validGender)
		v.RegisterValidation("eventType", validEventType)
		v.RegisterValidation("visitType", validVisitType)
	}

	route.POST("/users", server.createUser)
	route.POST("/users/login", server.loginUser)

	authRoutes := route.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccount)

	authRoutes.POST("/events", server.createEvent)
	authRoutes.GET("/events/:id", server.getEvent)
	authRoutes.GET("/events", server.listEvent)
	server.route = route

	return server, nil
}

func (server *Server) Start(address string) error {
	return server.route.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
