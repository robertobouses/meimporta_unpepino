package repository

import (
	"strconv"
	"strings"
)

func ConvertStringToFloatSlice(input string) ([]float64, error) {

	valuesStr := strings.Trim(input, "{}")
	valuesStr = strings.ReplaceAll(valuesStr, ",", " ")

	var result []float64
	for _, valueStr := range strings.Fields(valuesStr) {
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}

	return result, nil
}
