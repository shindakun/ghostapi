package ghostapi

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var (

	// BaseURL to start, https or die
	BaseURL = "https://%s"

	//APIPath is the default Ghost API path without the version number
	APIPath = "/ghost/api/%s/content/posts/"

	//LatestAPI is the currently supported version
	LatestAPI = "v3"
)

// Client struct
type Client struct {
	URL    string
	APIKey string
	Page   int
}

func makeURL(domain, ver, k string) string {
	host := fmt.Sprintf(BaseURL, domain)
	path := fmt.Sprintf(APIPath, ver)
	url := fmt.Sprintf("%s%s?key=%s", host, path, k)
	return url
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

// Get returns an http.Request ready to use
func (c *Client) Get(url string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// Do returns the cx
func (c *Client) Do(r *http.Request) ([]byte, error) {
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

	fmt.Println(r.URL.String())
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
