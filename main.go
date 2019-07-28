package main

type program map[int]byte

const (
	halt byte = iota
)

func runProgram(program map[int]byte) {
	for pc := 0; ; pc++ {
		instruction := program[pc]
		if instruction == halt {
			break
		}
	}
}

func main() {}
