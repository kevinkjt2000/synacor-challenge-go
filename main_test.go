package main

import "testing"

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
