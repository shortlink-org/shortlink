package application

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	rpc "github.com/batazor/shortlink/internal/metadata/domain"
)

var metaMock = rpc.Meta{
	ImageURL:    "",
	Keywords:    "",
	Description: "GitHub brings together the world’s largest community of developers to discover, share, and build better software. From open source projects to private team repositories, we’re your all-in-one platform for collaborative development.",
}

func TestSet(t *testing.T) { //nolint unused
	ctx := context.Background()
	url := "https://github.com/"

	r := Repository{}
	meta, err := r.Set(ctx, url)
	assert.Nil(t, err, "Error get body")

	// Check content
	assert.Equal(t, metaMock.ImageURL, meta.ImageURL)
	assert.Equal(t, metaMock.Keywords, meta.Keywords)
	assert.Equal(t, metaMock.Description, meta.Description)
}
