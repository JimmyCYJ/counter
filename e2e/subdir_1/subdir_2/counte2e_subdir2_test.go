package counter_test

import (
	"testing"
	"counter"
	"time"
)

func TestCountE2e1(t *testing.T) {
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

func TestCountE2e2(t *testing.T) {
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

func TestCountE2eSleep(t *testing.T) {
	counter.SetCount(0)
	if counter.Count() != 1 {
		t.Error("expected 1")
	}
	if counter.Count() != 2 {
		t.Error("expected 2")
	}

	time.Sleep(100 * time.Millisecond)

	if counter.Count() != 3 {
		t.Error("expected 3")
	}
}

func TestCountE2eSkip(t *testing.T) {
	counter.SetCount(0)
	if counter.Count() != 1 {
		t.Error("expected 1")
	}
	if counter.Count() != 2 {
		t.Error("expected 2")
	}

	if testing.Short() {
		t.Skip("skipping uint test in short mode.")
	}

	if counter.Count() != 3 {
		t.Error("expected 3")
	}
}

func TestCountE2eSkipAtTop(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping uint test in short mode.")
	}

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
