package lib

import (
	"bytes"
	"gojek_takehome/parkingLot"
	"io"
	"os"
	"testing"
)

func TestProcessCommandWithWrongCommand(t *testing.T) {
	expected := UNSUPPORTED_COMMAND
	actual := processCommand("test").Error()
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestProcessCommandWithWrongCommandArguments(t *testing.T) {
	expected := UNSUPPORTED_COMMAND_ARGUMENTS
	actual := processCommand("park").Error()
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestProcessCommandWithWithRightCommand(t *testing.T) {
	expected := ""
	actual := processCommand("create_parking_lot 1")
	if actual != nil {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual.Error())
	}
}

func TestProcessCommandWithRightCommand2(t *testing.T) {
	expected := parkingLot.NOT_FOUND_ERROR
	actual := processCommand("slot_numbers_for_cars_with_colour 1").Error()
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestProcessCommandWithRightCommand3(t *testing.T) {
	expected := ""
	actual := processCommand("park 1 red")
	if actual != nil {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual.Error())
	}
}

func TestProcessCommandWithRightCommand4(t *testing.T) {
	expected := parkingLot.NOT_FOUND_ERROR
	actual := processCommand("leave 2").Error()
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestProcessCommandWithRightCommand5(t *testing.T) {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	err := processCommand("status")
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC
	expected := "Slot No.\tRegistration No.\tColour\n1\t1\tRed\n"
	if err != nil || expected != out {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, out)
	}
}

func TestProcessCommandWithRightCommand6(t *testing.T) {
	expected := ""
	actual := processCommand("slot_number_for_registration_number 1")
	if actual != nil {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual.Error())
	}
}

func TestProcessCommandWithRightCommand7(t *testing.T) {
	expected := ""
	actual := processCommand("registration_numbers_for_cars_with_colour red")
	if actual != nil {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual.Error())
	}
}

func TestProcessCommandWithRightCommand8(t *testing.T) {
	expected := "strconv.Atoi: parsing \"a\": invalid syntax"
	actual := processCommand("leave a").Error()
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestProcessCommandWithRightCommand9(t *testing.T) {
	expected := "strconv.Atoi: parsing \"a\": invalid syntax"
	actual := processCommand("create_parking_lot a").Error()
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}
