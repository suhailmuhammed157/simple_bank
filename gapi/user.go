package gapi

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
	"github.com/suhailmuhammed157/simple_bank/db_source"
	"github.com/suhailmuhammed157/simple_bank/pb"
	"github.com/suhailmuhammed157/simple_bank/utils"
	"github.com/suhailmuhammed157/simple_bank/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	hashedPassword, err := utils.EncryptPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to hash password %s", err)
	}

	if violations := validateCreateUserRequest(req); len(violations) > 0 {
		return nil, invalidArgumentError(violations)
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

	violations := validateLoginRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

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

func (server *Server) GetUserDetails(ctx context.Context, req *pb.GetUserDetailsRequest) (*pb.GetUserDetailsResponse, error) {

	violations := validateGetUserDetailsRequest(req)

	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	user, err := server.store.GetUser(ctx, req.GetUsername())

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found: %s", err)

		}
		return nil, status.Errorf(codes.NotFound, "user not found: %s", err)

	}
	userResponse := &pb.GetUserDetailsResponse{
		User: convertUser(&user),
	}
	return userResponse, nil
}

func (server *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	violations := validateUpdateUserRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	args := db_source.UpdateUserParams{
		HashedPassword: sql.NullString{
			String: req.GetPassword(),
			Valid:  req.Password != nil,
		},
		FullName: sql.NullString{
			String: req.GetFullName(),
			Valid:  req.FullName != nil,
		},
		Email: sql.NullString{
			String: req.GetEmail(),
			Valid:  req.Email != nil,
		},
		Username: req.GetUsername(),
	}
	updatedUser, err := server.store.UpdateUser(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "update user failed %v", err)
	}
	response := &pb.UpdateUserResponse{
		User: convertUser(&updatedUser),
	}
	return response, nil
}

func validateCreateUserRequest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := val.ValidateUsername(req.Username); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if err := val.ValidatePassword(req.Password); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	if err := val.ValidateFullname(req.FullName); err != nil {
		violations = append(violations, fieldViolation("full_name", err))
	}

	if err := val.ValidateEmail(req.Email); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}

	return violations
}

func validateLoginRequest(req *pb.LoginUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := val.ValidateUsername(req.Username); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if err := val.ValidatePassword(req.Password); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}
	return violations
}

func validateGetUserDetailsRequest(req *pb.GetUserDetailsRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := val.ValidateUsername(req.Username); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	return violations
}

func validateUpdateUserRequest(req *pb.UpdateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := val.ValidateUsername(req.Username); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	return violations
}
