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
	g.Memory[0] = gmachine.OpNOOP
	g.Run()
	wantP = 2
	if wantP != g.P {
		t.Errorf("want P == %d, got P == %d", wantP, g.P)
	}
}
