package main

import (
	"fmt"
	"os"
)

func main() {
	err := router()
	if err != nil {
		fmt.Println("%v", err)
		os.Exit(1)
	}
}
