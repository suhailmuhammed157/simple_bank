package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/require"
	"github.com/suhailmuhammed157/simple_bank/utils"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)

	username := utils.RandomOwner()
	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := time.Now().Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)

}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)

	username := utils.RandomOwner()
	duration := -time.Minute
	token, err := maker.CreateToken(username, duration)

	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.Empty(t, payload)

}

func TestExpiredJWTTokenAlgoNone(t *testing.T) {

	username := utils.RandomOwner()
	duration := time.Minute

	payload, err := NewPayload(username, duration)
	require.NoError(t, err)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)

	require.NoError(t, err)
	require.NotEmpty(t, token)

	maker, err := NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)

}
