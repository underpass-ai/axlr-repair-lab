package lab

import "testing"

func TestMax(t *testing.T) {
	if got := Max(-3, -7); got != -3 {
		t.Errorf("Max(-3,-7) = %d, want -3", got)
	}
	if got := Max(1, 9, 4); got != 9 {
		t.Errorf("Max(1,9,4) = %d, want 9", got)
	}
	if got := Max(-7, -3); got != -3 {
		t.Errorf("Max(-7,-3) = %d, want -3", got)
	}
	if got := Max(-3); got != -3 {
		t.Errorf("Max(-3) = %d, want -3", got)
	}
	if got := Max(); got != 0 {
		t.Errorf("Max() = %d, want 0", got)
	}
}
