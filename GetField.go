package mirror

import (
	"errors"
	"reflect"
	"strings"
)

// GetField returns the field referenced by the given path string.
func GetField(root interface{}, path string) (*reflect.StructField, reflect.Type, reflect.Value, error) {
	var field reflect.StructField
	var found bool
	var err error

	t := reflect.TypeOf(root)
	v := reflect.ValueOf(root)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	// Nested properties
	parts := strings.Split(path, ".")

	for _, part := range parts {
		if strings.HasSuffix(part, "]") {
			// Array reference
			arrayStart := strings.Index(part, "[")
			arrayIndexString := part[arrayStart+1 : len(part)-1]
			part = part[:arrayStart]

			// Get array
			field, found = t.FieldByName(part)

			if !found {
				return nil, nil, reflect.Value{}, errors.New("Field '" + part + "' does not exist in type " + t.Name())
			}

			array := reflect.Indirect(v.FieldByName(field.Name))

			// Get slice element
			v, _, err = GetSliceElement(array.Interface(), arrayIndexString)

			if err != nil {
				return nil, nil, reflect.Value{}, err
			}

			t = v.Type()
		} else {
			// Non-array reference
			field, found = t.FieldByName(part)

			if !found {
				return nil, nil, reflect.Value{}, errors.New("Field '" + part + "' does not exist in type " + t.Name())
			}

			t = field.Type
			v = reflect.Indirect(v.FieldByName(field.Name))
		}

		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
	}

	return &field, t, v, nil
}
