package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashedPassword(t *testing.T) {
	password := "secret"
	hashedPassword, err := EncryptPassword(password)
	require.NoError(t, err)
	isPasswordValid := ValidatePassword(hashedPassword, password)

	require.NotEqual(t, hashedPassword, password)
	require.True(t, isPasswordValid)

	hashedPassword2, err := EncryptPassword(password)
	require.NoError(t, err)
	require.NotEqual(t, hashedPassword, hashedPassword2)

}
