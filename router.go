package main

import (
	"fmt"
)

func router() error {
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
		fmt.Println("Ошибка: чтения данных")
		return err
	}

	switch choice {
	case 1:
		err := SquareArea()
		if err != nil {
			return err
		}
	case 2:
		err := CircleData()
		if err != nil {
			return err
		}
	case 3:
		err := NumComposition()
		if err != nil {
			return err
		}
	case 4:
		err := Calc()
		if err != nil {
			return err
		}
	default:
		fmt.Println("Ошибка, введенное число не в диапазоне от 1 до 3, или не целое число")
		return err
	}
	return nil
}
