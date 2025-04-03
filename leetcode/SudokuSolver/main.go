package SudokuSolver

func getRectIdx(i, j uint16) uint16 {
	return (i/3)*3 + (j / 3)
}

type cell struct {
	row, col uint16
	visited  bool
}

type state struct {
	row, col, rect [10]uint16
}

func (st *state) getSet(i, j uint16) uint16 {
	return st.row[i] | st.col[j] | st.rect[getRectIdx(i, j)]
}

func getFreeBitsCount(v uint16) (res uint16) {
	tmp := uint16(1)
	for range 9 {
		if (tmp & v) == 0 {
			res++
		}
		tmp <<= 1
	}
	return
}

func getCellWithMinVars(emptyCells []cell, astate *state) (uint16, int) {
	minVal := uint16(10)
	idx := len(emptyCells)
	for i := 0; i < len(emptyCells); i++ {
		if emptyCells[i].visited {
			continue
		}
		v := getFreeBitsCount(astate.getSet(emptyCells[i].row, emptyCells[i].col))
		if v < minVal {
			minVal = v
			idx = i
		}
	}
	return minVal, idx
}

func solve(board [][]byte, emptyCells []cell, astate *state) bool {
	if len(emptyCells) == 0 {
		return true
	}
	_, cellIdx := getCellWithMinVars(emptyCells, astate)
	if cellIdx == len(emptyCells) {
		return true
	}
	c := emptyCells[cellIdx]
	rectIdx := getRectIdx(c.row, c.col)
	set := astate.getSet(c.row, c.col)
	if set == 0 {
		return false
	}
	emptyCells[cellIdx].visited = true
	for ch, bit := byte('1'), uint16(1); bit <= 0x100; ch, bit = ch+1, bit<<1 {
		if bit&set != 0 {
			continue
		}
		astate.row[c.row] |= bit
		astate.col[c.col] |= bit
		astate.rect[rectIdx] |= bit
		board[c.row][c.col] = ch
		if solve(board, emptyCells, astate) {
			return true
		}
		astate.row[c.row] &= ^bit
		astate.col[c.col] &= ^bit
		astate.rect[rectIdx] &= ^bit
	}
	emptyCells[cellIdx].visited = false
	return false
}

func solveSudoku(board [][]byte) {
	var astate state
	emptyCells := make([]cell, 0, len(board)*len(board[0]))
	for i, line := range board {
		for j, ch := range line {
			if ch == '.' {
				emptyCells = append(emptyCells, cell{uint16(i), uint16(j), false})
				continue
			}
			val := uint16(1 << (ch - '1'))
			astate.row[i] |= val
			astate.col[j] |= val
			astate.rect[getRectIdx(uint16(i), uint16(j))] |= val
		}
	}
	solve(board, emptyCells, &astate)
}
