package main

import (
	"reflect"
	"testing"
)

func TestInitializeBoard(t *testing.T) {
	height := 4
	width := 4
	totalFilledCells := 2
	b := &Board{height, width, 0, make([][]int, 0)}

	b.InitializeBoard(totalFilledCells)

	if len(b.board) != height {
		t.Errorf("expected height: %d, got: %d", height, len(b.board))
	}

	if len(b.board[0]) != width {
		t.Errorf("expected width: %d, got: %d", height, len(b.board[0]))
	}

	count := 0
	for y := range b.board {
		for x := range b.board[y] {
			if b.board[y][x] != 0 {
				count++
			}
		}
	}
	if count != totalFilledCells {
		t.Errorf("expected filled cells: %d, got: %d", totalFilledCells, count)
	}
}

func TestMergeLeft(t *testing.T) {
	b := &Board{}
	b.board = [][]int{
		{2,0,2,0}, // check if merging 2 of the same number ignore 0s
		{2,2,2,0}, // check if merging 3 of the same number will prioritize the first twos on the same direction
		{2,2,2,2}, // check if merging 4 of the same number will have the position right
		{2,2,4,8}, // check if merging 4 of two number 2s followed by number 4 and 8 won't affect the final board
	}
	expectedBoard:= [][]int{
		{4,0,0,0},
		{4,2,0,0},
		{4,4,0,0},
		{4,4,8,0},
	}

	b.MergeLeft()

	if !reflect.DeepEqual(expectedBoard, b.board) {
		t.Errorf("expected board:\n%v\n\ngot:\n%v", expectedBoard, b.board)
	}
}

func TestMergeRight(t *testing.T) {
	b := &Board{}
	b.board = [][]int{
		{2,0,2,0},
		{2,2,2,0},
		{2,2,2,2},
		{2,2,4,8},
	}
	expectedBoard:= [][]int{
		{0,0,0,4},
		{0,0,2,4},
		{0,0,4,4},
		{0,4,4,8},
	}

	b.MergeRight()

	if !reflect.DeepEqual(expectedBoard, b.board) {
		t.Errorf("expected board:\n%v\n\ngot:\n%v", expectedBoard, b.board)
	}
}

func TestMergeUp(t *testing.T) {
	b := &Board{}
	b.board = [][]int{
		{2,0,2,0},
		{2,2,2,0},
		{2,2,2,2},
		{2,2,4,8},
	}
	expectedBoard:= [][]int{
		{4,4,4,2},
		{4,2,2,8},
		{0,0,4,0},
		{0,0,0,0},
	}

	b.MergeUp()

	if !reflect.DeepEqual(expectedBoard, b.board) {
		t.Errorf("expected board:\n%v\n\ngot:\n%v", expectedBoard, b.board)
	}
}

func TestMergeDown(t *testing.T) {
	b := &Board{}
	b.board = [][]int{
		{2,0,2,0},
		{2,2,2,0},
		{2,2,2,2},
		{2,2,4,8},
	}
	expectedBoard:= [][]int{
		{0,0,0,0},
		{0,0,2,0},
		{4,2,4,2},
		{4,4,4,8},
	}

	b.MergeDown()

	if !reflect.DeepEqual(expectedBoard, b.board) {
		t.Errorf("expected board:\n%v\n\ngot:\n%v", expectedBoard, b.board)
	}
}

func TestMirror(t *testing.T) {
	b := &Board{}
	b.board = [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	expectedBoard := [][]int{
		{3,2,1},
		{6,5,4},
		{9,8,7},
	}

	b.mirror()

	if !reflect.DeepEqual(expectedBoard, b.board) {
		t.Errorf("expected board:\n%v\n\ngot:\n%v", expectedBoard, b.board)
	}
}

func TestTranspose(t *testing.T) {
	b := &Board{}
	b.board = [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	expectedBoard := [][]int{
		{1,4,7},
		{2,5,8},
		{3,6,9},
	}

	b.transpose()

	if !reflect.DeepEqual(expectedBoard, b.board) {
		t.Errorf("expected board:\n%v\n\ngot:\n%v", expectedBoard, b.board)
	}
}