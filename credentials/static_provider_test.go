package credentials

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStaticProviderAssignable(t *testing.T) {
	var provider interface{}
	provider = &StaticProvider{}

	_, ok := provider.(Provider)
	assert.True(t, ok, "")
}

func TestStaticProviderUsernamePassword(t *testing.T) {
	provider := &StaticProvider{Username: "username", Password: "password"}

	creds, err := provider.GetCredentials()
	assert.NoError(t, err, "")

	if assert.NotNil(t, creds, "") {
		assert.Empty(t, creds.Token, "")
		assert.Equal(t, "username", creds.Username, "")
		assert.Equal(t, "password", creds.Password, "")
	}
}

func TestStaticProviderToken(t *testing.T) {
	provider := &StaticProvider{Token: "token"}

	creds, err := provider.GetCredentials()
	assert.NoError(t, err, "")

	if assert.NotNil(t, creds, "") {
		assert.Equal(t, "token", creds.Token, "")
		assert.Empty(t, creds.Username, "")
		assert.Empty(t, creds.Password, "")
	}
}

func TestStaticProviderEmpty(t *testing.T) {
	provider := &StaticProvider{}

	creds, err := provider.GetCredentials()
	assert.NoError(t, err, "")

	if assert.NotNil(t, creds, "") {
		assert.Empty(t, creds.Token, "")
		assert.Empty(t, creds.Username, "")
		assert.Empty(t, creds.Password, "")
	}
}
