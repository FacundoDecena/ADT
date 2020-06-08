package Stack

import (
	"reflect"
	"testing"
)

func TestStack_Push(t *testing.T) {
	for _, tc := range pushTestCases {
		tc.stack.Push(tc.item)
		if !reflect.DeepEqual(tc.stack, tc.stackWanted) {
			t.Fatalf("FAIL: %s\nExpected %#v\nActual %#v", tc.description, tc.stack, tc.stackWanted)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func TestStack_Peek(t *testing.T) {
	for _, tc := range peekTestCases {
		actual, err := tc.stack.Peek()
		if err != nil && !tc.err {
			t.Fatalf("FAIL: %s\nExpected error", tc.description)
		}
		if err == nil && tc.err {
			t.Fatalf("FAIL: %s\nNot expecting error", tc.description)
		}
		if actual != tc.itemWanted {
			t.Fatalf("FAIL: %s\nExpected: %#v\nActual:%#v", tc.description, tc.itemWanted, actual)
		}
		if !reflect.DeepEqual(tc.stack, tc.stackWanted) {
			t.Fatalf("FAIL: %s\nExpected: %#v\nActual:%#v", tc.description, tc.stackWanted, tc.stack)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func TestStack_Pop(t *testing.T) {
	for _, tc := range popTestCases {
		actual, err := tc.stack.Pop()
		if err != nil && !tc.err {
			t.Fatalf("FAIL: %s\nExpected error", tc.description)
		}
		if err == nil && tc.err {
			t.Fatalf("FAIL: %s\nNot expecting error", tc.description)
		}
		if actual != tc.itemWanted {
			t.Fatalf("FAIL: %s\nExpected: %#v\nActual:%#v", tc.description, tc.itemWanted, actual)
		}
		if !reflect.DeepEqual(tc.stack, tc.stackWanted) {
			t.Fatalf("FAIL: %s\nExpected: %#v\nActual:%#v", tc.description, tc.stackWanted, tc.stack)
		}
		t.Logf("PASS: %s", tc.description)
	}
}
