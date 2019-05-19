package car

import (
	"testing"
)

func TestCar(t *testing.T) {
	car := Create("a", "red")

	expected := "a"
	actual := car.GetRegNo()
	if actual != expected {
		t.Errorf("%d != %d", actual, expected)
	}

	expected = "red"
	actual = car.GetColor()
	if actual != expected {
		t.Errorf("%d != %d", actual, expected)
	}

}
