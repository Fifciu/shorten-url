package utils

import "fmt"

func SortQueryParamToValues(sortQueryParam string, availableFields []string) (sortBy, direction string, forbiddenField bool) {
	if sortQueryParam == "" {
		return "updated_at", "ASC", false
	}
	for i := 0; i < len(availableFields); i++ {
		if availableFields[i] == sortQueryParam {
			return sortQueryParam, "ASC", false
		} else if fmt.Sprintf("%s:desc", availableFields[i]) == sortQueryParam {
			return sortQueryParam, "DESC", false
		}
	}
	return "", "", true
}
