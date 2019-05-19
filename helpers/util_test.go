package helpers

import (
	"container/heap"
	"testing"
)

func TestToCamel(t *testing.T) {
	cases := [][]string{
		[]string{"test_case", "TestCase"},
		[]string{"test", "Test"},
		[]string{"TestCase", "TestCase"},
		[]string{" test  case ", "TestCase"},
		[]string{"", ""},
		[]string{"many_many_words", "ManyManyWords"},
		[]string{"AnyKind of_string", "AnyKindOfString"},
		[]string{"odd-fix", "OddFix"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := ToCamelCase(in)
		if result != out {
			t.Error("'" + result + "' != '" + out + "'")
		}
	}
}

func TestHeap(t *testing.T) {

	testHeap := IntHeap{2, 4, 6, 1}
	heap.Init(&testHeap)

	out := heap.Pop(&testHeap)
	result := 1
	if result != out {
		t.Errorf("%d != %d", result, out)
	}

	out = heap.Pop(&testHeap)
	result = 2
	if result != out {
		t.Errorf("%d != %d", result, out)
	}

	heap.Push(&testHeap, 10)
	out = heap.Pop(&testHeap)
	result = 4
	if result != out {
		t.Errorf("%d != %d", result, out)
	}

}
