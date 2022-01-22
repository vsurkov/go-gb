package mySort

import (
	"strings"
)

func prapareData(raw string) ([]int, error) {
	// Очищаем входные данные, удаляя случайные двойные запятые, далее разбиваем по запятым на слайс
	raw = strings.ReplaceAll(raw, ",,", ",")
	slc := strings.Split(raw, ",")

	// Преобразуем слайс string в слайс int
	result, err := intifySlice(slc)
	if err != nil {
		return result, err
	}
	return result, nil
}
