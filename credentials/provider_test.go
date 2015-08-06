package credentials

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToken(t *testing.T) {
	token := "secret token"

	provider := Token(token)
	cred, err := provider.GetCredentials()

	require.NoError(t, err, "")
	require.NotNil(t, cred, "")
	assert.Equal(t, "secret token", cred.Token, "")
	assert.Empty(t, cred.Username, "")
	assert.Empty(t, cred.Password, "")
}

func TestUsernamePassword(t *testing.T) {
	username := "secret username"
	password := "secret password"

	provider := UsernamePassword(username, password)
	cred, err := provider.GetCredentials()

	require.NoError(t, err, "")
	require.NotNil(t, cred, "")
	assert.Equal(t, "secret username", cred.Username, "")
	assert.Equal(t, "secret password", cred.Password, "")
	assert.Empty(t, cred.Token, "")
}

func TestAnonymous(t *testing.T) {
	provider := Anonymous()
	cred, err := provider.GetCredentials()

	require.NoError(t, err, "")
	require.NotNil(t, cred, "")
	assert.Empty(t, cred.Token, "")
	assert.Empty(t, cred.Username, "")
	assert.Empty(t, cred.Password, "")
}
