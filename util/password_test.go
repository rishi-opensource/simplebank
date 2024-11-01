package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	pswd := RandomString(6)

	hashedPswd1, err := HashPassword(pswd)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPswd1)

	err = CheckPassword(pswd, hashedPswd1)
	require.NoError(t, err)

	wrongPswd := RandomString(6)
	err = CheckPassword(wrongPswd, hashedPswd1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPswd2, err := HashPassword(pswd)
	require.NoError(t, err)
	require.NotEqual(t, hashedPswd1, hashedPswd2)
}
