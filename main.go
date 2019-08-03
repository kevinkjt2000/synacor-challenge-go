package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

type memoryAddress uint16
type memoryWord uint16

const (
	halt memoryWord = iota
	set
	push
	pop
	eq
	gt
	jmp
	jt
	jf
	add
	mult
	mod
	and
	or
	not
	rmem
	wmem
	call
	ret
	out
	in
	noop
)

type machine struct {
	memory map[memoryAddress]memoryWord
	pc     memoryAddress
}

func newMachine(filename string) *machine {
	var m machine
	m.memory = make(map[memoryAddress]memoryWord)
	fileBytes := m.ReadProgramFromFile(filename)
	m.LoadProgramIntoMemory(fileBytes)
	return &m
}

func newTestMachine(bytes []byte) *machine {
	var m machine
	m.memory = make(map[memoryAddress]memoryWord)
	m.LoadProgramIntoMemory(bytes)
	return &m
}

func (m *machine) GetVal(addr memoryAddress) memoryWord {
	return m.memory[addr]
}

func (m *machine) SetVal(addr memoryAddress, val memoryWord) {
	m.memory[addr] = val
}

func (m *machine) RunProgram() {
	for {
		switch instruction := m.GetVal(m.pc); instruction {
		case halt:
			return
		case jmp:
			m.pc = memoryAddress(m.GetVal(m.pc + 1))
		case noop:
			m.pc++
			continue
		case out:
			char := m.GetVal(m.pc + 1)
			m.pc += 2
			fmt.Print(string(char))
		}
	}
}

func (m *machine) LoadProgramIntoMemory(fileBytes []byte) {
	r := bytes.NewReader(fileBytes)
	var i memoryAddress
	for data, err := parseNextMemoryWord(r); err == nil; data, err = parseNextMemoryWord(r) {
		m.SetVal(i, data)
		i++
	}
}

func parseNextMemoryWord(r *bytes.Reader) (ret memoryWord, err error) {
	err = binary.Read(r, binary.LittleEndian, &ret)
	return
}

func (m *machine) ReadProgramFromFile(filename string) []byte {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return fileBytes
}

func main() {
	m := newMachine("challenge.bin")
	m.RunProgram()
}
