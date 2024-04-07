package helpers

import (
	"encoding/json"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func ValidateRequest[T any](body []byte, modelReq *T) (*T, []string) {

	dataMap, err := BytesToMap(body)
	if err != nil {
		return nil, strings.Split(err.Error(), "\n")
	}

	if err := json.Unmarshal(body, modelReq); err != nil {
		return nil, strings.Split(err.Error(), "\n")
	}
	if err := validate.Struct(modelReq); err != nil {
		return nil, strings.Split(err.Error(), "\n")
	}

	differentAttributes := CompareAttributes(dataMap, *modelReq)
	if len(differentAttributes) > 0 {
		return nil, differentAttributes
	}

	return modelReq, nil
}
