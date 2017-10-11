package mirror

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// GetProperty returns the property referenced by the given path string.
func GetProperty(root interface{}, path string) (*reflect.StructField, reflect.Type, reflect.Value, error) {
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
			part = part[:arrayStart]

			if strings.Contains(arrayIndexString, "=") {
				keyValue := strings.Split(arrayIndexString, "=")
				queryKey := keyValue[0]
				queryValue := keyValue[1]

				var queryValueData interface{}
				err := json.Unmarshal([]byte(queryValue), &queryValueData)

				if err != nil {
					return nil, nil, reflect.Value{}, err
				}

				// Get array
				field, found = t.FieldByName(part)

				if !found {
					return nil, nil, reflect.Value{}, errors.New("Field '" + part + "' does not exist in type " + t.Name())
				}

				v = reflect.Indirect(v.FieldByName(field.Name))

				// Find array index with correct query value
				elementFound := false

				for i := 0; i < v.Len(); i++ {
					nextV := reflect.Indirect(v.Index(i))

					// elementField, found := t.FieldByName(queryKey)

					// if !found {
					// 	return nil, nil, reflect.Value{}, errors.New("Field '" + part + "' does not exist in type " + t.Name())
					// }

					elementField := reflect.Indirect(nextV.FieldByName(queryKey))

					if elementField.String() == fmt.Sprint(queryValueData) {
						elementFound = true
						v = nextV
						t = v.Type()
						break
					}
				}

				if !elementFound {
					return nil, nil, reflect.Value{}, fmt.Errorf("Could not find array item where %s = %s", queryKey, queryValue)
				}
			} else {
				arrayIndex, err := strconv.Atoi(arrayIndexString)

				if err != nil {
					return nil, nil, reflect.Value{}, err
				}

				// Get the slice first
				field, found = t.FieldByName(part)

				if !found {
					return nil, nil, reflect.Value{}, errors.New("Field '" + part + "' does not exist in type " + t.Name())
				}

				v = reflect.Indirect(v.FieldByName(field.Name))

				// Now get the object referenced at the given index
				v = reflect.Indirect(v.Index(arrayIndex))
				t = v.Type()
			}
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
