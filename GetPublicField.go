package mirror

import (
	"errors"
	"reflect"
	"strings"
)

// GetPublicField returns the field referenced by the given path string
// if none of the subfields have the "private" tag set to "true".
func GetPublicField(root interface{}, path string) (*reflect.StructField, reflect.Type, reflect.Value, error) {
	var field reflect.StructField
	var found bool

	t := reflect.TypeOf(root)
	v := reflect.ValueOf(root)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	for _, part := range strings.Split(path, ".") {
		field, found = t.FieldByName(part)

		if !found {
			return nil, nil, reflect.Value{}, errors.New("Field '" + part + "' does not exist in type " + t.Name())
		}

		if field.Tag.Get("private") == "true" {
			return nil, nil, reflect.Value{}, errors.New("Field '" + part + "' in type " + t.Name() + " is private")
		}

		t = field.Type
		v = v.FieldByName(field.Name)

		if t.Kind() == reflect.Ptr {
			t = t.Elem()
			v = v.Elem()
		}
	}

	return &field, t, v, nil
}
