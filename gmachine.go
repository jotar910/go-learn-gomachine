// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const DefaultMemSize = 1024

type GoMachine struct {
	Memory [DefaultMemSize]uint64
	P uint64
}

func New() GoMachine {
	return GoMachine{}
}
