package main

import (
	"math/rand"
	"time"
)

func InitializeBoard(height, width, totalFilledCell int) [][]int {
	return FillCells(CreateBoard(height, width), totalFilledCell)
}

func CreateBoard(height, width int) [][]int {
	board := make([][]int, height)
	for y := range board {
		board[y] = make([]int, width)
	}

	return board
}

func FillCells(board [][]int, totalFilledCell int) [][]int {
	random := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < totalFilledCell; {
		y := random.Intn(len(board))
		x := random.Intn(len(board[y]))

		if (board[y][x] == 0) {
			i++
			board[y][x] = 2
		}
	}

	return board
}

func MergeLeft(board [][]int) [][]int {
	newBoard := CreateBoard(len(board), len(board[0]))

	for y := 0; y < len(board); y++ {
		curPos := 0
		newBoard[y][curPos] = board[y][curPos]
		for x := 1; x < len(board[y]); x++ {
			if board[y][x] != 0 {
				if newBoard[y][curPos] == 0 {
					newBoard[y][curPos] = board[y][x]
				} else if newBoard[y][curPos] == board[y][x] {
					newBoard[y][curPos] = board[y][x] * 2
					curPos++
				} else if newBoard[y][curPos] != board[y][x] {
					curPos++
					newBoard[y][curPos] = board[y][x]
				}
			}
		}
	}

	return newBoard
}

func MergeRight(board [][]int) [][]int {
	return Mirror(MergeLeft(Mirror(board)))
}

func MergeUp(board [][]int) [][]int {
	return Transpose(MergeLeft(Transpose(board)))
}

func MergeDown(board [][]int) [][]int {
	return Transpose(MergeRight(Transpose(board)))
}

func Mirror(board [][]int) [][]int {
	for y := 0; y < len(board); y++ {
		maxX := len(board[y])
		for x := 0; x < len(board[y])/2; x++ {
			board[y][x], board[y][maxX-x-1] = board[y][maxX-x-1], board[y][x]
		}
	}

	return board
}

func Transpose(board [][]int) [][]int {
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if y > x {
				board[y][x], board[x][y] = board[x][y], board[y][x]
			}
		}
	}

	return board
}