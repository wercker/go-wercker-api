package credentials

type Provider interface {
	GetCredentials() (*Creds, error)
}

type Creds struct {
	Token    string
	Username string
	Password string
}

func Anonymous() Provider {
	return &StaticProvider{}
}

func Token(token string) Provider {
	return &StaticProvider{Token: token}
}

func UsernamePassword(username string, password string) Provider {
	return &StaticProvider{Username: username, Password: password}
}
