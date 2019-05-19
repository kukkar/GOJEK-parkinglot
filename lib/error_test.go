package lib

import (
	"errors"
	"testing"
)

func TestErrorWrapper(t *testing.T) {
	w := &ErrWrapper{}

	w.do(func() error {
		return errors.New("Test")
	})

	expected := "Test"
	actual := w.err.Error()
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	w.do(func() error {
		return nil
	})

	actual2 := w.err
	if actual2 != nil {
		t.Errorf("expected nil, got %s", actual2.Error())
	}
}
