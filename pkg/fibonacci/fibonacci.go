package fibonacci

import (
	"log"
	"time"
)

// Fib Main implementation of fibonacci function, calculate with recursion
func Fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)
}

// FibM Improved version of fibonacci function, use map[int]int to store already computed values
func FibM(n int) int {
	var cache = map[int]int{
		0: 0,
		1: 1,
		2: 1,
	}
	return worker(cache, n)
}

// Supply function for FibM(n int) int implementation
func worker(cache map[int]int, n int) int {
	if value, ok := cache[n]; ok {
		return value
	}
	cache[n] = worker(cache, n-1) + worker(cache, n-2)
	return cache[n]
}

// Supply function for calculating runtime of func()
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
