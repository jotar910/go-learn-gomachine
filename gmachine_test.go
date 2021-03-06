package gmachine_test

import (
	"gmachine"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	wantMemSize := gmachine.DefaultMemSize
	gotMemSize := len(g.Memory)
	if wantMemSize != gotMemSize {
		t.Errorf("want %d words of memory, got %d", wantMemSize, gotMemSize)
	}
	var wantP uint64 = 0
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
	var wantA uint64 = 0
	if wantA != g.A {
		t.Errorf("want initial P value %d, got %d", wantA, g.A)
	}
	var wantMemValue uint64 = 0
	gotMemValue := g.Memory[gmachine.DefaultMemSize-1]
	if wantMemValue != gotMemValue {
		t.Errorf("want last memory location to contain %d, got %d", wantMemValue, gotMemValue)
	}
}

func TestHALT(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	var wantP uint64 = 0
	if wantP != g.P {
		t.Errorf("want P == %d, got P == %d", wantP, g.P)
	}
	g.Run()
	wantP = 1
	if wantP != g.P {
		t.Errorf("want P == %d, got P == %d", wantP, g.P)
	}
}

func TestNOOP(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	var wantP uint64 = 0
	if wantP != g.P {
		t.Errorf("want P == %d, got P == %d", wantP, g.P)
	}
	g.RunProgram([]uint64{gmachine.OpNOOP})
	wantP = 2
	if wantP != g.P {
		t.Errorf("want P == %d, got P == %d", wantP, g.P)
	}
}

func TestINCA(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	var wantA uint64 = 0
	if wantA != g.A {
		t.Errorf("want A == %d, got A == %d", wantA, g.A)
	}
	g.Memory[0] = gmachine.OpINCA
	g.Run()
	wantA = 1
	if wantA != g.A {
		t.Errorf("want P == %d, got P == %d", wantA, g.A)
	}
}

func TestDECA(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	var wantA uint64 = 0
	if wantA != g.A {
		t.Errorf("want A == %d, got A == %d", wantA, g.A)
	}
	g.A = 2
	g.Memory[0] = gmachine.OpDECA
	g.Run()
	wantA = 1
	if wantA != g.A {
		t.Errorf("want P == %d, got P == %d", wantA, g.A)
	}
}

func TestSETA(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	var wantA uint64 = 0
	if wantA != g.A {
		t.Errorf("want A == %d, got A == %d", wantA, g.A)
	}
	g.Memory[0] = gmachine.OpSETA
	g.Memory[1] = 3
	g.Run()
	wantA = 3
	if wantA != g.A {
		t.Errorf("want A == %d, got A == %d", wantA, g.A)
	}
	var wantP uint64 = 3
	if wantP != g.P {
		t.Errorf("want P == %d, got P == %d", wantP, g.P)
	}
}

func TestSub2To3(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Memory[0] = gmachine.OpSETA
	g.Memory[1] = 3
	g.Memory[2] = gmachine.OpDECA
	g.Memory[3] = gmachine.OpDECA
	g.Run()
	var wantA uint64 = 1
	if wantA != g.A {
		t.Errorf("want A == %d, got A == %d", wantA, g.A)
	}
	var wantP uint64 = 5
	if wantP != g.P {
		t.Errorf("want P == %d, got P == %d", wantP, g.P)
	}
}

func TestRunProgram(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.RunProgram([]uint64{
		gmachine.OpNOOP,
		gmachine.OpHALT,
	})
	if g.P != 2 {
		t.Errorf("want P == 2, got %d", g.P)
	}

}
