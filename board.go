package main

import (
	"math/rand"
	"time"
)

type Merger interface{
	MergeLeft()
	MergeRight()
	MergeUp()
	MergeDown()
}

type Board struct {
	height, width, fillCount int
	board [][]int
}

func (b *Board) InitializeBoard(initialFillCount int) {
	b.board = createBoard(b.height, b.width)
	b.FillCells(initialFillCount)
}

func createBoard(height, width int) [][]int {
	board := make([][]int, height)
	for y := range board {
		board[y] = make([]int, width)
	}

	return board
}

func (b *Board) FillCells(fillCount int) {
	random := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < fillCount; {
		y := random.Intn(len(b.board))
		x := random.Intn(len(b.board[y]))

		if (b.board[y][x] == 0) {
			i++
			b.board[y][x] = 2
		}
	}
}

func (b *Board) CanFill() bool {
	count := 0
	for _, row := range b.board {
		for _, cell := range row {
			if cell == 0 {
				count++
				if count >= b.fillCount {
					return true
				}
			}
		}
	}

	return false
}

func (b *Board) MergeLeft() {
	newBoard := createBoard(len(b.board), len(b.board[0]))

	for y := 0; y < len(b.board); y++ {
		curPos := 0
		newBoard[y][curPos] = b.board[y][curPos]
		for x := 1; x < len(b.board[y]); x++ {
			if b.board[y][x] != 0 {
				if newBoard[y][curPos] == 0 {
					newBoard[y][curPos] = b.board[y][x]
				} else if newBoard[y][curPos] == b.board[y][x] {
					newBoard[y][curPos] = b.board[y][x] * 2
					curPos++
				} else if newBoard[y][curPos] != b.board[y][x] {
					curPos++
					newBoard[y][curPos] = b.board[y][x]
				}
			}
		}
	}

	b.board = newBoard
}

func (b *Board) MergeRight() {
	b.mirror()
	b.MergeLeft()
	b.mirror()
}

func (b *Board) MergeUp() {
	b.transpose()
	b.MergeLeft()
	b.transpose()
}

func (b *Board) MergeDown() {
	b.transpose()
	b.MergeRight()
	b.transpose()
}

func (b *Board) mirror() {
	for y := 0; y < len(b.board); y++ {
		maxX := len(b.board[y])
		for x := 0; x < len(b.board[y])/2; x++ {
			b.board[y][x], b.board[y][maxX-x-1] = b.board[y][maxX-x-1], b.board[y][x]
		}
	}
}

func (b *Board) transpose() {
	for y := 0; y < len(b.board); y++ {
		for x := 0; x < len(b.board[y]); x++ {
			if y > x {
				b.board[y][x], b.board[x][y] = b.board[x][y], b.board[y][x]
			}
		}
	}
}