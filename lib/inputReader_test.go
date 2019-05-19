package lib

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestReadAndProcessFromInput(t *testing.T) {
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()

	_, err = io.WriteString(in, "test")
	if err != nil {
		t.Fatal(err)
	}

	_, err = in.Seek(0, os.SEEK_SET)
	if err != nil {
		t.Fatal(err)
	}

	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	ReadAndProcessFromInput(in)
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC
	expected := UNSUPPORTED_COMMAND + "\n"
	if expected != out {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, out)
	}
}

func TestReadAndProcessStdIn(t *testing.T) {
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()

	_, err = io.WriteString(in, "test")
	if err != nil {
		t.Fatal(err)
	}

	_, err = in.Seek(0, os.SEEK_SET)
	if err != nil {
		t.Fatal(err)
	}

	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	old1 := os.Stdin
	os.Stdin = in
	ReadAndProcessStdIn()
	os.Stdin = old1
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC
	expected := UNSUPPORTED_COMMAND + "\n"
	if expected != out {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, out)
	}
}

func TestReadAndProcessFromFile(t *testing.T) {
	d1 := []byte("testn")
	ioutil.WriteFile("/tmp/dat1", d1, 0644)

	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	ReadAndProcessFromFile("/tmp/dat1")
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC
	expected := UNSUPPORTED_COMMAND + "\n"
	if expected != out {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, out)
	}

}
