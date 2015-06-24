package credentials

type Provider interface {
	GetCredentials() (*Creds, error)
}

type Creds struct {
	Token    string
	Username string
	Password string
}

// func AnonymousCreds() *Creds {
// 	return &Creds{}
// }
