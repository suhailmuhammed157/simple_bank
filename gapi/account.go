package gapi

import (
	"context"

	db_source "github.com/suhailmuhammed157/simple_bank/db_source/sqlc"
	"github.com/suhailmuhammed157/simple_bank/pb"
	"github.com/suhailmuhammed157/simple_bank/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateAccountRequest struct {
	Currency string `json:"currency" binding:"required,currency"`
}

func (server *Server) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {

	authPayload, err := server.authenticateUser(ctx)

	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "token error: %v", err)
	}

	if violations := validateCreateAccountRequest(req); len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	args := db_source.CreateAccountParams{
		Owner:    authPayload.Username,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, args)
	if err != nil {

		if code := db_source.ErrorCode(err); code == db_source.UniqueViolation || code == db_source.ForeignKeyViolation {
			return nil, status.Errorf(codes.Unauthenticated, "account already exists: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "Failed to create account %s", err)

	}

	acnt := &pb.CreateAccountResponse{
		Account: convertAccount(&account),
	}

	return acnt, nil

}

func (server *Server) GetAccountDetails(ctx context.Context, req *pb.Empty) (*pb.GetAccountDetailsResponse, error) {

	authPayload, err := server.authenticateUser(ctx)

	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "token error: %v", err)
	}

	account, err := server.store.GetAccount(ctx, authPayload.Username)
	if err != nil {
		if err == db_source.NoRowFound {
			return nil, status.Errorf(codes.Unauthenticated, "account not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

	acnt := &pb.GetAccountDetailsResponse{
		Account: convertAccount(&account),
	}

	return acnt, nil

}

func (server *Server) ListAccounts(ctx context.Context, req *pb.ListAccountRequest) (*pb.ListAccountResponse, error) {

	authPayload, err := server.authenticateUser(ctx)

	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "token error: %v", err)
	}
	accounts, err := server.store.ListAccounts(ctx, db_source.ListAccountsParams{
		Owner:  authPayload.Username,
		Limit:  req.PageSize,
		Offset: (req.PageId - 1) * req.PageSize,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	var pb_accounts []*pb.Account

	for _, acc := range accounts {
		pb_acc := convertAccount(&acc)
		pb_accounts = append(pb_accounts, pb_acc)
	}

	acnts := &pb.ListAccountResponse{
		Accounts: pb_accounts,
	}

	return acnts, nil

}

func validateCreateAccountRequest(req *pb.CreateAccountRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := val.ValidateCurrency(req.Currency); err != nil {
		violations = append(violations, fieldViolation("currency", err))
	}

	return violations
}
