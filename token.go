package wercker

import "github.com/jtacoma/uritemplates"

// tokenTemplates contains all UriTemplates indexed by name.
var tokenTemplates = make(map[string]*uritemplates.UriTemplate)

func init() {
	addURITemplate(tokenTemplates, "CreateToken", "/api/v3/tokens")
	addURITemplate(tokenTemplates, "DeleteToken", "/api/v3/tokens/{tokenId}")
	addURITemplate(tokenTemplates, "GetToken", "/api/v3/tokens/{tokenId}")
	addURITemplate(tokenTemplates, "GetTokens", "/api/v3/tokens")
	addURITemplate(tokenTemplates, "UpdateToken", "/api/v3/tokens/{tokenId}")
}

// CreateTokenOptions are the options associated with Client.CreateToken
type CreateTokenOptions struct {
	Name string `json:"name"`
}

// CreateToken creates a new Token.
func (c *Client) CreateToken(options *CreateTokenOptions) (*Token, error) {
	method := "POST"
	template := tokenTemplates["CreateToken"]

	result := &Token{}
	err := c.Do(method, template, options, options, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetTokenOptions are the options associated with Client.GetToken
type GetTokenOptions struct {
	TokenID string `map:"tokenId"`
}

// GetToken retrieves a single Token.
func (c *Client) GetToken(options *GetTokenOptions) (*Token, error) {
	method := "GET"
	template := tokenTemplates["GetToken"]

	result := &Token{}
	err := c.Do(method, template, options, nil, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetTokensOptions are the options associated with Client.GetTokens
type GetTokensOptions struct {
}

// GetTokens gets all tokens for the current user.
func (c *Client) GetTokens(options *GetTokensOptions) ([]*TokenSummary, error) {
	method := "GET"
	template := tokenTemplates["GetTokens"]

	result := []*TokenSummary{}
	err := c.Do(method, template, options, nil, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateTokenOptions are the options associated with Client.UpdateToken
type UpdateTokenOptions struct {
	// Required
	TokenID string `map:"tokenId"`

	// Optional
	Name string `json:"name,omitempty"`
}

// UpdateToken updates a Token.
func (c *Client) UpdateToken(options *UpdateTokenOptions) (*Token, error) {
	method := "PATCH"
	template := tokenTemplates["UpdateToken"]

	result := &Token{}
	err := c.Do(method, template, options, options, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteTokenOptions are the options associated with Client.DeleteToken
type DeleteTokenOptions struct {
	TokenID string `map:"tokenId"`
}

// DeleteToken deletes a token
func (c *Client) DeleteToken(options *DeleteTokenOptions) error {
	method := "DELETE"
	template := tokenTemplates["DeleteToken"]

	err := c.Do(method, template, options, nil, nil)
	return err
}
