package credentials

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type FakeProvider struct {
	Error       error
	Credentials *Credentials
}

func (p *FakeProvider) GetCredentials() (*Credentials, error) {
	if p.Credentials != nil {
		return p.Credentials, nil
	}

	if p.Error != nil {
		return nil, p.Error
	}

	panic("Specify either Credentials or Error")
}

func TestFakeProviderAssignable(t *testing.T) {
	var provider interface{}
	provider = &FakeProvider{}

	_, ok := provider.(Provider)
	assert.True(t, ok, "")
}
