package common

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"unicode"
)

func BytesToMap(data []byte) (map[string]any, error) {
	var result map[string]any
	err := json.Unmarshal(data, &result)
	if err != nil {
		log.Println("Error decoding bytes:", err)
		return nil, err
	}
	return result, nil
}

func ToSnakeCase(s string) string {
	var result strings.Builder
	var lastWasUpper bool

	for _, r := range s {
		if unicode.IsUpper(r) {
			if lastWasUpper {
				result.WriteRune(unicode.ToLower(r))
			} else {
				if result.Len() > 0 {
					result.WriteRune('_')
				}
				result.WriteRune(unicode.ToLower(r))
			}
			lastWasUpper = true
		} else {
			result.WriteRune(r)
			lastWasUpper = false
		}
	}
	return result.String()
}
func CompareAttributes(dataMap map[string]any, dataStruct interface{}) []string {
	var differentAttributes []string

	structType := reflect.TypeOf(dataStruct)

	for key := range dataMap {
		snakeCaseKey := ToSnakeCase(key)
		if _, ok := structType.FieldByNameFunc(func(fieldName string) bool {
			return ToSnakeCase(fieldName) == snakeCaseKey
		}); !ok {
			differentAttributes = append(differentAttributes, fmt.Sprintf("Unexpected '%v' field", key))
		}
	}

	return differentAttributes
}
