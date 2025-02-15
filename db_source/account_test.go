package db_source

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suhailmuhammed157/simple_bank/utils"
)

func TestCreateAccount(t *testing.T) {
	args := CreateAccountParams{
		Owner:    utils.RandomString(6),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, account.Owner, args.Owner)
	require.Equal(t, account.Balance, args.Balance)
	require.Equal(t, account.Currency, args.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}
