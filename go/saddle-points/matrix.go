package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix structure
type Matrix struct {
	w     int
	h     int
	data  [][]int
	dataT [][]int
}

// New creates a new martix object
func New(input string) (*Matrix, error) {
	lines := strings.Split(input, "\n")
	w := len(lines)
	matrix := make([][]int, w)

	h := len(strings.Fields(lines[0]))
	matrixT := make([][]int, h)

	for j := 0; j < h; j++ {
		matrixT[j] = make([]int, w)
	}
	for i, line := range lines {
		row := strings.Fields(line)
		if len(row) != h {
			return nil, errors.New("row has not same length")
		}
		matrix[i] = make([]int, h)
		for j := 0; j < h; j++ {
			val, err := strconv.Atoi(row[j])
			if err != nil {
				return nil, err
			}
			matrix[i][j] = val
			matrixT[j][i] = val
		}
	}
	return &Matrix{w, h, matrix, matrixT}, nil
}

// Rows returns all rows of the matrix
func (m *Matrix) Rows() [][]int {
	rows := make([][]int, m.w)
	for i := 0; i < m.w; i++ {
		rows[i] = make([]int, m.h)
		for j := 0; j < m.h; j++ {
			rows[i][j] = m.data[i][j]
		}
	}
	return rows
}

// Cols returns all cols of the matrix
func (m *Matrix) Cols() [][]int {
	cols := make([][]int, m.h)
	for i := 0; i < m.h; i++ {
		cols[i] = make([]int, m.w)
		for j := 0; j < m.w; j++ {
			cols[i][j] = m.dataT[i][j]
		}
	}
	return cols
}

// Set sets matrix[x][y] = val
func (m *Matrix) Set(x, y, val int) bool {
	if x < 0 || x >= m.w || y < 0 || y >= m.h {
		return false
	}
	m.data[x][y] = val
	m.dataT[y][x] = val
	return true
}
