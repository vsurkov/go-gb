package mySort

import (
	"fmt"
	"os"
)

func main() {

	// Запрашиваем ввод
	str, err := request()
	if err != nil {
		fmt.Println("ОШИБКА: при чтении ввода \n", err)
		os.Exit(1)
	}
	// Подготавливаем данные для работы, очищаем строку, выполняем преобразование типов
	slc, err := prepare(str)
	if err != nil {
		fmt.Println("ОШИБКА: при конвертации значений, одно из значений не целое число \n", err)
		os.Exit(1)
	}

	// Сортируем указанным методом сортировки и выводим пользователю результат
	srd := insertsSort(slc)
	respond(srd)
}
