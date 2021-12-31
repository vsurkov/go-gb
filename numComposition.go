package main

import (
	"fmt"
	"strconv"
)

// NumComposition С клавиатуры вводится трехзначное число.
// Выведите цифры, соответствующие количество сотен, десятков и единиц в этом числе.
func NumComposition() error {
	var num int
	fmt.Print("Введите целое число в диапазоне от 100 до 999: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Ошибка: чтения данных")
		return err
	}
	if len(strconv.Itoa(num)) == 3 {
		hundreds := num / 100
		hundredsRemainder := num % 100
		tens := hundredsRemainder / 10
		units := hundredsRemainder % 10

		fmt.Printf("Сотен: %d\n", hundreds)
		fmt.Printf("Десятков: %d\n", tens)
		fmt.Printf("Единиц: %d\n", units)
	} else {
		fmt.Println("Ошибка: введенное значение не трехзначное число")
		return err
	}
	return nil
}
