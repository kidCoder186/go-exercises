package transpose

// Transpose returns a transpose matrix of a given matrix.
func Transpose(matrix []string) []string {
	var n int
	m := len(matrix)
	if m == 0 {
		return matrix
	}
	for i := m - 1; i >= 0; i-- {
		if len(matrix[i]) > n {
			n = len(matrix[i])
		} else {
			for len(matrix[i]) != n {
				matrix[i] += " "
			}
		}
	}
	var res []string
	for j := 0; j < n; j++ {
		var output string
		for i := 0; i < m; i++ {
			if j >= len(matrix[i]) {
				break
			}
			output += string(matrix[i][j])
		}
		res = append(res, output)
	}
	return res
}
