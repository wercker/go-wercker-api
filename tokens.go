package wercker

import "errors"

func init() {
	// Add templates to the route map
	addURITemplate("tokens.GetToken", "")
}

type CreateTokenOptions struct {
	*Options
}
type GetTokenOptions struct {
	*Options
}
type FetchTokenOptions struct {
	*Options
}
type UpdateTokenOptions struct {
	*Options
}
type DeleteTokenOptions struct {
	*Options
}

func (c *Client) CreateToken(options *CreateTokenOptions) (*Token, error) {
	return nil, errors.New("not implemented")
}

func (c *Client) GetToken(options *GetTokenOptions) (*Token, error) {
	return nil, errors.New("not implemented")
}

func (c *Client) FetchTokens(options *FetchTokenOptions) ([]*Token, error) {
	return nil, errors.New("not implemented")
}

func (c *Client) UpdateToken(options *UpdateTokenOptions) (*Token, error) {
	return nil, errors.New("not implemented")
}

func (c *Client) DeleteToken(options *DeleteTokenOptions) error {
	return errors.New("not implemented")
}
