package minesweeper

import "errors"

func (b Board) isValid() bool {
	if len(b) == 0 {
		return true
	}
	m := len(b)
	n := len(b[0])
	// check rectangle shape
	for i := 0; i < m; i++ {
		if len(b[i]) != n {
			return false
		}
	}
	// check 4 corners
	if b[0][0] != '+' || b[0][n-1] != '+' || b[m-1][0] != '+' || b[m-1][n-1] != '+' {
		return false
	}
	// check 2 horizontal borders
	for j := 1; j < n-1; j++ {
		if b[0][j] != '-' || b[m-1][j] != '-' {
			return false
		}
	}
	// check 2 vertical borders
	for i := 1; i < m-1; i++ {
		if b[i][0] != '|' || b[i][n-1] != '|' {
			return false
		}
	}
	// check all elements in board
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if !(b[i][j] == '*' || b[i][j] == ' ' || (b[i][j] >= '1' && b[i][j] <= '9')) {
				return false
			}
		}
	}
	return true
}

// Count adds the numbers to a minesweeper board
func (b Board) Count() error {
	if !b.isValid() {
		return errors.New("invalid board")
	}
	m := len(b)
	if m > 0 {
		n := len(b[0])
		dx := []int{0, 1, 1, 1, 0, -1, -1, -1}
		dy := []int{1, 1, 0, -1, -1, -1, 0, 1}
		for i := 1; i < m-1; i++ {
			for j := 1; j < n-1; j++ {
				if b[i][j] == ' ' {
					var mines byte = '0'
					for k := 0; k < 8; k++ {
						if b[i+dx[k]][j+dy[k]] == '*' {
							mines++
						}
					}
					if mines != '0' {
						b[i][j] = mines
					}
				}
			}
		}
	}
	return nil
}
