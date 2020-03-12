package shapes

import "testing"

func TestShapes(t *testing.T) {

	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	wanted := 40.0

	if got != wanted {
		t.Errorf("got %.2f want %.2f", got, wanted)
	}

}
