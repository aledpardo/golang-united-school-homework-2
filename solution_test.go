package square

import (
	"math/big"
	"testing"
)

func TestCalcSqareForSquare(t *testing.T) {
	ans := CalcSquare(3, 4)
	want := 9.0
	if ans != want {
		t.Errorf("Got %f; want %f", ans, want)
	}
}

func TestCalcSqareForTriangle(t *testing.T) {
	ans := CalcSquare(3, 3)
	want := 3.897114
	result := big.NewFloat(ans).Cmp(big.NewFloat(want))
	/*
		if result != 0 {
			t.Errorf("Got %d result", result)
		}
		if ans != want {
			t.Errorf("Got %f; want %f", ans, want)
		}
		if ans < want {
			t.Errorf("%f < %f", ans, want)
		}
		if ans > want {
			t.Errorf("%f > %f", ans, want)
		}
	*/
}
