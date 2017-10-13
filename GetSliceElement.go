package mirror

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// GetSliceElement returns the element with the given index or query.
func GetSliceElement(arrayObj interface{}, arrayIndexString string) (value reflect.Value, index int, err error) {
	array := reflect.ValueOf(arrayObj)

	// Query
	if strings.Contains(arrayIndexString, "=") {
		keyValue := strings.Split(arrayIndexString, "=")
		queryKey := keyValue[0]
		queryValue := keyValue[1]

		var queryValueData interface{}
		err := json.Unmarshal([]byte(queryValue), &queryValueData)

		if err != nil {
			return reflect.Value{}, -1, err
		}

		// Find array index with correct query value
		for i := 0; i < array.Len(); i++ {
			nextV := reflect.Indirect(array.Index(i))
			_, found := nextV.Type().FieldByName(queryKey)

			if !found {
				return reflect.Value{}, -1, errors.New("Field '" + queryKey + "' does not exist in type " + nextV.Type().Name())
			}

			elementField := reflect.Indirect(nextV.FieldByName(queryKey))

			if elementField.String() == fmt.Sprint(queryValueData) {
				return nextV, i, nil
			}
		}

		return reflect.Value{}, -1, fmt.Errorf("Could not find array item where %s = %s", queryKey, queryValue)
	}

	// Normal integer index
	arrayIndex, err := strconv.Atoi(arrayIndexString)

	if err != nil {
		return reflect.Value{}, -1, err
	}

	// Now get the object referenced at the given index
	return reflect.Indirect(array.Index(arrayIndex)), arrayIndex, nil
}
