package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/suhailmuhammed157/simple_bank/db_source"
)

type Server struct {
	store  *db_source.Store
	router *gin.Engine
}

func NewServer(store *db_source.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validateCurrency)
	}

	router.POST("/users", server.CreateUser)
	router.POST("/accounts", server.CreateAccount)
	router.GET("/accounts/:id", server.GetAccountDetails)
	router.GET("/accounts", server.ListAccounts)
	router.POST("/transfers", server.CreateTransfer)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
