package stackvm

import "fmt"

type I32 uint32

/*
 * Instruction format
 * header: 2 bits
 * data: 30 bits
 *
 * header format:
 * 0 => positive integer
 * 1 => primitive instruction
 * 2 => negative integer
 * 3 => undefined
 * */

//NewStackVM create a new stack VM
func NewStackVM() *StackVM {
	stackVM := &StackVM{pc: 100, memory: make([]I32, 1000000), running: 1}
	return stackVM
}

//VM virtual machine interface
type VM interface {
	fetch()
	decode()
	execute()
	doPrimitive()
	getData(instruction I32) I32
	getType(instruction I32) I32

	Run()
	LoadProgram(prog []I32)
}

//StackVM the stack vm
type StackVM struct {
	pc      I32
	sp      I32
	memory  []I32
	running I32
	dat     I32
	typ     I32
}

func (s *StackVM) fetch() {
	s.pc++
}

func (s *StackVM) decode() {
	s.typ = s.getType(s.memory[s.pc])
	s.dat = s.getData(s.memory[s.pc])
}

func (s *StackVM) execute() {
	if s.typ == 0 || s.typ == 2 {
		s.sp++
		s.memory[s.sp] = s.dat
	} else {
		s.doPrimitive()
	}
}

func (s *StackVM) doPrimitive() {
	switch s.dat {
	case 0: // halt
		fmt.Println("Halting")
		s.running = 0
	case 1: // add
		fmt.Printf("add %d %d \n", s.memory[s.sp-1], s.memory[s.sp])
		s.memory[s.sp-1] = s.memory[s.sp-1] + s.memory[s.sp]
		s.sp--

	case 2: // sub
		fmt.Printf("sub %d %d \n", s.memory[s.sp-1], s.memory[s.sp])
		s.memory[s.sp-1] = s.memory[s.sp-1] - s.memory[s.sp]

		s.sp--

	case 3: // mul
		fmt.Printf("mul %d %d \n", s.memory[s.sp-1], s.memory[s.sp])
		s.memory[s.sp-1] = s.memory[s.sp-1] * s.memory[s.sp]
		s.sp--

	case 4: // div
		fmt.Printf("div %d %d \n", s.memory[s.sp-1], s.memory[s.sp])
		s.memory[s.sp-1] = s.memory[s.sp-1] / s.memory[s.sp]
		s.sp--

	}
}

func (s *StackVM) getData(instruction I32) I32 {
	var data I32 = 0x3fffffff
	data = data & instruction

	return data
}

func (s *StackVM) getType(instruction I32) I32 {
	var typ I32 = 0xc0000000
	typ = (typ & instruction) >> 30

	return typ
}

//LoadProgram Load a program
func (s *StackVM) LoadProgram(prog []I32) {
	var a I32
	for i := 0; i < len(prog); i++ {
		s.memory[s.pc+a] = prog[i]
		a++
	}
}

//Run run vm
func (s *StackVM) Run() {
	s.pc--
	for s.running == 1 {
		s.fetch()
		s.decode()
		s.execute()
	}
}
