package main

import (
	"sort"
	"testing"
)

// Ensure items are inserted into the tree
func TestInsertSingle(t *testing.T) {
	tests := []struct {
		name    string
		tree    *Node
		input   []MyInt
		wantErr bool
	}{
		{name: "zero nodes", input: []MyInt{}},
		{name: "one node", input: []MyInt{9}},
		{name: "add dupe node", input: []MyInt{9, 9}, wantErr: true},
		{name: "three nodes", input: []MyInt{2, 0, 3}},
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
func TestInsertBulkCheckCorrectStructure(t *testing.T) {
	tree := &Node{}
	input := []MyInt{10, 5, 15, 4, 6, 14, 16}

	num, err := insertBulk(tree, input)
	if err != nil {
		t.Errorf("failed to Insert all %v/%v. %v", num, len(input), err)
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

// Make sure that we form a tree of the expected height
func TestHeightCount(t *testing.T) {
	tests := []struct {
		name       string
		tree       *Node
		input      []MyInt
		wantHeight int
	}{
		{name: "empty root", input: []MyInt{}, wantHeight: 0},
		{name: "single node", input: []MyInt{9}, wantHeight: 1},
		{name: "height 2", input: []MyInt{2, 0, 3}, wantHeight: 2},
		{name: "height 3", input: []MyInt{3, 2, 1}, wantHeight: 3},
		{name: "height 3 all left", input: []MyInt{3, 2, 1}, wantHeight: 3},
		{name: "height 4 all right", input: []MyInt{1, 2, 3, 4}, wantHeight: 4},
		{name: "height 3 full", input: []MyInt{10, 8, 7, 9, 12, 11, 13}, wantHeight: 3},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.tree = &Node{}
			_, _ = insertBulk(test.tree, test.input)
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

// Test find Min/Max functionality of the BST
func TestMinMax(t *testing.T) {
	tests := []struct {
		name        string
		tree        *Node
		input       []MyInt
		expectError bool
	}{
		{name: "empty root", input: []MyInt{}, expectError: true},
		{name: "single node", input: []MyInt{9}},
		{name: "height 2", input: []MyInt{2, 0, 3}},
		{name: "height 3", input: []MyInt{3, 2, 1}},
		{name: "height 3 all left", input: []MyInt{3, 2, 1}},
		{name: "height 4 all right", input: []MyInt{1, 2, 3, 4}},
		{name: "height 3 full", input: []MyInt{10, 8, 7, 9, 12, 11, 13}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.tree = &Node{}
			_, _ = insertBulk(test.tree, test.input)
			// Test Min
			got, err := test.tree.Min()
			gotErr := err != nil
			if test.expectError != gotErr {
				t.Errorf("error expectation mismatch in. got %v from %v", got, test.input)
			}
			if len(test.input) > 0 {
				if want := minMyInt(test.input...); want != got {
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
				if want := maxMyInt(test.input...); want != got {
					t.Errorf("wrong Max(). got %v, want %v, from %v", got, want, test.input)
				}
			}
		})
	}
}

// Fulfill interface for Sort, see: https://gobyexample.com/sorting-by-functions
type myIntSlice []MyInt

func (s myIntSlice) Len() int           { return len(s) }
func (s myIntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s myIntSlice) Less(i, j int) bool { return s[i] < s[j] }

// Ensure that in-order traversal works as expected
func TestInOrder(t *testing.T) {
	tests := []struct {
		name        string
		tree        *Node
		input       myIntSlice
		expectError bool
	}{
		{name: "empty root", input: []MyInt{}, expectError: true},
		{name: "single node", input: []MyInt{9}},
		{name: "height 2", input: []MyInt{2, 0, 3}},
		{name: "height 3", input: []MyInt{3, 2, 1}},
		{name: "height 3 all left", input: []MyInt{3, 2, 1}},
		{name: "height 4 all right", input: []MyInt{1, 2, 3, 4}},
		{name: "height 3 full", input: []MyInt{10, 8, 7, 9, 12, 11, 13}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.tree = &Node{}
			_, _ = insertBulk(test.tree, test.input)
			got := test.tree.InOrder()
			sort.Sort(myIntSlice(test.input)) // Use Sort to verify traversal
			want := test.input
			for i := 0; i < len(want)-1; i++ {
				if !(got[i].Equals(want[i])) {
					t.Errorf("wrong InOrder(). got %+v, want %+v", got, want)
				}
			}
		})
	}
}

// Test BST's search function
func TestSearch(t *testing.T) {
	tests := []struct {
		name      string
		tree      *Node
		input     []MyInt
		searchVal MyInt
		wantFound bool
	}{
		{name: "empty root", input: []MyInt{}, searchVal: 0, wantFound: false},
		{name: "empty root", input: []MyInt{}, searchVal: 10, wantFound: false},
		{name: "empty root", input: []MyInt{12}, searchVal: 12, wantFound: true},
		{name: "empty root", input: []MyInt{5, 3, 10}, searchVal: 5, wantFound: true},
		{name: "empty root", input: []MyInt{5, 3, 10}, searchVal: 10, wantFound: true},
		{name: "empty root", input: []MyInt{5, 3, 10}, searchVal: 0, wantFound: false},
		{name: "empty root", input: []MyInt{5, 3, 10, 0}, searchVal: 0, wantFound: true},
		{name: "empty root", input: []MyInt{5, 3, 10, 0, 6}, searchVal: 6, wantFound: true},
		{name: "empty root", input: []MyInt{5, 3, 10, 9}, searchVal: 9, wantFound: true},
		{name: "empty root", input: []MyInt{5, 3, 10, 9, 11}, searchVal: 11, wantFound: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.tree = &Node{}
			_, _ = insertBulk(test.tree, test.input)
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
		input       []MyInt
		removeVal   MyInt
		wantSucceed bool
	}{
		{name: "empty root", input: []MyInt{}, removeVal: 0, wantSucceed: false},
		{name: "empty root", input: []MyInt{}, removeVal: 10, wantSucceed: false},
		{name: "empty root", input: []MyInt{12}, removeVal: 12, wantSucceed: true},
		{name: "remove root of 3", input: []MyInt{5, 3, 10}, removeVal: 5, wantSucceed: true},
		{name: "remove right of 3", input: []MyInt{5, 3, 10}, removeVal: 10, wantSucceed: true},
		{name: "remove left of 3", input: []MyInt{5, 3, 10}, removeVal: 3, wantSucceed: true},
		{name: "remove non-exist of 3", input: []MyInt{5, 3, 10}, removeVal: 0, wantSucceed: false},
		{name: "remove LL of 4", input: []MyInt{5, 3, 10, 0}, removeVal: 0, wantSucceed: true},
		{name: "remove LR of 5", input: []MyInt{5, 3, 10, 0, 6}, removeVal: 6, wantSucceed: true},
		{name: "remove RL of 4", input: []MyInt{5, 3, 10, 9}, removeVal: 9, wantSucceed: true},
		{name: "remove RR of 5", input: []MyInt{5, 3, 10, 9, 11}, removeVal: 11, wantSucceed: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.tree = &Node{}
			_, _ = insertBulk(test.tree, test.input)
			test.tree = test.tree.Remove(test.removeVal)
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
				if !(gotOrdered[i-i].Less(gotOrdered[i])) {
					t.Errorf("Remove(%v) not in ascending order %v %v %v", test.removeVal, gotOrdered, test.input, i)
				}
			}

		})
	}
}

// Utility max/min function for myInt
func maxMyInt(x ...MyInt) MyInt {
	max := x[0]
	for _, v := range x {
		if v > max {
			max = v
		}
	}
	return max
}
func minMyInt(x ...MyInt) MyInt {
	min := x[0]
	for _, v := range x {
		if v < min {
			min = v
		}
	}
	return min
}
