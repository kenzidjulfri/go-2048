package main

import "testing"

type BoardStub struct {
	calledFunctionCount map[string]int
}

func (bs *BoardStub) Init() {
	bs.calledFunctionCount = make(map[string]int)
}

func (bs *BoardStub) MergeLeft() {
	bs.calledFunctionCount["MergeLeft"]++
}

func (bs *BoardStub) MergeRight() {
	bs.calledFunctionCount["MergeRight"]++
}

func (bs *BoardStub) MergeUp() {
	bs.calledFunctionCount["MergeUp"]++
}

func (bs *BoardStub) MergeDown() {
	bs.calledFunctionCount["MergeDown"]++
}

func TestProcessCommand_MovementLeft(t *testing.T) {
	bs := &BoardStub{}
	bs.Init()
	funcName := "MergeLeft"
	command := "left"
	expectedCallCount := 1
	
	ProcessCommand(bs, command)

	if bs.calledFunctionCount[funcName] != expectedCallCount {
		t.Errorf("expected '%s' function calls: %d, got: %d", funcName, expectedCallCount, bs.calledFunctionCount[funcName])
	}
}

func TestProcessCommand_MovementRight(t *testing.T) {
	bs := &BoardStub{}
	bs.Init()
	funcName := "MergeRight"
	command := "right"
	expectedCallCount := 1
	
	ProcessCommand(bs, command)

	if bs.calledFunctionCount[funcName] != expectedCallCount {
		t.Errorf("expected '%s' function calls: %d, got: %d", funcName, expectedCallCount, bs.calledFunctionCount[funcName])
	}
}

func TestProcessCommand_MovementUp(t *testing.T) {
	bs := &BoardStub{}
	bs.Init()
	funcName := "MergeUp"
	command := "up"
	expectedCallCount := 1
	
	ProcessCommand(bs, command)

	if bs.calledFunctionCount[funcName] != expectedCallCount {
		t.Errorf("expected '%s' function calls: %d, got: %d", funcName, expectedCallCount, bs.calledFunctionCount[funcName])
	}
}

func TestProcessCommand_MovementDown(t *testing.T) {
	bs := &BoardStub{}
	bs.Init()
	funcName := "MergeDown"
	command := "down"
	expectedCallCount := 1
	
	ProcessCommand(bs, command)

	if bs.calledFunctionCount[funcName] != expectedCallCount {
		t.Errorf("expected '%s' function calls: %d, got: %d", funcName, expectedCallCount, bs.calledFunctionCount[funcName])
	}
}

func TestProcessCommand_Unrecognized(t *testing.T) {
	bs := &BoardStub{}
	tt := []string{"","upp"}
	funcNames := []string{"MergeLeft","MergeRight","MergeUp","MergeDown"}

	for _, tc := range tt {
		bs.Init()
		ProcessCommand(bs, tc)

		for _, funcName := range funcNames {
			if bs.calledFunctionCount[funcName] > 0 {
				t.Errorf("expected '%s' function calls: %d, got: %d", funcName, 0, bs.calledFunctionCount[funcName])
			}
		}
	}
}