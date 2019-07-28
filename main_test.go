package main

import (
	"testing"
	"time"
)

func panicOnTimeout() {
	<-time.After(500 * time.Millisecond)
	panic("Test timed out")
}

func init() {
	go panicOnTimeout()
}

func TestInstructionHalt(t *testing.T) {
	program := program{
		0: halt,
	}
	runProgram(program)
}

func TestEmptyProgramHalts(t *testing.T) {
	program := program{}
	runProgram(program)
}
