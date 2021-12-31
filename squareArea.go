package main

import (
	"fmt"
)

// SquareArea Программа для вычисления площади прямоугольника. Длины сторон
// прямоугольника должны вводиться пользователем с клавиатуры.
func SquareArea() error {
	var a, b float32
	fmt.Print("Введите сторону a прямоугольника: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Ошибка: чтения данных")
		return err
	}

	fmt.Print("Введите сторону b прямоугольника: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println("Ошибка: чтения данных")
		return err
	}
	fmt.Printf("Площадь прямоугольника равна: %f\n", a*b)
	return nil
}
