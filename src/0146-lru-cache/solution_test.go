package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(1, 1)
	cache.Put(2, 2)

	var ret int

	ret = cache.Get(1)
	if !reflect.DeepEqual(1, ret) {
		t.Fatalf("FAIL ----> Test 1, expects: 1, ret: %v", ret)
	}

	cache.Put(3, 3)

	ret = cache.Get(2)
	if !reflect.DeepEqual(-1, ret) {
		t.Fatalf("FAIL ----> Test 2, expects: -1, ret: %v", ret)
	}

	cache.Put(4, 4)

	ret = cache.Get(1)
	if !reflect.DeepEqual(-1, ret) {
		t.Fatalf("FAIL ----> Test 3, expects: -1, ret: %v", ret)
	}

	ret = cache.Get(3)
	if !reflect.DeepEqual(3, ret) {
		t.Fatalf("FAIL ----> Test 4, expects: 3, ret: %v", ret)
	}

	ret = cache.Get(4)
	if !reflect.DeepEqual(4, ret) {
		t.Fatalf("FAIL ----> Test 5, expects: 4, ret: %v", ret)
	}
}
