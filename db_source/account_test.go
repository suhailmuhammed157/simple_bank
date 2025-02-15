package db_source

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/suhailmuhammed157/simple_bank/utils"
)

func createRandomAccount(t *testing.T) Account {
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

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	newAccount := createRandomAccount(t)

	account, err := testQueries.GetAccount(context.Background(), newAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, newAccount.ID, account.ID)
	require.Equal(t, newAccount.Owner, account.Owner)
	require.Equal(t, newAccount.Balance, account.Balance)
	require.Equal(t, newAccount.Currency, account.Currency)

	require.WithinDuration(t, newAccount.CreatedAt.Time, account.CreatedAt.Time, time.Second)
}
