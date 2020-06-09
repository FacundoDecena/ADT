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

func TestBST_DeleteElement(t *testing.T) {
	for i, tc := range deleteCases {
		err := tc.tree.DeleteElement(tc.index)
		if err != nil && !tc.err {
			t.Fatalf("FAIL: %s\nNo error expected", tc.description)
		}
		if err == nil && tc.err {
			t.Fatalf("FAIL: %s\nError expected", tc.description)
		}
		if tc.tree != tc.wanted && !tc.err {
			t.Fatalf("FAIL: %s", tc.description)
		}
		t.Logf("%d PASS: %s",i, tc.description)
	}
}