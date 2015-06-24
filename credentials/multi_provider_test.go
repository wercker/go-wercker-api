package credentials

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiProviderAssignable(t *testing.T) {
	var provider interface{}
	provider = &MultiProvider{}

	_, ok := provider.(Provider)
	assert.True(t, ok, "")
}

func TestMultProviderGet(t *testing.T) {
	provider := &MultiProvider{
		Providers: []Provider{
			&FakeProvider{Error: errors.New("Fake result1")},
			&FakeProvider{Error: errors.New("Fake result2")},
			&FakeProvider{Creds: &Creds{Token: "keepitgreen"}},
		},
	}

	creds, err := provider.GetCredentials()

	assert.Nil(t, err, "Expected err to be nil, but was %#v", err)

	if assert.NotNil(t, creds, "") {
		assert.Equal(t, "keepitgreen", creds.Token, "")
		assert.Empty(t, creds.Username, "")
		assert.Empty(t, creds.Password, "")
	}
}

func TestMultProviderInvalid(t *testing.T) {
	provider := &MultiProvider{
		Providers: []Provider{
			&FakeProvider{Error: errors.New("Fake result1")},
			&FakeProvider{Error: errors.New("Fake result2")},
			&FakeProvider{Error: errors.New("Fake result3")},
		},
	}

	creds, err := provider.GetCredentials()

	assert.Nil(t, creds, "")
	if assert.NotNil(t, err, "") {
		assert.Equal(t, ErrNoValidProvidersFound, err, "")
	}
}

func TestMultiProviderNewEmpty(t *testing.T) {
	provider := NewMultiProvider()

	assert.Equal(t, 0, len(provider.Providers), "")
}

func TestMultiProviderNewSingle(t *testing.T) {
	provider := NewMultiProvider(&StaticProvider{})

	assert.Equal(t, 1, len(provider.Providers), "")
}

func TestMultiProviderNewMulti(t *testing.T) {
	provider := NewMultiProvider(&StaticProvider{}, &StaticProvider{})

	assert.Equal(t, 2, len(provider.Providers), "")
}
