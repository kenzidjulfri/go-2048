package main

import (
	"reflect"
	"testing"
)

func TestInitializeBoard(t *testing.T) {
	height := 4
	width := 4
	totalFilledCell := 2

	board := InitializeBoard(height, width, totalFilledCell)

	if len(board) != height {
		t.Logf("expected height: %d, got: %d", height, len(board))
		t.Fail()
	}

	if len(board[0]) != width {
		t.Logf("expected width: %d, got: %d", height, len(board[0]))
		t.Fail()
	}

	count := 0
	for y := range board {
		for x := range board[y] {
			if board[y][x] != 0 {
				count++
			}
		}
	}
	if count != totalFilledCell {
		t.Logf("expected filled cells: %d, got: %d", totalFilledCell, count)
		t.Fail()
	}
}

func TestMergeLeft(t *testing.T) {
	board := [][]int{
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

	board = MergeLeft(board)

	if !reflect.DeepEqual(expectedBoard, board) {
		t.Logf("expected board:\n%v\n\ngot:\n%v", expectedBoard, board)
		t.Fail()
	}
}

func TestMergeRight(t *testing.T) {
	board := [][]int{
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

	board = MergeRight(board)

	if !reflect.DeepEqual(expectedBoard, board) {
		t.Logf("expected board:\n%v\n\ngot:\n%v", expectedBoard, board)
		t.Fail()
	}
}

func TestMergeUp(t *testing.T) {
	board := [][]int{
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

	board = MergeUp(board)

	if !reflect.DeepEqual(expectedBoard, board) {
		t.Logf("expected board:\n%v\n\ngot:\n%v", expectedBoard, board)
		t.Fail()
	}
}

func TestMergeDown(t *testing.T) {
	board := [][]int{
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

	board = MergeDown(board)

	if !reflect.DeepEqual(expectedBoard, board) {
		t.Logf("expected board:\n%v\n\ngot:\n%v", expectedBoard, board)
		t.Fail()
	}
}

func TestMirror(t *testing.T) {
	board := [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	expectedBoard := [][]int{
		{3,2,1},
		{6,5,4},
		{9,8,7},
	}

	board = Mirror(board)

	if !reflect.DeepEqual(expectedBoard, board) {
		t.Logf("expected board:\n%v\n\ngot:\n%v", expectedBoard, board)
		t.Fail()
	}
}

func TestTranspose(t *testing.T) {
	board := [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	expectedBoard := [][]int{
		{1,4,7},
		{2,5,8},
		{3,6,9},
	}

	board = Transpose(board)

	if !reflect.DeepEqual(expectedBoard, board) {
		t.Logf("expected board:\n%v\n\ngot:\n%v", expectedBoard, board)
		t.Fail()
	}
}