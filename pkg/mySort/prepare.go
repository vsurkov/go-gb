package mySort

func prepare(str string) ([]int, error) {
	// Очищаем входные данные, удаляя случайные двойные запятые, далее разбиваем по запятым на слайс
	str = clearData(str)
	slc := toSlice(str)

	// Преобразуем слайс string в слайс int
	return toIntIN(slc)
}
