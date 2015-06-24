package credentials

import (
	"fmt"
	"os"
)

var ErrWerckerTokenEnvNotFound = fmt.Errorf("WERCKER_TOKEN not found in environment")

type EnvProvider struct{}

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
