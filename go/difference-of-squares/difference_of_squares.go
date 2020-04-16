package diffsquares

// SquareOfSum returns square of sum of the first 'n' natural number
func SquareOfSum(n int) int {
	res := n * (n + 1) / 2
	return res * res
}

// SumOfSquares returns sum of squares of the first 'n' natural number
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference returns difference between SquareOfSum and SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
