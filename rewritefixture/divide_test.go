package rewritefixture

import "testing"

func TestDivideByZeroReturnsZero(t *testing.T) {
	if got := Divide(10, 0); got != 0 {
		t.Fatalf("Divide(10, 0) = %d, want 0", got)
	}
}
