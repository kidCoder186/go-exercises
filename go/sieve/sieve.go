package sieve

import "math"

// Sieve calculates all prime numbers from 2 up to n.
func Sieve(n int) []int {
	res := []int{}
	visit := map[int]bool{}
	if n >= 2 {
		for i := 2; i <= int(math.Sqrt(float64(n))+1); i++ {
			if !visit[i] {
				for j := i * i; j <= n; j = j + i {
					visit[j] = true
				}
			}
		}

		for i := 2; i <= n; i++ {
			if !visit[i] {
				res = append(res, i)
			}
		}
	}
	return res
}
