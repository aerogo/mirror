package mirror

import (
	"errors"
	"reflect"
)

// GetChildField returns the direct descendant field referenced by the given string.
func GetChildField(root interface{}, name string) (*reflect.StructField, reflect.Type, reflect.Value, error) {
	t := reflect.TypeOf(root)
	v := reflect.ValueOf(root)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	field, found := t.FieldByName(name)

	if !found {
		return nil, nil, reflect.Value{}, errors.New("Field '" + name + "' does not exist in type " + t.Name())
	}

	t = field.Type
	v = v.FieldByName(field.Name)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	return &field, t, v, nil
}
