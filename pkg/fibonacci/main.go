package fibonacci

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// int num, use for calculating functions F(num)
	num := 40
	runFibonacci(num)
	runFibonacciMapped(num)
}

// Run regular Fib(n int) int function with time tracking
func runFibonacci(num int) {
	if num < 0 {
		log.Panicf("Fibonacci can be counted only for positive numbers, but received %d", num)
	}
	defer timeTrack(time.Now(), "regular fibonacci")
	fmt.Printf("Fib(%d) = %d\n", num, Fib(int(num)))
}

// Run improved FibM(n int) int function with time tracking
func runFibonacciMapped(num int) {
	if num < 0 {
		log.Panicf("Fibonacci can be counted only for positive numbers, but received %d", num)
	}
	defer timeTrack(time.Now(), "mapped fibonacci")
	fmt.Printf("FibM(%d) = %d\n", num, GetFibonacci(int(num)))
}
