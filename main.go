package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

type memoryAddress uint16
type memoryWord uint16

type program map[memoryAddress]memoryWord

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

func runProgram(program program) {
	for pc := memoryAddress(0); ; pc++ {
		switch instruction := program[pc]; instruction {
		case halt:
			return
		case noop:
			continue
		case out:
			char := program[pc+1]
			fmt.Print(string(char))
		}
	}
}

func loadProgram(filename string) program {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	program := make(program)
	r := bytes.NewReader(fileBytes)
	var i memoryAddress
	var data memoryWord
	for err := binary.Read(r, binary.LittleEndian, &data); err == nil; err = binary.Read(r, binary.LittleEndian, &data) {
		program[i] = memoryWord(data)
		i++
	}

	return program
}

func main() {
	program := loadProgram("challenge.bin")
	runProgram(program)
}
