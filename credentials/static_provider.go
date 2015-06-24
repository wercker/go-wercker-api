package credentials

type StaticProvider struct {
	Username string
	Password string
	Token    string
}

func (p *StaticProvider) GetCredentials() (*Creds, error) {
	return &Creds{
		Username: p.Username,
		Password: p.Password,
		Token:    p.Token,
	}, nil
}
