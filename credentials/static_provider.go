package credentials

// StaticProvider always returns a Credentials object filled with Username,
// Password, and Token.
type StaticProvider struct {
	Username string
	Password string
	Token    string
}

// GetCredentials returns a Credentials object based on p.Username, p.Password,
// and p.Token. It never returns a error.
func (p *StaticProvider) GetCredentials() (*Credentials, error) {
	return &Credentials{
		Username: p.Username,
		Password: p.Password,
		Token:    p.Token,
	}, nil
}
