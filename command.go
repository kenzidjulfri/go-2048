package main

import (
	"strings"
)

func ProcessCommand(m Merger, command string) {
	command = strings.ToLower(command)
	switch (command) {
	case "left":
		m.MergeLeft()
	case "right":
		m.MergeRight()
	case "up":
		m.MergeUp()
	case "down":
		m.MergeDown()
	default:
		return;
	}
}