package main

import (
	"fmt"
	"math"
)

func Calc() error {
	var a, b, res float32
	var op string

	fmt.Print("Введите арифметическую операцию (+, -, *, /, sqrt, pow): ")
	_, err := fmt.Scanln(&op)
	if err != nil {
		fmt.Println("Ошибка: чтения данных")
		return err
	}

	fmt.Print("Введите число A: ")
	_, err = fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Ошибка: чтения данных")
		return err
	}

	if op != "sqrt" {
		fmt.Print("Введите число B: ")
		_, err = fmt.Scanln(&b)
		if err != nil {
			fmt.Println("Ошибка: чтения данных")
			return err
		}
	}

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: в элементарной алгебре деление на 0 запрещено")
			return err
		}
		res = a / b
	case "sqrt":
		res = float32(math.Sqrt(float64(a)))
	case "pow":
		res = float32(math.Pow(float64(a), float64(b)))
	default:
		fmt.Printf("Выбрана неизвестная операция - '%s\n'", op)
		return err
	}

	if op != "sqrt" {
		fmt.Printf("Ответ: %f %s %f = %f\n", a, op, b, res)
	} else {
		fmt.Printf("Ответ: √%f = %f\n", a, res)
	}
	return nil
}
