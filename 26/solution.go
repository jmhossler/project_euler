package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Reciprocal Cycles")

	var primes []int
	longest_cycle := 0
	num_longest_cycle := 0
	for i := 999; i > 1 && longest_cycle < i; i-- {
		if is_prime(i) {
			primes = append(primes, i)
			x := get_cycle_len(i)
			fmt.Printf("%d: %d\n", i, x)
			if x >= longest_cycle {
				longest_cycle = x
				num_longest_cycle = i
			}
		}
	}

	fmt.Println("Primes: ", primes)
	fmt.Printf("Longest cycle: %d for d=%d\n", longest_cycle, num_longest_cycle)
}

func get_cycle_len(val int) int {
	for i := 0; i < val; i++ {
		if math.Mod(math.Pow10(i), float64(val)) == 1 {
			return i
		}
	}

	return 0
}

func is_prime(val int) bool {
	for i := 2; i < int(math.Sqrt(float64(val))); i++ {
		if val%i == 0 {
			return false
		}
	}
	return true
}
