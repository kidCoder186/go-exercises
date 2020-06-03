package matrix

// Pair structure
type Pair [2]int

// Saddle returns all saddle points of the matrix.
func (m *Matrix) Saddle() []Pair {
	res := []Pair{}
	maxRows := make([]int, m.w)
	minCols := make([]int, m.h)
	for i := 0; i < m.w; i++ {
		maxRows[i] = m.data[i][0]
		for j := 1; j < m.h; j++ {
			if maxRows[i] < m.data[i][j] {
				maxRows[i] = m.data[i][j]
			}
		}
	}
	for j := 0; j < m.h; j++ {
		minCols[j] = m.data[0][j]
		for i := 1; i < m.w; i++ {
			if minCols[j] > m.data[i][j] {
				minCols[j] = m.data[i][j]
			}
		}
	}
	for i := 0; i < m.w; i++ {
		for j := 0; j < m.h; j++ {
			if m.data[i][j] == maxRows[i] && m.data[i][j] == minCols[j] {
				res = append(res, Pair{i, j})
			}
		}
	}
	return res
}
