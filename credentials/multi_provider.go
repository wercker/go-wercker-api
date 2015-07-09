package credentials

import "errors"

var (
	// ErrNoValidProvidersFound is the error returned when none of the wrapped
	// Providers returns a valid credentials object.
	ErrNoValidProvidersFound = errors.New("No Provider returned a valid credentials object")
)

// NewMultiProvider creates a new MultiProvider, which wraps providers.
func NewMultiProvider(providers ...Provider) *MultiProvider {
	return &MultiProvider{
		Providers: providers,
	}
}

// MultiProvider wraps Providers and will call all Providers when GetCredentials
// is called. It will ignore any errors from the wrapped providers.
type MultiProvider struct {
	Providers []Provider
}

// GetCredentials returns the first Creds object returned from p.Providers. It
// will ignore any errors coming from the Providers, however, it will return
// ErrNoValidProvidersFound if no wrapped Providers return a Creds object.
func (p *MultiProvider) GetCredentials() (*Creds, error) {
	for _, p := range p.Providers {
		c, err := p.GetCredentials()
		if err == nil {
			return c, nil
		}
	}

	return nil, ErrNoValidProvidersFound
}
