package mirror

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

// GetProperty returns the property referenced by the given path string.
func GetProperty(root interface{}, path string) (reflect.Type, *reflect.Value, error) {
	var field reflect.StructField
	var found bool

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
			arrayIndex, err := strconv.Atoi(arrayIndexString)
			part = part[:arrayStart]

			if err != nil {
				return nil, nil, err
			}

			// Get the slice first
			field, found = t.FieldByName(part)

			if !found {
				return nil, nil, errors.New("Field '" + part + "' does not exist in type " + t.Name())
			}

			v = reflect.Indirect(v.FieldByName(field.Name))

			// Now get the object referenced at the given index
			v = reflect.Indirect(v.Index(arrayIndex))
			t = v.Type()
		} else {
			// Non-array reference
			field, found = t.FieldByName(part)

			if !found {
				return nil, nil, errors.New("Field '" + part + "' does not exist in type " + t.Name())
			}

			t = field.Type
			v = reflect.Indirect(v.FieldByName(field.Name))
		}

		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
	}

	return t, &v, nil
}
