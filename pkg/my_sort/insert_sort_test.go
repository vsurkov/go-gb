package mySort

import "fmt"

func ExampleReverse() {
	s := []int{9, 0, 2, 4, 5, 3, 1, 7, 8}
	fmt.Println(insertsSort(s))
	// Output: [0 1 2 3 4 5 7 8 9]
}
