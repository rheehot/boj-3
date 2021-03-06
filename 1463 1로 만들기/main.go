package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(solve(n))
}

func solve(n int) int {
	m := make([]int, n)
	return f(n, m)
}

func f(n int, m []int) int {
	if n == 1 {
		return 0
	}
	if x := m[n-1]; x != 0 {
		return x
	}

	min := -1
	tryKeep := func(x int) {
		if min == -1 || x < min {
			min = x
		}
	}

	if n%3 == 0 {
		tryKeep(f(n/3, m))
	}
	if n%2 == 0 {
		tryKeep(f(n/2, m))
	}
	tryKeep(f(n-1, m))

	x := 1 + min
	m[n-1] = x
	return x
}

// f7322464 copies rickmccoy's solution.
// There's no memoization but much faster.
func f7322464(n int) int {
	if n < 2 {
		return 0
	}

	a := f7322464(n/2) + n%2
	b := f7322464(n/3) + n%3

	if a < b {
		return 1 + a
	}
	return 1 + b
}
