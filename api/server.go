package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/suhailmuhammed157/simple_bank/db_source"
	"github.com/suhailmuhammed157/simple_bank/token"
	"github.com/suhailmuhammed157/simple_bank/utils"
)

type Server struct {
	store      *db_source.Store
	router     *gin.Engine
	tokenMaker token.Maker
	config     *utils.Config
}

func NewServer(config *utils.Config, store *db_source.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.SymmetricKey)
	if err != nil {
		return nil, err
	}

	server := &Server{store: store, tokenMaker: tokenMaker, config: config}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validateCurrency)
	}

	server.setupApiRoutes()

	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) setupApiRoutes() {
	router := gin.Default()
	router.POST("/users", server.CreateUser)
	router.POST("/users/login", server.Login)
	router.POST("/accounts", server.CreateAccount)
	router.GET("/accounts/:id", server.GetAccountDetails)
	router.GET("/accounts", server.ListAccounts)
	router.POST("/transfers", server.CreateTransfer)
	server.router = router
}
