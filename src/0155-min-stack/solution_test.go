package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)

	min1 := obj.GetMin()
	if !reflect.DeepEqual(-3, min1) {
		t.Fatalf("FAIL ----> Test 1, expects: -3, ret: %v", min1)
	}

	obj.Pop()
	top1 := obj.Top()
	if !reflect.DeepEqual(0, top1) {
		t.Fatalf("FAIL ----> Test 2, expects: 0, ret: %v", top1)
	}

	min2 := obj.GetMin()
	if !reflect.DeepEqual(-2, min2) {
		t.Fatalf("FAIL ----> Test 3, expects: -2, ret: %v", min2)
	}

	obj.Pop()
	obj.Pop()
	obj.Pop()
	top2 := obj.Top()
	if !reflect.DeepEqual(0, top2) {
		t.Fatalf("FAIL ----> Test 4, expects: 0, ret: %v", top2)
	}
	min3 := obj.GetMin()
	if !reflect.DeepEqual(0, min3) {
		t.Fatalf("FAIL ----> Test 5, expects: 0, ret: %v", top2)
	}
}
