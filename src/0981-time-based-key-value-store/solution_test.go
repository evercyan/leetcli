package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	obj := Constructor()
	var ret string

	obj.Set("foo", "bar", 1)

	ret = obj.Get("foo", 1)
	if !reflect.DeepEqual("bar", ret) {
		t.Fatalf("FAIL ----> Test 1, expects: bar, ret: %v", ret)
	}

	ret = obj.Get("foo", 3)
	if !reflect.DeepEqual("bar", ret) {
		t.Fatalf("FAIL ----> Test 2, expects: bar, ret: %v", ret)
	}

	obj.Set("foo", "bar2", 4)

	ret = obj.Get("foo", 4)
	if !reflect.DeepEqual("bar2", ret) {
		t.Fatalf("FAIL ----> Test 3, expects: bar2, ret: %v", ret)
	}

	ret = obj.Get("foo", 5)
	if !reflect.DeepEqual("bar2", ret) {
		t.Fatalf("FAIL ----> Test 4, expects: bar2, ret: %v", ret)
	}

	ret = obj.Get("foooo", 5)
	if !reflect.DeepEqual("", ret) {
		t.Fatalf("FAIL ----> Test 5, expects: , ret: %v", ret)
	}

	ret = obj.Get("foo", 0)
	if !reflect.DeepEqual("", ret) {
		t.Fatalf("FAIL ----> Test 6, expects: , ret: %v", ret)
	}
}
