package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var choice int
	fmt.Println("Введите целое число в диапазоне от 1 до 3, включительно: ")
	fmt.Println("1 - подчет площади прямоугольника")
	fmt.Println("2 - подсчет диаметра и длины окружности по площади круга")
	fmt.Println("3 - состав числа")
	fmt.Print("Ваш выбор: ")

	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Print("Ошибка: вероятно, указанное значение не число")
		return
	}

	switch choice {
	case 1:
		squareArea()
	case 2:
		circleData()
	case 3:
		numComposition()
	default:
		fmt.Println("Ошибка, введенное число не в диапазоне от 1 до 3, или не целое число")
	}
}

// 1. Напишите программу для вычисления площади прямоугольника.
// Длины сторон прямоугольника должны вводиться пользователем с клавиатуры.
func squareArea() {
	var a, b float32
	fmt.Print("Введите сторону a прямоугольника: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Print("Ошибка: вероятно, указанное значение не число")
		return
	}

	fmt.Print("Введите сторону b прямоугольника: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Print("Ошибка: вероятно, указанное значение не число")
		return
	}
	fmt.Printf("Площадь прямоугольника равна: %f\n", a*b)
}

// 2. Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга.
// Площадь круга должна вводиться пользователем с клавиатуры.
func circleData() {
	var s float32
	fmt.Print("Введите площадь круга: ")
	_, err := fmt.Scanln(&s)
	if err != nil {
		fmt.Print("Ошибка: вероятно, указанное значение не число")
		return
	}

	d := 2 * math.Sqrt(float64(s/math.Pi))
	fmt.Printf("Диаметр круга: %f\n", d)

	l := math.Pi * d
	fmt.Printf("Длина окружности: %f\n", l)
}

// 3. С клавиатуры вводится трехзначное число.
// Выведите цифры, соответствующие количество сотен, десятков и единиц в этом числе.
func numComposition() {
	var num int
	fmt.Print("Введите целое число в диапазоне от 100 до 999: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Print("Ошибка: вероятно, указанное значение не число")
		return
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
	}
}
