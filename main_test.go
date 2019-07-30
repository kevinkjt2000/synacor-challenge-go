package main

import (
	"io/ioutil"
	"os"
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

func TestJumpInstruction(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	program := program{
		0: jmp,
		1: 4,
		2: out,
		3: memoryWord('n'),
		4: out,
		5: memoryWord('y'),
	}
	runProgram(program)

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	if string(out) != "y" {
		t.Errorf("Expected 'y', got: %s", string(out))
	}
}

func TestInstructionOut(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	program := program{
		0: out,
		1: memoryWord('h'),
	}
	runProgram(program)

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	if string(out) != "h" {
		t.Errorf("Expected 'h', got: %s", string(out))
	}
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
