package parkingLot

import (
	"bytes"
	"gojek_takehome/car"
	"io"
	"os"
	"testing"
)

func TestLeaveWithError(t *testing.T) {
	expected := PARKING_LOT_NOT_CREATED_ERROR
	actual := Leave(1).Error()
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestParkWithError(t *testing.T) {
	expected := PARKING_LOT_NOT_CREATED_ERROR
	car := car.Create("a", "red")
	actual := Park(car).Error()
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestInitializeWithError(t *testing.T) {
	expected := WRONG_SIZE_PARKING_LOT_ERROR
	actual := Initialize(0).Error()
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestInitializeWithoutError(t *testing.T) {
	expected := ""
	actual := Initialize(2)
	if actual != nil {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual.Error())
	}
}

func TestParkWithoutError(t *testing.T) {
	expected := ""
	car := car.Create("a", "red")
	actual := Park(car)
	car1, _ := GetInstance().getCarAtSlot(1)
	if actual != nil && car != car1 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual.Error())
	}
}

func TestGetSlotNoForRegNo(t *testing.T) {
	expected := 1
	actual, _ := GetSlotNoForRegNo("a", false)
	car, _ := GetInstance().getCarAtSlot(1)
	if actual != 1 || car.GetRegNo() != "a" {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
	}
}

func TestGetRegNosForCarsWithColor(t *testing.T) {
	expected := []string{"A"}
	actual, _ := GetRegNosForCarsWithColor("red", false)

	for i, v := range actual {
		if v != expected[i] {
			t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
		}
	}

}

func TestGetRegNosForCarsWithColor2(t *testing.T) {
	expected := NOT_FOUND_ERROR
	_, err := GetRegNosForCarsWithColor("white", false)

	if err != nil && err.Error() != expected {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, err.Error())
	}

}

func TestStatus(t *testing.T) {
	// Try to Pipe the stdout of the program to match the status output against what is expected
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	err := Status()
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC
	expected := "Slot No.\tRegistration No.\tColour\n1\tA\tRed\n"
	if err != nil || expected != out {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, out)
	}
}

func TestGetSlotNosForCarsWithColor(t *testing.T) {
	expected := []int{1}
	actual, _ := GetSlotNosForCarsWithColor("red")

	for i, v := range actual {
		if v != expected[i] {
			t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
		}
	}

}

func TestLeaveWithoutError(t *testing.T) {
	expected := ""
	actual := Leave(1)
	_, exists := GetInstance().getCarAtSlot(1)
	if actual != nil && exists == false {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual.Error())
	}
}
