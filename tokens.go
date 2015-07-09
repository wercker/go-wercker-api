package wercker

import "errors"

func init() {
	// Add templates to the route map
	addURITemplate("tokens.GetToken", "")
}

type tokensService struct {
	client *Client
}

// CreateTokenOptions are the options associated with tokensService.Create
type CreateTokenOptions struct {
	Config *Config
}

// Create a new Token.
func (c *tokensService) Create(options *CreateTokenOptions) (*Token, error) {
	return nil, errors.New("not implemented")
}

// GetTokenOptions are the options associated with tokensService.Get
type GetTokenOptions struct {
	Config *Config
}

// Get a single Token.
func (c *tokensService) Get(options *GetTokenOptions) (*Token, error) {
	return nil, errors.New("not implemented")
}

// FetchTokenOptions are the options associated with tokensService.Fetch
type FetchTokenOptions struct {
	Config *Config
}

// Fetch multiple Tokens.
func (c *tokensService) Fetch(options *FetchTokenOptions) ([]*Token, error) {
	return nil, errors.New("not implemented")
}

// UpdateTokenOptions are the options associated with tokensService.Update
type UpdateTokenOptions struct {
	Config *Config
}

// Update a Token.
func (c *tokensService) Update(options *UpdateTokenOptions) (*Token, error) {
	return nil, errors.New("not implemented")
}

// DeleteTokenOptions are the options associated with tokensService.Delete
type DeleteTokenOptions struct {
	Config *Config
}

// Delete a token
func (c *tokensService) Delete(options *DeleteTokenOptions) error {
	return errors.New("not implemented")
}
