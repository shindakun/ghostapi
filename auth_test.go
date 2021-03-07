package ghostapi

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/shindakun/ghostapi/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientFromAuth(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	c, err := NewClientFromAuth(mock.Domain, "", mock.Key)
	require.NoError(t, err)
	assert.IsType(t, &Client{}, c)
	//assert.Equal(t, c.ctx, auth)
	assert.Equal(t, c.APIKey, "somekey")
	assert.Equal(t, c.URL, "https://shindakun.net/ghost/api/v3/content/posts/?key=somekey")
	assert.Equal(t, c.Page, 1)
}

func TestClientFromAuthWithVersion(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	c, err := NewClientFromAuth(mock.Domain, mock.V2, mock.Key)
	require.NoError(t, err)
	assert.IsType(t, &Client{}, c)
	//assert.Equal(t, c.ctx, auth)
	assert.Equal(t, c.APIKey, "somekey")
	assert.Equal(t, c.URL, "https://shindakun.net/ghost/api/v2/content/posts/?key=somekey")
	assert.Equal(t, c.Page, 1)
}
