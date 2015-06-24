package credentials

import "errors"

var (
	ErrNoValidProvidersFound = errors.New("")
)

func NewMultiProvider(providers ...Provider) *MultiProvider {
	return &MultiProvider{
		Providers: providers,
	}
}

type MultiProvider struct {
	Providers []Provider
}

func (p *MultiProvider) GetCredentials() (*Creds, error) {
	for _, p := range p.Providers {
		c, err := p.GetCredentials()
		if err == nil {
			return c, nil
		}
	}

	return nil, ErrNoValidProvidersFound
}
