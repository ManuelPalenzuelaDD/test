package rewritefixture

import "testing"

func TestDivideByZeroPanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("Divide(10, 0) did not panic")
		}
	}()
	_ = Divide(10, 0)
}
