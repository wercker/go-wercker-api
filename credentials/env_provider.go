package credentials

import (
	"fmt"
	"os"
)

// ErrWerckerTokenEnvNotFound is the error when the WERCKER_TOKEN env var is not
// found
var ErrWerckerTokenEnvNotFound = fmt.Errorf("WERCKER_TOKEN not found in environment")

// EnvProvider fetches the wercker token from the environment. By default it
// uses $WERCKER_TOKEN.
type EnvProvider struct{}

// GetCredentials returns the value of the WERCKER_TOKEN environment variable.
// If this is not set, it will return a ErrWerckerTokenEnvNotFound error.
func (p *EnvProvider) GetCredentials() (*Creds, error) {
	t := os.Getenv("WERCKER_TOKEN")

	if t != "" {
		c := Creds{
			Token: t,
		}
		return &c, nil
	}

	return nil, ErrWerckerTokenEnvNotFound
}
