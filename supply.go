package main

import (
	"fmt"
	"os"
)

func exitWrongInput() {
	fmt.Println("Ошибка: вероятно, указанное значение не число ")
	os.Exit(1)
}
