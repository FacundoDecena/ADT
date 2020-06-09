package BST

import "testing"

func TestBST_Insert(t *testing.T) {
	for _, tc := range insertCases {
		err := tc.tree.Insert(tc.element)
		if tc.tree.Cursor.Apex.Index != tc.element.Index {
			t.Fatalf("FAIL: %s\nExpected %#v\nActual %#v", tc.description, tc.tree.Cursor.Apex.Index, tc.element.Index)
		}
		if err != nil && !tc.err {
			t.Fatalf("FAIL: %s\nNo error expected", tc.description)
		}
		if err == nil && tc.err {
			t.Fatalf("FAIL: %s\nError expected", tc.description)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func TestBST_GetElement(t *testing.T) {
	for _, tc := range getCases {
		actual, err := tc.tree.GetElement(tc.element.Index)
		if err != nil && !tc.err {
			t.Fatalf("FAIL: %s\nNo error expected", tc.description)
		}
		if err == nil && tc.err {
			t.Fatalf("FAIL: %s\nError expected", tc.description)
		}
		if actual != tc.element && !tc.err {
			t.Fatalf("FAIL: %s\nExpected %#v\nActual %#v", tc.description, actual, tc.element)
		}
		t.Logf("PASS: %s", tc.description)
	}
}
