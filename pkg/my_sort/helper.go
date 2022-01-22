package mySort

import (
	"strconv"
)

func intifySlice(elements []string) ([]int, error) {
	result := []int{}
	for _, element := range elements {
		intVal, err := strconv.Atoi(element)
		if err != nil {
			return nil, err
		}
		result = append(result, intVal)
	}
	return result, nil
}
