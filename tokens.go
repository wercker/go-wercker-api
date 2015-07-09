package wercker

import "errors"

func init() {
	addURITemplate("tokens.DeleteToken", "/api/v3/tokens/{tokenId}")
	addURITemplate("tokens.FetchTokens", "/api/v3/tokens")
	addURITemplate("tokens.GetToken", "/api/v3/tokens/{tokenId}")
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

	TokenID string `map:"tokenId"`
}

// Get a single Token.
func (c *tokensService) Get(options *GetTokenOptions) (*Token, error) {
	method := "GET"
	template := routes["tokens.GetToken"]

	result := &Token{}
	err := c.client.makeRequest(method, template, options, nil, nil, options.Config, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// FetchTokenOptions are the options associated with tokensService.Fetch
type FetchTokenOptions struct {
	Config *Config
}

// Fetch multiple Tokens.
func (c *tokensService) Fetch(options *FetchTokenOptions) ([]*TokenSummary, error) {
	method := "GET"
	template := routes["tokens.FetchTokens"]

	result := []*TokenSummary{}
	err := c.client.makeRequest(method, template, options, nil, nil, options.Config, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
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

	TokenID string `map:"tokenId"`
}

// Delete a token
func (c *tokensService) Delete(options *DeleteTokenOptions) error {
	method := "DELETE"
	template := routes["tokens.DeleteToken"]

	err := c.client.makeRequest(method, template, options, nil, nil, options.Config, nil)
	return err
}
