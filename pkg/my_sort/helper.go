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

//
//func stringify(slc []int) string {
//	var stringified []string
//	for _, i := range slc {
//		stringified = append(stringified, strconv.Itoa(i))
//	}
//
//	str := strings.Join(stringified, ",")
//	return str
//}
