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
	height, width int
	mineCount int
	board [][]int
}

func (b *Board) InitializeBoard(totalFilledCell int) [][]int {
	return FillCells(CreateBoard(b.height, b.width), totalFilledCell)
}

func CreateBoard(height, width int) [][]int {
	board := make([][]int, height)
	for y := range board {
		board[y] = make([]int, width)
	}

	return board
}

func FillCells(board [][]int, totalFilledCells int) [][]int {
	random := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < totalFilledCells; {
		y := random.Intn(len(board))
		x := random.Intn(len(board[y]))

		if (board[y][x] == 0) {
			i++
			board[y][x] = 2
		}
	}

	return board
}

func (b *Board) MergeLeft() {
	newBoard := CreateBoard(len(b.board), len(b.board[0]))

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