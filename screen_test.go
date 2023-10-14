package main

import (
	"bytes"
	"testing"
)

func TestPrintBoard(t *testing.T) {
	b := &Board{}
	b.board = [][]int{
		{0,2,0,0},
		{0,32,2,0},
		{4,2,1024,2},
		{4,128,4,8},
	}
	expected := "    |   2|    |    \n" +
				"---- ---- ---- ----\n" +
				"    |  32|   2|    \n" +
				"---- ---- ---- ----\n" +
				"   4|   2|1024|   2\n" +
				"---- ---- ---- ----\n" +
				"   4| 128|   4|   8\n"

	var output bytes.Buffer

	PrintBoard(&output, b)

	if expected != output.String() {
		t.Errorf("expected filled cells:\n%s\ngot:\n%s", expected, output.String())
	}
}