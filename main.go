package main

import "fmt"

// MaxProfit returns the best profit achievable from one buy followed by one
// sell, given the price on each day. Returns 0 if no profitable trade exists.
func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minSoFar, best := prices[0], 0
	for _, p := range prices[1:] {
		if p-minSoFar > best {
			best = p - minSoFar
		}
		if p < minSoFar {
			minSoFar = p
		}
	}
	return best
}

func main() {
	fmt.Println(MaxProfit([]int{7, 1, 5, 3, 6, 4}))
}
