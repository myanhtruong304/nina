package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPasswor(t *testing.T) {
	pass := "Myanhtruong304"

	hashedPassword, err := HashPassword(pass)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	wrongPassword := "Myanhtruong"
	err = CheckPassword(wrongPassword, hashedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
