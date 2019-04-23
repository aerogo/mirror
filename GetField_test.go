package mirror_test

import (
	"reflect"
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

func TestGetField(t *testing.T) {
	movie := &Movie{
		Title:    "The Last Samurai",
		Director: &Person{"Edward Zwick"},
		Actors: []*Person{
			{"Tom Cruise"},
		},
	}

	// Direct descendant
	field, dataType, value, err := mirror.GetField(movie, "Title")
	assert.NoError(t, err)
	assert.Equal(t, "Title", field.Name)
	assert.Equal(t, "string", dataType.Name())
	assert.Equal(t, "The Last Samurai", value.String())

	// Nested descendant
	field, dataType, value, err = mirror.GetField(movie, "Director.Name")
	assert.NoError(t, err)
	assert.Equal(t, "Name", field.Name)
	assert.Equal(t, "string", dataType.Name())
	assert.Equal(t, "Edward Zwick", value.String())

	// Array index
	field, dataType, value, err = mirror.GetField(movie, "Actors[0]")
	assert.NoError(t, err)
	assert.Equal(t, "Actors", field.Name)
	assert.Equal(t, "Person", dataType.Name())
	assert.Equal(t, movie.Actors[0], value.Addr().Interface())

	// Query
	field, dataType, value, err = mirror.GetField(movie, `Actors[Name="Tom Cruise"]`)
	assert.NoError(t, err)
	assert.Equal(t, "Actors", field.Name)
	assert.Equal(t, "Person", dataType.Name())
	assert.Equal(t, movie.Actors[0], value.Addr().Interface())

	// Field of array index
	field, dataType, value, err = mirror.GetField(movie, "Actors[0].Name")
	assert.NoError(t, err)
	assert.Equal(t, "Name", field.Name)
	assert.Equal(t, "string", dataType.Name())
	assert.Equal(t, "Tom Cruise", value.String())

	// Non-existant field
	field, dataType, value, err = mirror.GetField(movie, "Nirvana")
	assert.Error(t, err)
	assert.Nil(t, field)
	assert.Nil(t, dataType)
	assert.Equal(t, reflect.Value{}, value)

	// Non-existant array field
	field, dataType, value, err = mirror.GetField(movie, "Nirvana[0]")
	assert.Error(t, err)
	assert.Nil(t, field)
	assert.Nil(t, dataType)
	assert.Equal(t, reflect.Value{}, value)

	// Non-existant array field with query
	field, dataType, value, err = mirror.GetField(movie, "Nirvana[ID=0]")
	assert.Error(t, err)
	assert.Nil(t, field)
	assert.Nil(t, dataType)
	assert.Equal(t, reflect.Value{}, value)

	// Invalid array index
	field, dataType, value, err = mirror.GetField(movie, "Actors[wtf]")
	assert.Error(t, err)
	assert.Nil(t, field)
	assert.Nil(t, dataType)
	assert.Equal(t, reflect.Value{}, value)

	// Invalid query
	field, dataType, value, err = mirror.GetField(movie, `Actors[Name="Tom]`)
	assert.Error(t, err)
	assert.Nil(t, field)
	assert.Nil(t, dataType)
	assert.Equal(t, reflect.Value{}, value)

	// Query with no result
	field, dataType, value, err = mirror.GetField(movie, `Actors[Name="non-existent"]`)
	assert.Error(t, err)
	assert.Nil(t, field)
	assert.Nil(t, dataType)
	assert.Equal(t, reflect.Value{}, value)

	// Invalid field name
	field, dataType, value, err = mirror.GetField(movie, `Actors[WTF="Tom"]`)
	assert.Error(t, err)
	assert.Nil(t, field)
	assert.Nil(t, dataType)
	assert.Equal(t, reflect.Value{}, value)
}
