package main

import (
	"fmt"
)

// SquareArea Программа для вычисления площади прямоугольника. Длины сторон
// прямоугольника должны вводиться пользователем с клавиатуры.
func SquareArea() {
	var a, b float32
	fmt.Print("Введите сторону a прямоугольника: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		exitWrongInput()
	}

	fmt.Print("Введите сторону b прямоугольника: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		exitWrongInput()
	}
	fmt.Printf("Площадь прямоугольника равна: %f\n", a*b)
}
