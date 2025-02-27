package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suhailmuhammed157/simple_bank/db_source"
)

type CreateTransferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,min=1"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) CreateTransfer(ctx *gin.Context) {

	var req CreateTransferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	validateFromAccount := server.validateAccount(ctx, req.FromAccountID, req.Currency)
	if !validateFromAccount {
		return
	}

	validateToAccount := server.validateAccount(ctx, req.ToAccountID, req.Currency)
	if !validateToAccount {
		return
	}

	args := db_source.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	account, err := server.store.TransferTx(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)

}

func (server *Server) validateAccount(ctx *gin.Context, id int64, currency string) bool {

	account, err := server.store.GetAccount(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	if account.Currency != currency {
		err := fmt.Errorf("currency %v does not match with one of the given accounts", currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}
	return true

}
