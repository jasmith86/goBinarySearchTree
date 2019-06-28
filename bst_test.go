package main

import (
	"reflect"
	"sort"
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
		//{name: "add dupe node", input: []int{9, 9}, wantErr: true},
		//{name: "three nodes", input: []int{2, 0, 3}},
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

func TestMinMax(t *testing.T) {
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
				t.Errorf("error expectation mismatch in. got %v from %v", got, test.input)
			}
			if len(test.input) > 0 {
				if want := minInt(test.input...); want != got {
					t.Errorf("wrong Min(). got %v, want %v, from %v", got, want, test.input)

				}
			}
			// Test Max
			got, err = test.tree.Max()
			gotErr = err != nil
			if test.expectError != gotErr {
				t.Errorf("error expectation mismatch in Max(). got %v from %v", got, test.input)
			}
			if len(test.input) > 0 {
				if want := maxInt(test.input...); want != got {
					t.Errorf("wrong Max(). got %v, want %v, from %v", got, want, test.input)
				}
			}
		})
	}
}

func TestInOrder(t *testing.T) {
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
			got := test.tree.InOrder()
			sort.Ints(test.input)
			want := test.input
			if !reflect.DeepEqual(want, got) {
				if !(len(want) == 0 && len(got) == 0) { // TODO DeepEqual on empty slices
					t.Errorf("wrong InOrder(). got %+v, want %+v", got, want)
				}
			}
		})
	}
}

func TestInSearch(t *testing.T) {
	tests := []struct {
		name      string
		tree      *Node
		input     []int
		searchVal int
		wantFound bool
	}{
		{name: "empty root", input: []int{}, searchVal: 0, wantFound: false},
		{name: "empty root", input: []int{}, searchVal: 10, wantFound: false},
		{name: "empty root", input: []int{12}, searchVal: 12, wantFound: true},
		{name: "empty root", input: []int{5, 3, 10}, searchVal: 5, wantFound: true},
		{name: "empty root", input: []int{5, 3, 10}, searchVal: 10, wantFound: true},
		{name: "empty root", input: []int{5, 3, 10}, searchVal: 0, wantFound: false},
		{name: "empty root", input: []int{5, 3, 10, 0}, searchVal: 0, wantFound: true},
		{name: "empty root", input: []int{5, 3, 10, 0, 6}, searchVal: 6, wantFound: true},
		{name: "empty root", input: []int{5, 3, 10, 9}, searchVal: 9, wantFound: true},
		{name: "empty root", input: []int{5, 3, 10, 9, 11}, searchVal: 11, wantFound: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.tree = &Node{}
			for _, v := range test.input {
				if err := test.tree.Insert(v); err != nil {
					t.Errorf("failed to Insert: %v: %v", v, err)
				}
			}
			// Test Search
			wasFound, err := test.tree.Search(test.searchVal)
			if (err != nil) && len(test.input) > 0 {
				t.Errorf("search %v failed on %v, wanted %v. got %v. %v", test.searchVal, test.input, test.wantFound, wasFound, err)
			}
			if test.wantFound != wasFound {
				t.Errorf("search %v failed on %v, wanted %v. got %v. %v", test.searchVal, test.input, test.wantFound, wasFound, err)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name        string
		tree        *Node
		input       []int
		removeVal   int
		wantSucceed bool
	}{
		{name: "empty root", input: []int{}, removeVal: 0, wantSucceed: false},
		{name: "empty root", input: []int{}, removeVal: 10, wantSucceed: false},
		{name: "empty root", input: []int{12}, removeVal: 12, wantSucceed: true},
		{name: "remove root of 3", input: []int{5, 3, 10}, removeVal: 5, wantSucceed: true},
		{name: "remove right of 3", input: []int{5, 3, 10}, removeVal: 10, wantSucceed: true},
		{name: "remove left of 3", input: []int{5, 3, 10}, removeVal: 3, wantSucceed: true},
		{name: "remove non-exist of 3", input: []int{5, 3, 10}, removeVal: 0, wantSucceed: false},
		{name: "remove LL of 4", input: []int{5, 3, 10, 0}, removeVal: 0, wantSucceed: true},
		{name: "remove LR of 5", input: []int{5, 3, 10, 0, 6}, removeVal: 6, wantSucceed: true},
		{name: "remove RL of 4", input: []int{5, 3, 10, 9}, removeVal: 9, wantSucceed: true},
		{name: "remove RR of 5", input: []int{5, 3, 10, 9, 11}, removeVal: 11, wantSucceed: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.tree = &Node{}
			for _, v := range test.input {
				if err := test.tree.Insert(v); err != nil {
					t.Errorf("failed to Insert: %v: %v", v, err)
				}
			}
			// Test Remove
			//fmt.Println("Test tree bf remove", test.tree)
			test.tree = test.tree.Remove(test.removeVal)
			//fmt.Println("Test tree af remove", test.tree)

			// test that only the specified node was removed
			for _, inputVal := range test.input {
				stillPresent, _ := test.tree.Search(inputVal)
				if stillPresent && inputVal == test.removeVal {
					t.Errorf("Remove(%v) failed on %v. wanted %v. got stillPresent=%v", test.removeVal, test.input, test.wantSucceed, stillPresent)
				}
				if !stillPresent && inputVal != test.removeVal {
					t.Errorf("Remove(%v) failed and removed %v on %v. wanted %v. got stillPresent=%v", test.removeVal, inputVal, test.input, test.wantSucceed, stillPresent)
				}
			}
			// Test that 1 or 0 Nodes were removed
			if test.wantSucceed {
				want := len(test.input) - 1
				got := test.tree.Count()
				if want != got {
					t.Errorf("wrong node count. want %v, got %v in %v", want, got, test.input)
				}
			}
			// make sure the ordering is still correct
			gotOrdered := test.tree.InOrder()
			for i := 1; i < len(gotOrdered); i++ {
				//fmt.Println("hi")
				if !(gotOrdered[i-i] < gotOrdered[i]) {
					t.Errorf("Remove(%v) not in ascending order %v %v %v", test.removeVal, gotOrdered, test.input, i)
				}
			}

		})
	}
}
