package shapes

import "testing"

func TestShapes(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	wanted := 40.0

	if got != wanted {
		t.Errorf("got %.2f want %.2f", got, wantwantwant)
	}

}
