package main

type program map[int]uint16

const (
	halt uint16 = iota
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
	for pc := 0; ; pc++ {
		switch instruction := program[pc]; instruction {
		case halt:
			return
		case noop:
			continue
		}
	}
}

func main() {}
