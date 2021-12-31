package mySort

import (
	"strconv"
	"strings"
)

func toIntIN(slc []string) ([]int, error) {
	result, err := intify(slc)
	if err != nil {
		return result, err
	}
	return result, err
}

func clearData(str string) string {
	// TODO: обрабатывать другие разделители, большее количество разделителей
	return strings.ReplaceAll(str, ",,", ",")
}

func toSlice(str string) []string {
	return strings.Split(strings.ReplaceAll(str, ",,", ","), ",")
}

func intify(slc []string) ([]int, error) {
	result := []int{}
	for _, element := range slc {
		intVal, err := strconv.Atoi(element)
		if err != nil {
			return nil, err
		}
		result = append(result, intVal)
	}
	return result, nil
}

func stringify(slc []int) string {
	var stringified []string
	for _, i := range slc {
		stringified = append(stringified, strconv.Itoa(i))
	}

	str := strings.Join(stringified, ",")
	return str
}
