package main

import (
	"testing"
)

func TestBSTInsertSimple(t *testing.T) {
	tests := []struct {
		name    string
		tree    *Node
		input   []int
		wantErr bool
	}{
		{name: "zero nodes", input: []int{}},
		{name: "one node", input: []int{9}},
		{name: "add dupe node", input: []int{9, 9}, wantErr: true},
		{name: "three nodes", input: []int{2, 0, 3}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.tree = &Node{}
			for _, v := range test.input { // build the tree
				err := test.tree.Insert(v)
				if err != nil { // got an error
					if !test.wantErr { // but we didn't want an error
						t.Errorf("got unexpected error: %v. \n%v into %v", err, v, test.tree)
					}
				}
			}
			if len(test.input) > 0 { // if we added nodes
				if !test.tree.isReady { // check the root node marked as ready
					t.Errorf("tree not ready: %+v", test.tree)
				}
				if want, got := test.input[0], test.tree.val; want != got { // make sure root node has correct value
					t.Errorf("wrong value. want %v, got %v. %+v", want, got, test.tree)
				}
			}

		})
	}
}

// Check tree structure
func TestBSTInsertCheckCorrectStructure(t *testing.T) {
	tree := &Node{}
	input := []int{10, 5, 15, 4, 6, 14, 16}

	for _, v := range input {
		if err := tree.Insert(v); err != nil {
			t.Errorf("failed to Insert: %v", err)
		}
	}
	location := []*Node{
		tree,             // 10
		tree.left,        // 5
		tree.right,       // 15
		tree.left.left,   // 4
		tree.left.right,  // 6
		tree.right.left,  //14
		tree.right.right, // 16
	}
	for i, v := range input {
		if location[i].val != v {
			t.Errorf("item in wrong location %v", v)
		}
	}
}

func TestBSTHeightCount(t *testing.T) {
	tests := []struct {
		name       string
		tree       *Node
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
			test.tree = &Node{}
			for _, v := range test.input {
				if err := test.tree.Insert(v); err != nil {
					t.Errorf("failed to Insert: %v: %v", v, err)
				}
			}
			// Test Height
			if got := test.tree.Height(); got != test.wantHeight {
				t.Errorf("wrong height. want %v, got %v in %v", test.wantHeight, got, test.input)
			}
			// Test Count
			if want, got := len(test.input), test.tree.Count(); want != got {
				t.Errorf("wrong node count. want %v, got %v in %v", want, got, test.input)
			}
		})
	}
}

func TestBSTMinMax(t *testing.T) {
	tests := []struct {
		name        string
		tree        *Node
		input       []int
		expectError bool
	}{
		{name: "empty root", input: []int{}, expectError: true},
		{name: "single node", input: []int{9}},
		{name: "height 2", input: []int{2, 0, 3}},
		{name: "height 3", input: []int{3, 2, 1}},
		{name: "height 3 all left", input: []int{3, 2, 1}},
		{name: "height 4 all right", input: []int{1, 2, 3, 4}},
		{name: "height 3 full", input: []int{10, 8, 7, 9, 12, 11, 13}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.tree = &Node{}
			for _, v := range test.input {
				if err := test.tree.Insert(v); err != nil {
					t.Errorf("failed to Insert: %v: %v", v, err)
				}
			}
			// Test Min
			got, err := test.tree.Min()
			gotErr := err != nil
			if test.expectError != gotErr {
				t.Errorf("error expectation mismatch. got %v from %v", got, test.input)
			}
			if len(test.input) > 0 {
				if want := minInt(test.input...); want != got {
					t.Errorf("wrong Min(). got %v, want %v, from %v", got, want, test.input)

				}
			}
		})
	}
}
