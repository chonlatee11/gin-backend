package api

import (
	db "github.com/chonlatee11/gin-backend/db/sqlc"
	"github.com/chonlatee11/gin-backend/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config util.Config
	router *gin.Engine
	store  db.Store
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	server := Server{
		config: config,
		store: *store,
	}

	server.setupRouter()
	return &server, nil
}

func (s *Server) setupRouter() {

	router := gin.Default()

	apiv1Route := router.Group("/api/v1")

	apiv1Route.POST("/register", s.createUser)
	apiv1Route.POST("/login", s.loginUser)

	s.router = router
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
