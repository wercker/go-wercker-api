package credentials

// StaticProvider always returns a Creds object filled with Username, Password,
// and Token.
type StaticProvider struct {
	Username string
	Password string
	Token    string
}

// GetCredentials returns a Creds object based on p.Username, p.Password, and
// p.Token. It never returns a error.
func (p *StaticProvider) GetCredentials() (*Creds, error) {
	return &Creds{
		Username: p.Username,
		Password: p.Password,
		Token:    p.Token,
	}, nil
}
