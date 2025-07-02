package gapi

import (
	"context"
	"fmt"
	"time"

	db_source "github.com/suhailmuhammed157/simple_bank/db_source/sqlc"
	"github.com/suhailmuhammed157/simple_bank/pb"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type NewTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type NewTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func (server *Server) IssueNewToken(ctx context.Context, req *pb.NewTokenRequest) (*pb.NewTokenResponse, error) {

	if violations := validateIssueNewTokenRequest(req); len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	refreshPayload, err := server.tokenMaker.VerifyToken(req.RefreshToken)

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "failed to renew token: %v", err)
	}

	session, err := server.store.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		if err == db_source.NoRowFound {
			return nil, status.Errorf(codes.NotFound, "invalid token: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}

	if session.IsBlocked {
		return nil, status.Errorf(codes.Unauthenticated, "blocked session")
	}

	if session.Username != refreshPayload.Username {
		return nil, status.Errorf(codes.Unauthenticated, "invalid user")

	}

	access_token, accessPayload, err := server.tokenMaker.CreateToken(refreshPayload.Username, server.config.TokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}

	response := &pb.NewTokenResponse{
		AccessToken:          access_token,
		AccessTokenExpiresAt: timestamppb.New(accessPayload.ExpiredAt),
	}
	return response, nil

}

func validateIssueNewTokenRequest(req *pb.NewTokenRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if len(req.RefreshToken) == 0 {
		violations = append(violations, fieldViolation("refresh_token", fmt.Errorf("refresh_token is required")))
	}

	return violations
}
