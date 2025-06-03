package gapi

import (
	"context"
	"database/sql"
	"time"

	"github.com/hibiken/asynq"
	"github.com/lib/pq"
	"github.com/suhailmuhammed157/simple_bank/db_source"
	"github.com/suhailmuhammed157/simple_bank/pb"
	"github.com/suhailmuhammed157/simple_bank/utils"
	"github.com/suhailmuhammed157/simple_bank/val"
	"github.com/suhailmuhammed157/simple_bank/worker"
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

	args := db_source.CreateUserTxParams{
		CreateUserParams: db_source.CreateUserParams{
			Username:       req.GetUsername(),
			HashedPassword: hashedPassword,
			FullName:       req.GetFullName(),
			Email:          req.GetEmail(),
		},

		//after create user need to send email
		AfterCreateUser: func(user db_source.User) error {
			payload := &worker.PayloadSendVerifyEmail{Username: req.GetUsername()}
			opts := []asynq.Option{
				asynq.MaxRetry(10),
				asynq.ProcessIn(10 * time.Second),
				asynq.Queue("critical"),
			}
			return server.taskDistributor.DistributeTaskSendVerifyEmail(ctx, payload, opts...)
		},
	}

	txResult, err := server.store.CreateUserTx(ctx, args)
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
		User: convertUser(&txResult.User),
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

	payload, err := server.authenticateUser(ctx)

	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "token error: %v", err)
	}

	violations := validateGetUserDetailsRequest(payload.Username)

	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	user, err := server.store.GetUser(ctx, payload.Username)

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
	payload, err := server.authenticateUser(ctx)

	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "token error: %v", err)
	}
	violations := validateUpdateUserRequest(req, payload.Username)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	args := db_source.UpdateUserParams{

		FullName: sql.NullString{
			String: req.GetFullName(),
			Valid:  req.FullName != nil,
		},
		Email: sql.NullString{
			String: req.GetEmail(),
			Valid:  req.Email != nil,
		},
		Username: payload.Username,
	}

	if req.Password != nil {
		hashedPassword, err := utils.EncryptPassword(req.GetPassword())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to hash password %s", err)
		}
		args.HashedPassword = sql.NullString{
			String: hashedPassword,
			Valid:  req.Password != nil,
		}

		args.PasswordChangedAt = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
	}

	updatedUser, err := server.store.UpdateUser(ctx, args)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found %v", err)
		}

		return nil, status.Errorf(codes.Internal, "update user failed %v", err)
	}
	response := &pb.UpdateUserResponse{
		User: convertUser(&updatedUser),
	}
	return response, nil
}

func (server *Server) VerifyUser(ctx context.Context, req *pb.VerifyUserRequest) (*pb.VerifyUserResponse, error) {

	violations := validateVerifyUserRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	verifyEmail, err := server.store.UpdateVerifyEmail(ctx, db_source.UpdateVerifyEmailParams{
		ID:         int64(req.GetEmailId()),
		SecretCode: req.GetSecretCode(),
	})

	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "invalid secret code: %v", err)
	}

	updatedUser, err := server.store.UpdateUser(ctx, db_source.UpdateUserParams{
		IsUserVerified: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
		Username: verifyEmail.Username,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found %v", err)
		}
		return nil, status.Errorf(codes.Internal, "update user failed %v", err)
	}

	response := &pb.VerifyUserResponse{
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

func validateGetUserDetailsRequest(username string) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := val.ValidateUsername(username); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	return violations
}

func validateUpdateUserRequest(req *pb.UpdateUserRequest, username string) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := val.ValidateUsername(username); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if req.Password != nil {
		if err := val.ValidatePassword(*req.Password); err != nil {
			violations = append(violations, fieldViolation("password", err))
		}
	}
	if req.FullName != nil {
		if err := val.ValidateFullname(*req.FullName); err != nil {
			violations = append(violations, fieldViolation("full_name", err))
		}
	}
	if req.Email != nil {
		if err := val.ValidateEmail(*req.Email); err != nil {
			violations = append(violations, fieldViolation("email", err))
		}
	}

	return violations
}

func validateVerifyUserRequest(req *pb.VerifyUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := val.ValidateSecretCode(req.GetSecretCode()); err != nil {
		violations = append(violations, fieldViolation("secret_code", err))

	}

	return violations
}
