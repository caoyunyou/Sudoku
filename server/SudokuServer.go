package server

import (
	"com.cyy/sudoku/types"
	"math/rand"
)

// 生成完整的数独终盘
func generateCompleteSudoku() _types.Game {
	var sudoku _types.Game
	backtrack(&sudoku, 0, 0)
	return sudoku
}

// 回溯算法填充数独
func backtrack(sudoku *_types.Game, row, col int) bool {
	if row == 9 {
		return true
	}

	nextRow, nextCol := row, col+1
	if nextCol == 9 {
		nextRow = row + 1
		nextCol = 0
	}

	if sudoku[row][col].Num != 0 {
		return backtrack(sudoku, nextRow, nextCol)
	}

	nums := rand.Perm(9)
	for _, n := range nums {
		num := n + 1
		if IsValid(sudoku, row, col, num) {
			sudoku[row][col].Num = num
			sudoku[row][col].IsHole = false
			if backtrack(sudoku, nextRow, nextCol) {
				return true
			}
			sudoku[row][col].Num = 0
			sudoku[row][col].IsHole = true
		}
	}
	return false
}

// IsValid 检查填入数字是否合法
func IsValid(sudoku *_types.Game, row, col, num int) bool {

	// 检查行
	for c := 0; c < 9; c++ {
		if sudoku[row][c].Num == num {
			return false
		}
	}

	// 检查列
	for r := 0; r < 9; r++ {
		if sudoku[r][col].Num == num {
			//log.Printf("列内存在相等元素 row:%d ,col:%d ,num:%d", r, col, num)
			return false
		}
	}

	// 检查宫格
	startRow, startCol := row/3*3, col/3*3
	for r := startRow; r < startRow+3; r++ {
		for c := startCol; c < startCol+3; c++ {
			if sudoku[r][c].Num == num {
				return false
			}
		}
	}
	return true
}

// GenerateSudokuPuzzle 生成数独题目（保留指定数量的已知数字）
func GenerateSudokuPuzzle(clues int) _types.Game {
	complete := generateCompleteSudoku()
	puzzle := complete

	cells := rand.Perm(81)
	digitsToRemove := 81 - clues

	for _, idx := range cells {
		if digitsToRemove <= 0 {
			break
		}
		row := idx / 9
		col := idx % 9
		original := puzzle[row][col]

		if original.Num == 0 {
			continue
		}
		// 尝试挖洞
		puzzle[row][col].Num = 0
		if countSolutions(&puzzle) == 1 {
			digitsToRemove--
			puzzle[row][col].IsHole = true
		} else {
			// 恢复数字
			puzzle[row][col] = original
		}
	}

	return puzzle
}

// 计算解的个数（最多计算到2个）
func countSolutions(sudoku *_types.Game) int {
	var copySudoku _types.Game
	copySudoku = *sudoku
	count := 0

	var backtrackSolutions func(int, int)
	backtrackSolutions = func(row, col int) {
		if count >= 2 {
			return
		}

		if row == 9 {
			count++
			return
		}

		nextRow, nextCol := row, col+1
		if nextCol == 9 {
			nextRow = row + 1
			nextCol = 0
		}

		if copySudoku[row][col].Num != 0 {
			backtrackSolutions(nextRow, nextCol)
			return
		}

		for num := 1; num <= 9; num++ {
			if IsValid(&copySudoku, row, col, num) {
				copySudoku[row][col].Num = num
				backtrackSolutions(nextRow, nextCol)
				copySudoku[row][col].Num = 0

				if count >= 2 {
					return
				}
			}
		}
	}

	backtrackSolutions(0, 0)
	return count
}
