package main

import (
	"fmt"
	"math"
)

// CircleData Программа, вычисляющую диаметр и длину окружности по заданной
// площади круга. Площадь круга должна вводиться пользователем с клавиатуры.
func CircleData() error {
	var s float32
	fmt.Print("Введите площадь круга: ")
	_, err := fmt.Scanln(&s)
	if err != nil {
		fmt.Println("Ошибка: чтения данных")
		return err
	}

	d := 2 * math.Sqrt(float64(s/math.Pi))
	fmt.Printf("Диаметр круга: %f\n", d)

	l := math.Pi * d
	fmt.Printf("Длина окружности: %f\n", l)
	return nil
}
