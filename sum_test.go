package lab

import "testing"

func TestSum(t *testing.T) {
	if got := Sum(1, 2, 3); got != 6 {
		t.Errorf("Sum(1,2,3) = %d, want 6", got)
	}
	if got := Sum(5); got != 5 {
		t.Errorf("Sum(5) = %d, want 5", got)
	}
}
