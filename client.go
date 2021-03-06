package ghostapi

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var (

	// BaseURL to start
	BaseURL = "https://%s"

	//APIPath is the default Ghost API path for posts without the version number
	APIPath = "/ghost/api/%s/content/posts/"

	//LatestAPI is the currently supported version
	LatestAPI = "v3"
)

func makeURL(domain, ver, k string) string {
	if ver == "" {
		ver = LatestAPI
	}
	host := fmt.Sprintf(BaseURL, domain)
	path := fmt.Sprintf(APIPath, ver)
	url := fmt.Sprintf("%s%s?key=%s", host, path, k)
	return url
}

// Client struct
type Client struct {
	URL    string
	APIKey string
	Page   int
}

// NewClient takes a domain name, API version, and API key strings and returns a *Client
func NewClient(domain, ver, k string) *Client {
	url := makeURL(domain, ver, k)
	return &Client{
		URL:    url,
		APIKey: k,
		Page:   1,
	}
}

// Get returns []byte
func (c *Client) Get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.do(req)
}

// Do returns []byte, uses c.Page for paging
func (c *Client) do(r *http.Request) ([]byte, error) {
	if c.Page > 1 {
		i := strconv.Itoa(c.Page)
		q := r.URL.Query()
		q.Add("page", i)
		r.URL.RawQuery = q.Encode()
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
