package main


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
		}
	}
}

func main() {}
