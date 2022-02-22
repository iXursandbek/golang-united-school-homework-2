package square

import "testing"

func TestCalcSquare(t *testing.T) {
	got := CalcSquare(3.0, 3)
	want := 3.8971143170299736

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
