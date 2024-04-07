package helpers

import "strconv"

func ValidatePaginationParams(offset, limit string) (int, int, []string) {

	errors := []string{}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		errors = append(errors, err.Error())
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		errors = append(errors, err.Error())
	}

	if len(errors) > 0 {
		return 0, 0, errors
	}

	if offsetInt < 0 {
		offsetInt = 0
	}
	if limitInt%5 != 0 {
		limitInt = 10
	}
	return offsetInt, limitInt, nil
}
