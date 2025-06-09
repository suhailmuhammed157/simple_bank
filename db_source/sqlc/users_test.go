package db_source

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
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
	user, err := testStore.CreateUser(context.Background(), args)

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

	account, err := testStore.GetUser(context.Background(), newUser.Username)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, newUser.Username, account.Username)
	require.Equal(t, newUser.HashedPassword, account.HashedPassword)
	require.Equal(t, newUser.FullName, account.FullName)
	require.Equal(t, newUser.Email, account.Email)

	require.WithinDuration(t, newUser.PasswordChangedAt, account.PasswordChangedAt, time.Second)
	require.WithinDuration(t, newUser.CreatedAt, account.CreatedAt, time.Second)
}

func TestUpdateUserEmail(t *testing.T) {
	account := createRandomUser(t)

	newEmail := utils.RandomEmail()

	args := UpdateUserParams{
		Email: pgtype.Text{
			Valid:  true,
			String: newEmail,
		},
		Username: account.Username,
	}

	updatedAccount, err := testStore.UpdateUser(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, updatedAccount.Username, account.Username)
	require.Equal(t, updatedAccount.HashedPassword, account.HashedPassword)
	require.Equal(t, updatedAccount.FullName, account.FullName)
	require.NotEqual(t, updatedAccount.Email, account.Email)
	require.Equal(t, updatedAccount.Email, newEmail)

	require.WithinDuration(t, account.PasswordChangedAt, account.PasswordChangedAt, time.Second)
	require.WithinDuration(t, account.CreatedAt, account.CreatedAt, time.Second)
}
