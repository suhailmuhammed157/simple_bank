package gapi

import (
	db_source "github.com/suhailmuhammed157/simple_bank/db_source/sqlc"
	"github.com/suhailmuhammed157/simple_bank/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user *db_source.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}

func convertAccount(account *db_source.Account) *pb.Account {
	return &pb.Account{
		Id:        account.ID,
		Owner:     account.Owner,
		Balance:   float32(account.Balance),
		Currency:  account.Currency,
		CreatedAt: timestamppb.New(account.CreatedAt),
	}
}
