package mySort

import "fmt"

func request() (string, error) {
	// TODO добавить альтернативные методы ввода, ex: чтение из параметров запуска
	var str string
	fmt.Print("Введите через запятую последовательность целых чисел для сортировки: ")

	_, err := fmt.Scan(&str)
	if err != nil {
		return "", err
	}
	return str, nil
}
