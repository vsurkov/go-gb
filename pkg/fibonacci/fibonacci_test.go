package fibonacci

import (
	"fmt"
	"testing"
)

// Test for regular Fib(n int) int function
func TestFibonacci(t *testing.T) {
	answers := populate()

	for i := 0; i <= 20; i++ {
		fmt.Println(i)
		expected := answers[i]
		received := Fib(i)

		if expected != received {
			t.Errorf("Expected %d, but received %d, n=%d", expected, received, i)
		}
	}
}

// Test for improved FibM(n int) int function
func TestFibonacciMaped(t *testing.T) {
	answers := populate()

	for i := 0; i <= 20; i++ {
		expected := answers[i]
		received := FibM(i)

		if expected != received {
			t.Errorf("Expected %d, but received %d, n=%d", expected, received, i)
		}
	}
}

// Prepare answers map for testing
func populate() map[int]int {
	return map[int]int{
		0:  0,
		1:  1,
		2:  1,
		3:  2,
		4:  3,
		5:  5,
		6:  8,
		7:  13,
		8:  21,
		9:  34,
		10: 55,
		11: 89,
		12: 144,
		13: 233,
		14: 377,
		15: 610,
		16: 987,
		17: 1597,
		18: 2584,
		19: 4181,
		20: 6765,
	}
}
