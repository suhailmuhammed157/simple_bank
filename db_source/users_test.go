package db_source

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/suhailmuhammed157/simple_bank/utils"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := utils.EncryptPassword(utils.RandomString(6))
	require.NoError(t, err)
	args := CreateUserParams{
		Username:       utils.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       utils.RandomOwner(),
		Email:          utils.RandomEmail(),
	}
	user, err := testQueries.CreateUser(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.Username, args.Username)
	require.Equal(t, user.HashedPassword, args.HashedPassword)
	require.Equal(t, user.FullName, args.FullName)
	require.Equal(t, user.Email, args.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	newUser := createRandomUser(t)

	account, err := testQueries.GetUser(context.Background(), newUser.Username)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, newUser.Username, account.Username)
	require.Equal(t, newUser.HashedPassword, account.HashedPassword)
	require.Equal(t, newUser.FullName, account.FullName)
	require.Equal(t, newUser.Email, account.Email)

	require.WithinDuration(t, newUser.PasswordChangedAt, account.PasswordChangedAt, time.Second)
	require.WithinDuration(t, newUser.CreatedAt, account.CreatedAt, time.Second)
}
