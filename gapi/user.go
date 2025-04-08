package gapi

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
	"github.com/suhailmuhammed157/simple_bank/db_source"
	"github.com/suhailmuhammed157/simple_bank/pb"
	"github.com/suhailmuhammed157/simple_bank/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	hashedPassword, err := utils.EncryptPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to hash password %s", err)
	}

	args := db_source.CreateUserParams{
		Username:       req.GetUsername(),
		HashedPassword: hashedPassword,
		FullName:       req.GetFullName(),
		Email:          req.GetEmail(),
	}

	user, err := server.store.CreateUser(ctx, args)
	if err != nil {
		if pqError, ok := err.(*pq.Error); ok {
			switch pqError.Code.Name() {

			case "unique_violation":
				return nil, status.Errorf(codes.Unauthenticated, "user already exists: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "Failed to create user %s", err)
	}

	usr := &pb.CreateUserResponse{
		User: convertUser(&user),
	}

	return usr, nil

}

func (server *Server) Login(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {

	user, err := server.store.GetUser(ctx, req.GetUsername())

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user or password not found: %s", err)

		}
		return nil, status.Errorf(codes.NotFound, "user or password not found: %s", err)

	}

	validateRes := utils.ValidatePassword(user.HashedPassword, req.GetPassword())

	if !validateRes {
		return nil, status.Errorf(codes.NotFound, "user or password incorrect: %s", err)
	}

	access_token, accessPayload, err := server.tokenMaker.CreateToken(user.Username, server.config.TokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "user or password incorrect: %s", err)

	}

	refresh_token, refreshPayload, err := server.tokenMaker.CreateToken(user.Username, server.config.RefreshTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to generate token: %s", err)

	}

	mtd := server.extractMetadata(ctx)

	server.store.CreateSession(ctx, db_source.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refresh_token,
		UserAgent:    mtd.UserAgent,
		ClientIp:     mtd.ClientIp,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})

	response := &pb.LoginUserResponse{
		AccessToken:           access_token,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshToken:          refresh_token,
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
		User:                  convertUser(&user),
	}

	return response, nil
}
