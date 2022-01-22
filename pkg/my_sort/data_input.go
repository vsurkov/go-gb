package mySort

import "fmt"

func getUserInput() (string, error) {
	// TODO добавить альтернативные методы ввода, ex: чтение из параметров запуска
	var raw string
	fmt.Print("Введите через запятую последовательность целых чисел для сортировки: ")

	_, err := fmt.Scan(&raw)
	if err != nil {
		return "", err
	}
	return raw, nil
}
