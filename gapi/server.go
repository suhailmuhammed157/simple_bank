package gapi

import (
	"github.com/suhailmuhammed157/simple_bank/db_source"
	"github.com/suhailmuhammed157/simple_bank/pb"
	"github.com/suhailmuhammed157/simple_bank/token"
	"github.com/suhailmuhammed157/simple_bank/utils"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	store      *db_source.Store
	tokenMaker token.Maker
	config     *utils.Config
}

func NewServer(config *utils.Config, store *db_source.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.SymmetricKey)
	if err != nil {
		return nil, err
	}

	server := &Server{store: store, tokenMaker: tokenMaker, config: config}

	return server, nil
}
