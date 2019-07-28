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
