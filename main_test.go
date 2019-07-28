package main

import "testing"

func TestInstructionHalt(t *testing.T) {
	program := make(program)
	program[0] = halt
	runProgram(program)
}

func TestEmptyProgramHalts(t *testing.T) {
	program := make(program)
	runProgram(program)
}
