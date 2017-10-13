package mirror_test

import (
	"testing"

	"github.com/aerogo/mirror"
	"github.com/stretchr/testify/assert"
)

func TestGetSliceElement(t *testing.T) {
	movies := []*Movie{
		&Movie{
			Title: "The Last Samurai",
		},
		&Movie{
			Title: "Harry Potter",
		},
	}

	// Find by index
	value, index, err := mirror.GetSliceElement(movies, `1`)
	assert.NoError(t, err)
	assert.Equal(t, 1, index)
	assert.Equal(t, movies[1], value.Addr().Interface())

	// Find by query
	value, index, err = mirror.GetSliceElement(movies, `Title="Harry Potter"`)
	assert.NoError(t, err)
	assert.Equal(t, 1, index)
	assert.Equal(t, movies[1], value.Addr().Interface())
}
