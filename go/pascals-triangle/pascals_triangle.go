package pascal

// Triangle computes Pascal's triangle up to a given number of rows.
func Triangle(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, i+1)
		res[i][0] = 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j] + res[i-1][j-1]
		}
		res[i][i] = 1
	}
	return res
}
