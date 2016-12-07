package main

import (
	"flag"
	"fmt"
	"time"
)

var max = flag.Int("max", 4000000, "Maximum fib number to calculate")

func main() {
	flag.Parse()
	fmt.Println("Project Euler : Problem 2")

	// Memoized fibs so that each fib is only calculated once
	fibs := []int{1, 2}

	start := time.Now()
	sum := 2
	limit := *max
	for i := 2; fibs[len(fibs)-1] < limit; i++ {
		next_fib := fibs[i-1] + fibs[i-2]
		fibs = append(fibs, next_fib)

		if next_fib%2 == 0 {
			sum += next_fib
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("M1 Sum of even fibonacci numbers below %d: %d --- %s\n", limit, sum, elapsed)

	start = time.Now()
	//Comparison of algorithms
	sum_2 := 0
	for i := 1; i < limit; i++ {
		next_fib := fib(i)
		if next_fib%2 == 0 {
			sum_2 += next_fib
		}
	}

	elapsed = time.Since(start)
	fmt.Printf("M2 Sum of even fibonacci numbers below %d: %d --- %s\n", limit, sum, elapsed)
}

func fib(n int) int {
	if n < 2 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}
