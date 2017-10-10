package mirror_test

import (
	"testing"

	"github.com/aerogo/mirror"
	"github.com/stretchr/testify/assert"
)

type Movie struct {
	Title    string
	Director *Person
	Actors   []*Person
}

type Person struct {
	Name string
}

func TestGetProperty(t *testing.T) {
	movie := &Movie{
		Title:    "The Last Samurai",
		Director: &Person{"Edward Zwick"},
		Actors: []*Person{
			&Person{"Tom Cruise"},
		},
	}

	// Direct descendant
	dataType, value, err := mirror.GetProperty(movie, "Title")
	assert.NoError(t, err)
	assert.Equal(t, "string", dataType.Name())
	assert.Equal(t, "The Last Samurai", value.String())

	// Direct descendant
	dataType, value, err = mirror.GetProperty(movie, "Director.Name")
	assert.NoError(t, err)
	assert.Equal(t, "string", dataType.Name())
	assert.Equal(t, "Edward Zwick", value.String())

	// Array index
	dataType, value, err = mirror.GetProperty(movie, "Actors[0]")
	assert.NoError(t, err)
	assert.Equal(t, "Person", dataType.Name())
	assert.Equal(t, *movie.Actors[0], value.Interface())

	// Field of array index
	dataType, value, err = mirror.GetProperty(movie, "Actors[0].Name")
	assert.NoError(t, err)
	assert.Equal(t, "string", dataType.Name())
	assert.Equal(t, "Tom Cruise", value.String())

	// Non-existant field
	dataType, value, err = mirror.GetProperty(movie, "Nirvana")
	assert.Error(t, err)
	assert.Nil(t, dataType)
	assert.Nil(t, value)

	// Non-existant array field
	dataType, value, err = mirror.GetProperty(movie, "Nirvana[0]")
	assert.Error(t, err)
	assert.Nil(t, dataType)
	assert.Nil(t, value)

	// Invalid array index
	dataType, value, err = mirror.GetProperty(movie, "Actors[wtf]")
	assert.Error(t, err)
	assert.Nil(t, dataType)
	assert.Nil(t, value)
}
