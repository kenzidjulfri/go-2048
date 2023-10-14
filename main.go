package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	b := &Board{height: 4, width: 4, fillCount: 1}
	b.InitializeBoard(2)
	isEnding := false
	for !isEnding {
		oldBoard := b.board
		PrintBoard(os.Stdout, b)
		fmt.Print("Input: ")
		scanner.Scan()
		command := scanner.Text()
		ProcessCommand(b, command)
		isEnding = isEndGame(oldBoard, b)
	}
}

func isEndGame(oldBoard [][]int, b *Board) bool {
	if !reflect.DeepEqual(oldBoard, b.board) {
		if isWinning(b.board) {
			fmt.Println("You win! Congratulations!")
			return true
		}

		if (!b.CanFill()) {
			fmt.Println("You lose! Better luck next time")
			return true
		}
		b.FillCells(1)
	}
	return false
}

func isWinning(board [][]int) bool {
	for _, row := range board {
		for _, cell := range row {
			if cell == 2048 {
				return true
			}
		}
	}
	return false
}