package main

import (
	"fmt"
	"math"
	"os"
)

func Calc() {
	var a, b, res float32
	var op string

	fmt.Print("Введите арифметическую операцию (+, -, *, /, sqrt, pow): ")
	_, err := fmt.Scanln(&op)
	if err != nil {
		exitWrongInput()
	}

	fmt.Print("Введите число A: ")
	_, err = fmt.Scanln(&a)
	if err != nil {
		exitWrongInput()
	}

	if op != "sqrt" {
		fmt.Print("Введите число B: ")
		_, err = fmt.Scanln(&b)
		if err != nil {
			exitWrongInput()
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
			os.Exit(1)
		}
		res = a / b
	case "sqrt":
		res = float32(math.Sqrt(float64(a)))
	case "pow":
		res = float32(math.Pow(float64(a), float64(b)))
	default:
		fmt.Println("Выбрана неизвестная операция")
		os.Exit(1)
	}

	if op != "sqrt" {
		fmt.Printf("Ответ: %f %s %f = %f\n", a, op, b, res)
	} else {
		fmt.Printf("Ответ: √%f = %f\n", a, res)
	}
}
