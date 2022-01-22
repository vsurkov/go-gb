package mySort

import (
	"os"
)

func main() {
	err := mySort()
	if err != nil {
		os.Exit(1)
	}
}
