package main

import (
	fmt "fmt"
	"math"
)

func main() {
	var choise int
	fmt.Println("Введите целое число в диапазоне от 1 до 3, включительно: ")
	fmt.Println("1 - подчет площади прямоугольника")
	fmt.Println("2 - подсчет диаметра и длины окружности по площади круга")
	fmt.Println("3 - состав числа")
	fmt.Print("Ваш выбор: ")
	fmt.Scanln(&choise)

	switch choise {
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
	fmt.Scanln(&a)
	fmt.Print("Введите сторону b прямоугольника: ")
	fmt.Scanln(&b)
	fmt.Printf("Площадь прямоугольника равна: %f\n", a*b)
}

// 2. Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга.
// Площадь круга должна вводиться пользователем с клавиатуры.
func circleData() {
	var s float32
	fmt.Print("Введите площадь круга: ")
	fmt.Scanln(&s)

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
	fmt.Scan(&num)
	hungreds := num / 100
	hungredsRemainder := num % 100
	tens := hungredsRemainder / 10
	units := hungredsRemainder % 10

	fmt.Printf("Сотен: %d\n", hungreds)
	fmt.Printf("Десятков: %d\n", tens)
	fmt.Printf("Единиц: %d\n", units)
}
