package counter_test

import (
	"testing"
	"counter"
)

func TestCountInteg1(t *testing.T) {
	counter.SetCount(0)
	if counter.Count() != 1 {
		t.Error("expected 1")
	}
	if counter.Count() != 2 {
		t.Error("expected 2")
	}
	if counter.Count() != 3 {
		t.Error("expected 3")
	}
}

func TestCountInteg2(t *testing.T) {
	counter.SetCount(3)
	if counter.Count2() != 5 {
		t.Error("expected 5")
	}
	if counter.Count2() != 7 {
		t.Error("expected 7")
	}
	if counter.Count2() != 9 {
		t.Error("expected 9")
	}
}
