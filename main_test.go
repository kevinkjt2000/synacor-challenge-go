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

func assemble(words ...memoryWord) []byte {
	ret := make([]byte, len(words) * 2)
	for i, word := range words {
		ret[i*2] = byte(word)
	}
	return ret
}

func TestJumpInstruction(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	m := newTestMachine(assemble(jmp, 4, out, 'n', out, 'y'))
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

	m := newTestMachine(assemble(out, 'h'))
	m.RunProgram()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	if string(out) != "h" {
		t.Errorf("Expected 'h', got: %s", string(out))
	}
}

func TestInstructionHalt(t *testing.T) {
	m := newTestMachine(assemble(halt))
	m.RunProgram()
}

func TestEmptyProgramHalts(t *testing.T) {
	m := newTestMachine(assemble())
	m.RunProgram()
}
