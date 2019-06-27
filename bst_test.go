package main

import (
	"testing"
)

func TestBSTInsertSimple(t *testing.T) {
	tests := []struct {
		name    string
		tree    *BST
		input   []int
		wantVal int
	}{
		{name: "empty root", input: []int{}, wantVal: 0},
		{name: "single root", input: []int{9}, wantVal: 9},
		{name: "single root", input: []int{2, 0, 3}, wantVal: 2},
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

func TestBSTInsertWholeTree(t *testing.T) {
	tree := &BST{}
	input := []int{10, 5, 15, 4, 6, 14, 16}

	for _, v := range input {
		if _, err := tree.Insert(v); err != nil {
			t.Errorf("failed to insert: %v", err)
		}
	}
	location := []*Node{
		tree.head,             // 10
		tree.head.left,        // 5
		tree.head.right,       // 15
		tree.head.left.left,   // 4
		tree.head.left.right,  // 6
		tree.head.right.left,  //14
		tree.head.right.right, // 16
	}
	for i, v := range input {
		if location[i].val != v {
			t.Errorf("item in wrong location %v", v)
		}
	}
}

func TestBSTHeight(t *testing.T) {
	tests := []struct {
		name       string
		tree       *BST
		input      []int
		wantHeight int
	}{
		{name: "empty root", input: []int{}, wantHeight: 0},
		{name: "single node", input: []int{9}, wantHeight: 1},
		{name: "height 2", input: []int{2, 0, 3}, wantHeight: 2},
		{name: "height 3", input: []int{3, 2, 1}, wantHeight: 3},
		{name: "height 3 all left", input: []int{3, 2, 1}, wantHeight: 3},
		{name: "height 4 all right", input: []int{1, 2, 3, 4}, wantHeight: 4},
		{name: "height 3 full", input: []int{10, 8, 7, 9, 12, 11, 13}, wantHeight: 3},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.tree = &BST{}
			for _, v := range test.input {
				if _, err := test.tree.Insert(v); err != nil {
					t.Errorf("failed to insert: %v: %v", v, err)
				}
			}
			if got := test.tree.head.Height(); got != test.wantHeight {
				t.Errorf("wrong height. want %v, got %v", test.wantHeight, got)
			}
		})
	}
}
