package server

import (
	"fmt"
	"log"
	"math/rand"
)

// Sudoku 表示一个9x9的数独棋盘
type Sudoku [9][9]int

// 生成完整的数独终盘
func generateCompleteSudoku() Sudoku {
	var sudoku Sudoku
	backtrack(&sudoku, 0, 0)
	return sudoku
}

// 回溯算法填充数独
func backtrack(sudoku *Sudoku, row, col int) bool {
	if row == 9 {
		return true
	}

	nextRow, nextCol := row, col+1
	if nextCol == 9 {
		nextRow = row + 1
		nextCol = 0
	}

	if sudoku[row][col] != 0 {
		return backtrack(sudoku, nextRow, nextCol)
	}

	nums := rand.Perm(9)
	for _, n := range nums {
		num := n + 1
		if IsValid(sudoku, row, col, num) {
			sudoku[row][col] = num
			if backtrack(sudoku, nextRow, nextCol) {
				return true
			}
			sudoku[row][col] = 0
		}
	}
	return false
}

// 检查数字是否合法
func IsValid(sudoku *Sudoku, row, col, num int) bool {

	//log.Printf("该位置是否可填 i:%d ,j:%d ,num:%d", row, col, num)
	// 检查行
	for c := 0; c < 9; c++ {
		if sudoku[row][c] == num {
			//log.Printf("行内存在相等元素 row:%d ,col:%d ,num:%d", row, c, num)
			return false
		}
	}

	// 检查列
	for r := 0; r < 9; r++ {
		if sudoku[r][col] == num {
			//log.Printf("列内存在相等元素 row:%d ,col:%d ,num:%d", r, col, num)
			return false
		}
	}

	// 检查宫格
	startRow, startCol := row/3*3, col/3*3
	for r := startRow; r < startRow+3; r++ {
		for c := startCol; c < startCol+3; c++ {
			if sudoku[r][c] == num {
				//log.Printf("宫格内存在相等元素 row:%d ,col:%d ,num:%d", r, c, num)
				return false
			}
		}
	}
	return true
}

func CanPlaceNumber(board [9][9]int, row, col, num int) bool {
	// 检查数独是否为9x9
	//if len(board) != 9 {
	//	return false
	//}
	//for _, r := range board {
	//	if len(r) != 9 {
	//		return false
	//	}
	//}

	// 检查行和列的范围
	if row < 0 || row >= 9 || col < 0 || col >= 9 {
		log.Printf("行列范围错误 i:%d ,j:%d ,num:%d", row, col, num)
		return false
	}

	// 检查数字是否合法
	if num < 1 || num > 9 {
		log.Printf("数字格式错误 i:%d ,j:%d ,num:%d", row, col, num)
		return false
	}

	// 检查当前格子是否为空
	if board[row][col] != 0 {
		log.Printf("该窗格已存在值 i:%d ,j:%d ,num:%d", row, col, num)
		return false
	}

	// 检查同一行是否存在重复
	for c := 0; c < 9; c++ {
		if c != col && board[row][c] == num {
			log.Printf("行内重复 i:%d ,j:%d ,num:%d", row, c, num)
			return false
		}
	}

	// 检查同一列是否存在重复
	for r := 0; r < 9; r++ {
		if r != row && board[r][col] == num {
			log.Printf("列内重复 i:%d ,j:%d ,num:%d", r, col, num)
			return false
		}
	}

	// 检查3x3宫格是否存在重复
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for r := startRow; r < startRow+3; r++ {
		for c := startCol; c < startCol+3; c++ {
			if r == row && c == col { // 跳过当前格子
				continue
			}
			if board[r][c] == num {
				log.Printf("窗格内重复 i:%d ,j:%d ,num:%d", r, c, num)
				return false
			}
		}
	}

	return true
}

// GenerateSudokuPuzzle 生成数独题目（保留指定数量的已知数字）
func GenerateSudokuPuzzle(clues int) Sudoku {
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

		if original == 0 {
			continue
		}

		// 尝试挖洞
		puzzle[row][col] = 0
		if countSolutions(&puzzle) == 1 {
			digitsToRemove--
		} else {
			// 恢复数字
			puzzle[row][col] = original
		}
	}

	return puzzle
}

// 计算解的个数（最多计算到2个）
func countSolutions(sudoku *Sudoku) int {
	var copySudoku Sudoku
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

		if copySudoku[row][col] != 0 {
			backtrackSolutions(nextRow, nextCol)
			return
		}

		for num := 1; num <= 9; num++ {
			if IsValid(&copySudoku, row, col, num) {
				copySudoku[row][col] = num
				backtrackSolutions(nextRow, nextCol)
				copySudoku[row][col] = 0

				if count >= 2 {
					return
				}
			}
		}
	}

	backtrackSolutions(0, 0)
	return count
}

// 打印数独
func printSudoku(sudoku Sudoku) {
	for i, row := range sudoku {
		if i%3 == 0 && i != 0 {
			fmt.Println("------+-------+------")
		}
		for j, num := range row {
			if j%3 == 0 && j != 0 {
				fmt.Print("| ")
			}
			if num == 0 {
				fmt.Print(". ")
			} else {
				fmt.Printf("%d ", num)
			}
		}
		fmt.Println()
	}
}
