package main

import "fmt"

func main() {
	var N int
	_, _ = fmt.Scan(&N)
	board := make([][]int, N)
	for i := range board {
		board[i] = make([]int, N)
	}

	answer := queen(0, board)
	fmt.Println(answer)
}

func queen(row int, board [][]int) int {
	if row == len(board) {
		return 1
	}

	count := 0
	for col, cell := range board[row] {
		if cell == 0 {
			place(board, row, col, 1)
			count += queen(row+1, board)
			place(board, row, col, -1)
		}
	}
	return count
}

func place(board [][]int, row, col, add int) {
	board[row][col] += add
	d := 1
	for {
		count := 0
		count += set(board, row-d, col+0, add)
		count += set(board, row-d, col+d, add)
		count += set(board, row+0, col+d, add)
		count += set(board, row+d, col+d, add)
		count += set(board, row+d, col+0, add)
		count += set(board, row+d, col-d, add)
		count += set(board, row+0, col-d, add)
		count += set(board, row-d, col-d, add)
		if count == 0 {
			break
		}
		d++
	}
}

func set(board [][]int, row, col, add int) int {
	N := len(board)
	if !(0 <= row && row < N) || !(0 <= col && col < N) {
		return 0
	}
	board[row][col] += add
	return 1
}
