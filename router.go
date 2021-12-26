package main

import (
	"fmt"
	"os"
)

func router() {
	var choice int
	msg := `Введите целое число в диапазоне от 1 до 3, включительно:
	1 - подчет площади прямоугольника
	2 - подсчет диаметра и длины окружности по площади круга
	3 - состав числа
	4 - калькулятор
Ваш выбор: `
	fmt.Print(msg)

	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Ошибка: вероятно, указанное значение не число ")
		os.Exit(1)
	}

	switch choice {
	case 1:
		SquareArea()
	case 2:
		CircleData()
	case 3:
		NumComposition()
	case 4:
		Calc()
	default:
		fmt.Println("Ошибка, введенное число не в диапазоне от 1 до 3, или не целое число")
	}
}
