package square

import "testing"

func TestCalcSquare(t *testing.T) {
	got := CalcSquare(3.0, 6)
	want := 0.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}

	got1 := CalcSquare(4.0, 4)
	want1 := 16.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got1, want1)
	}
}
