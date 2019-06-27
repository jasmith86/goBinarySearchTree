package main

import "testing"

func TestBSTInsert(t *testing.T) {
	tests := []struct {
		name    string
		tree    *BST
		input   []int
		wantVal int
	}{
		{name: "empty root", input: []int{}, wantVal: 0},
		{name: "single root", input: []int{9}, wantVal: 9},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.tree = &BST{}
			for _, v := range test.input {
				if _, err := test.tree.Insert(v); err != nil {
					t.Errorf("failed to insert: %v", err)
				}
			}
			if test.tree.count != len(test.input) {
				t.Errorf("wrong count. want %v, got %v", len(test.input), test.tree.count)
			}
			if len(test.input) > 0 {
				if test.tree.head.val != test.wantVal {
					t.Errorf("wrong value. want %v, got %v", test.tree.head.val, test.wantVal)
				}
			}
		})
	}
}
