package fibonacci

import (
	"testing"
)

// Test for regular Fib(n int) int function
func TestFibonacci(t *testing.T) {
	answers := populate()

	for i := 0; i <= 20; i++ {
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
		received := GetFibonacci(i)

		if expected != received {
			t.Errorf("Expected %d, but received %d, n=%d", expected, received, i)
		}
	}
}

// Prepare answers map for testing
func populate() map[int]uint64 {
	fibonacciKnown := []uint64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765}
	m := make(map[int]uint64)
	for i, v := range fibonacciKnown {
		m[i] = v
	}
	return m
}
