package main

import stackvm "github.com/Jon3123/Go-Stack-VM/pkg/stack-vm"

func main() {
	vm := stackvm.NewStackVM()
	prog := []stackvm.I32{3, 4, 0x40000001, 0x40000000}
	vm.LoadProgram(prog)
	vm.Run()
}
