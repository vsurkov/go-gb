package fibonacci

import (
	"log"
	"time"
)

// Fib Main implementation of fibonacci function, calculate with recursion
func Fib(n int) uint64 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)
}

// GetFibonacci Improved version of fibonacci function, use map[int]int to store already computed values
var cache = make(map[int]uint64)

func GetFibonacci(n int) uint64 {
	if len(cache) < 3 {
		cache = map[int]uint64{
			0: 0,
			1: 1,
			2: 1,
		}
	}

	if value, ok := cache[n]; ok {
		return value
	}
	cache[n] = GetFibonacci(n-1) + GetFibonacci(n-2)
	return cache[n]
}

// Supply function for calculating runtime of func()
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
