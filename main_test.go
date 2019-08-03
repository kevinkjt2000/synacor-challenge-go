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

	m := newTestMachine([]byte{byte(jmp), 0, 4, 0, byte(out), 0, 'n', 0, byte(out), 0, 'y', 0})
	m.RunProgram()

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

	m := newTestMachine([]byte{byte(out), 0, 'h', 0})
	m.RunProgram()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	if string(out) != "h" {
		t.Errorf("Expected 'h', got: %s", string(out))
	}
}

func TestInstructionHalt(t *testing.T) {
	m := newTestMachine([]byte{byte(halt), 0})
	m.RunProgram()
}

func TestEmptyProgramHalts(t *testing.T) {
	m := newTestMachine([]byte{})
	m.RunProgram()
}
