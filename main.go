package main

import (
	"fmt"
	"io"
)

func PrintBoard(output io.Writer, board [][]int) {
	for y, rows := range board {
		line := ""
		for x, cell := range rows {
			if x != 0 {
				fmt.Fprint(output, "|")
				line += " "
			}

			if cell == 0 {
				fmt.Fprint(output, "    ")
			} else if cell/1000 >= 1 {
				fmt.Fprint(output, cell)
			} else if cell/100 >= 1 {
				fmt.Fprintf(output, "%s%d", " ", cell)
			} else if cell/10 >= 1 {
				fmt.Fprintf(output, "%s%d", "  ", cell)
			} else {
				fmt.Fprintf(output, "%s%d", "   ", cell)
			}
			line += "----"
		}
		fmt.Fprintln(output)
		if y != len(board) - 1 {
			fmt.Fprintln(output, line)
		}
	}
}