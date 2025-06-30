package gapi

import (
	"context"

	db_source "github.com/suhailmuhammed157/simple_bank/db_source/sqlc"
	"github.com/suhailmuhammed157/simple_bank/pb"
	"github.com/suhailmuhammed157/simple_bank/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		Account: &pb.Account{
			Id:        account.ID,
			Owner:     account.Owner,
			Balance:   float32(account.Balance),
			Currency:  account.Currency,
			CreatedAt: timestamppb.New(account.CreatedAt),
		},
	}

	return acnt, nil

}

// type GetAccountParam struct {
// 	Id int64 `uri:"id" binding:"required"`
// }

// func (server *Server) GetAccountDetails(ctx *gin.Context) {

// 	var req GetAccountParam
// 	if err := ctx.ShouldBindUri(&req); err != nil {

// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	account, err := server.store.GetAccount(ctx, req.Id)
// 	if err != nil {
// 		if err == db_source.NoRowFound {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	authPayload := ctx.MustGet(authorizationPayload).(*token.Payload)
// 	if account.Owner != authPayload.Username {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("account does not belong to the current user ")))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, account)

// }

// type ListAccountsParams struct {
// 	PageId   int32 `form:"page_id" binding:"required,min=1"`
// 	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
// }

// func (server *Server) ListAccounts(ctx *gin.Context) {

// 	authPayload := ctx.MustGet(authorizationPayload).(*token.Payload)

// 	var req ListAccountsParams
// 	if err := ctx.ShouldBindQuery(&req); err != nil {

// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	accounts, err := server.store.ListAccounts(ctx, db_source.ListAccountsParams{
// 		Owner:  authPayload.Username,
// 		Limit:  req.PageSize,
// 		Offset: (req.PageId - 1) * req.PageSize,
// 	})
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, accounts)

// }

func validateCreateAccountRequest(req *pb.CreateAccountRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := val.ValidateCurrency(req.Currency); err != nil {
		violations = append(violations, fieldViolation("currency", err))
	}

	return violations
}
