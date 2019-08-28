package solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	obj := Constructor(3)
	var ret int

	if obj.EnQueue(1) != true {
		t.Fatalf("FAIL ----> Test 1, expects: true, ret: false")
	}
	if obj.EnQueue(2) != true {
		t.Fatalf("FAIL ----> Test 2, expects: true, ret: false")
	}
	ret = obj.Rear()
	if ret != 2 {
		t.Fatalf("FAIL ----> Test 3, expects: 2, ret: %v", ret)
	}
	if obj.EnQueue(3) != true {
		t.Fatalf("FAIL ----> Test 4, expects: true, ret: false")
	}
	if obj.EnQueue(4) == true {
		t.Fatalf("FAIL ----> Test 5, expects: false, ret: true")
	}
	ret = obj.Rear()
	if ret != 3 {
		t.Fatalf("FAIL ----> Test 6, expects: 3, ret: %v", ret)
	}
	ret = obj.Front()
	if ret != 1 {
		t.Fatalf("FAIL ----> Test 7, expects: 1, ret: %v", ret)
	}
	if obj.IsFull() != true {
		t.Fatalf("FAIL ----> Test 8, expects: true, ret: false")
	}
	if obj.IsEmpty() == true {
		t.Fatalf("FAIL ----> Test 9, expects: true, ret: false")
	}
	if obj.DeQueue() != true {
		t.Fatalf("FAIL ----> Test 10, expects: true, ret: false")
	}
	if obj.DeQueue() != true {
		t.Fatalf("FAIL ----> Test 11, expects: true, ret: false")
	}
	if obj.DeQueue() != true {
		t.Fatalf("FAIL ----> Test 12, expects: true, ret: false")
	}
	if obj.DeQueue() == true {
		t.Fatalf("FAIL ----> Test 13, expects: false, ret: true")
	}
	ret = obj.Rear()
	if ret != -1 {
		t.Fatalf("FAIL ----> Test 14, expects: -1, ret: %v", ret)
	}
	ret = obj.Front()
	if obj.Front() != -1 {
		t.Fatalf("FAIL ----> Test 15, expects: -1, ret: %v", ret)
	}

}
