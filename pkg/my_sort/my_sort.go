package mySort

import "fmt"

func mySort() error {
	// Запрашиваем ввод
	raw, err := getUserInput()
	if err != nil {
		fmt.Println("ОШИБКА: при чтении ввода \n", err)
		return err
	}
	// Подготавливаем данные для работы, очищаем строку, выполняем преобразование типов
	slc, err := prapareData(raw)
	if err != nil {
		fmt.Println("ОШИБКА: при конвертации значений, одно из значений не целое число \n", err)
		return err
	}

	// Сортируем указанным методом сортировки и выводим пользователю результат
	sorted := insertsSort(slc)
	doOutput(sorted)
	return nil
}
