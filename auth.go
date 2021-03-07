package ghostapi

import "errors"

// NewClientFromAuth gets domain, API version and token from store
// and returns a *Client.
func NewClientFromAuth(domain, ver, token string) (*Client, error) {
	if token == "" {
		return nil, errors.New("no token found in auth")
	}
	if domain == "" {
		return nil, errors.New("no domain found in auth")
	}
	return NewClient(domain, ver, token), nil
}

// Validate attempts to retreive a single page of posts returns an
// error if not succesful
func (c *Client) Validate() error {
	if _, err := c.Get(c.URL); err != nil {
		return err
	}
	return nil
}
