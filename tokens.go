package wercker

import "errors"

func init() {
	// Add templates to the route map
	addURITemplate("tokens.GetToken", "")
}

type tokensService struct {
	client *Client
}

type CreateTokenOptions struct {
	Config *Config
}

func (c *tokensService) Create(options *CreateTokenOptions) (*Token, error) {
	return nil, errors.New("not implemented")
}

type GetTokenOptions struct {
	Config *Config
}

func (c *tokensService) Get(options *GetTokenOptions) (*Token, error) {
	return nil, errors.New("not implemented")
}

type FetchTokenOptions struct {
	Config *Config
}

func (c *tokensService) Fetch(options *FetchTokenOptions) ([]*Token, error) {
	return nil, errors.New("not implemented")
}

type UpdateTokenOptions struct {
	Config *Config
}

func (c *tokensService) Update(options *UpdateTokenOptions) (*Token, error) {
	return nil, errors.New("not implemented")
}

type DeleteTokenOptions struct {
	Config *Config
}

func (c *tokensService) Delete(options *DeleteTokenOptions) error {
	return errors.New("not implemented")
}
