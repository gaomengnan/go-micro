package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {

	maker, err := NewJWTMaker(RandomString(32))
	require.NoError(t, err)

	username := RandomOwner()

	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)

	require.NoError(t, err)

	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.UserName)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)

}

func TestExpiredJWTToken(t *testing.T) {

	maker, err := NewJWTMaker(RandomString(32))
	require.NoError(t, err)

	username := RandomOwner()

	token, err := maker.CreateToken(username, -time.Minute)

	require.NoError(t, err)

	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}
