package square

import "testing"

func TestCalcSquare(t *testing.T) {
	got := CalcSquare(3.0, 6)
	want := 0.0

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
