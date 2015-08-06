package credentials

// Provider is an abstraction for retrieving wercker Credentials.
type Provider interface {
	GetCredentials() (*Credentials, error)
}

// Credentials should be either a Token, or Username/Password.
type Credentials struct {
	Token    string
	Username string
	Password string
}

// Anonymous returns a provider which will always return a empty Credentials
// object.
func Anonymous() Provider {
	return &StaticProvider{}
}

// Token returns a Provider which will always return a Credentials object with
// Token set to token.
func Token(token string) Provider {
	return &StaticProvider{Token: token}
}

// UsernamePassword returns a Provider which will always return a Credentials
// object with Username and Password set to username and password.
func UsernamePassword(username string, password string) Provider {
	return &StaticProvider{Username: username, Password: password}
}
