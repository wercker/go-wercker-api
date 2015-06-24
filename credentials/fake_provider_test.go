package credentials

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type FakeProvider struct {
	Error error
	Creds *Creds
}

func (p *FakeProvider) GetCredentials() (*Creds, error) {
	if p.Creds != nil {
		return p.Creds, nil
	}

	if p.Error != nil {
		return nil, p.Error
	}

	panic("Specify either Creds or Error")
}

func TestFakeProviderAssignable(t *testing.T) {
	var provider interface{}
	provider = &FakeProvider{}

	_, ok := provider.(Provider)
	assert.True(t, ok, "")
}
