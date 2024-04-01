package common

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
)

func BytesToMap(data []byte) (map[string]string, error) {
	var result map[string]string
	err := json.Unmarshal(data, &result)
	if err != nil {
		log.Fatal("Error decoding bytes:", err)
		return nil, err
	}
	return result, nil
}

func CompareAttributes(dataMap map[string]string, dataStruct interface{}) []string {
	var differentAttributes []string

	structType := reflect.TypeOf(dataStruct)

	for key := range dataMap {
		lowercaseKey := strings.ToLower(key)

		if _, ok := structType.FieldByNameFunc(func(fieldName string) bool {
			return strings.ToLower(fieldName) == lowercaseKey
		}); !ok {
			differentAttributes = append(differentAttributes, fmt.Sprintf("Unexpected '%v' field", key))
		}
	}

	return differentAttributes
}
