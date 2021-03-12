// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const DefaultMemSize = 1024

const (
	OpHALT = iota
	OpNOOP
	OpINCA
	OpDECA
)

type GoMachine struct {
	Memory [DefaultMemSize]uint64
	P      uint64 // Program Counter
	A      uint64 // Accumulator
}

func New() GoMachine {
	return GoMachine{}
}

func (g *GoMachine) Run() {
	stop := false
	for !stop {
		switch g.Memory[g.P] {
		case OpINCA:
			g.A += 1
		case OpDECA:
			g.A -= 1
		case OpHALT:
			stop = g.Memory[g.P] == OpHALT
		}
		g.P += 1
	}
}
