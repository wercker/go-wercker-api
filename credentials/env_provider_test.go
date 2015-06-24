package credentials

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvProviderAssignable(t *testing.T) {
	var provider interface{}
	provider = &EnvProvider{}

	_, ok := provider.(Provider)
	assert.True(t, ok, "")
}

func TestEnvProviderGet(t *testing.T) {
	os.Clearenv()
	os.Setenv("WERCKER_TOKEN", "keepitgreen")

	provider := &EnvProvider{}

	creds, err := provider.GetCredentials()

	assert.NoError(t, err, "")

	if assert.NotNil(t, creds, "") {
		assert.Equal(t, "keepitgreen", creds.Token, "")
		assert.Empty(t, creds.Username, "")
		assert.Empty(t, creds.Password, "")
	}
}

func TestEnvProviderEmptyEnv(t *testing.T) {
	os.Clearenv()

	provider := &EnvProvider{}

	creds, err := provider.GetCredentials()

	assert.Nil(t, creds, "")
	if assert.Error(t, err, "") {
		assert.Equal(t, ErrWerckerTokenEnvNotFound, err, "")
	}
}
