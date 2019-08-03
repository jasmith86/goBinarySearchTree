package main

import (
	"errors"
	"fmt"
)

// Fulfil Item interface specified in bst.go
type MyInt int

func (a MyInt) Less(than Item) bool    { return a < than.(MyInt) }
func (a MyInt) Greater(than Item) bool { return a > than.(MyInt) }
func (a MyInt) Equals(to Item) bool    { return a == to.(MyInt) }

// insertBulk is a utility function to add a slice of Items, in this case MyInt's.
// Used heavily in bst_test.go
func insertBulk(n *Node, values []MyInt) (int, error) {
	var numInserted int
	for _, v := range values {
		err := n.Insert(v)
		if err == nil {
			numInserted += 1
		}
	}
	if numInserted != len(values) {
		return numInserted, errors.New("failed trying to Insert duplicate value(s)")
	}
	return numInserted, nil
}

func main() {
	var list []MyInt
	list = []MyInt{1, 2, 0, -1, 0, 0, 0, 0, 0, 1000, 9999}
	root := &Node{}
	// Insert all the elements in list
	cnt, err := insertBulk(root, list)
	if err != nil {
		fmt.Println("Failed to insert", len(list)-cnt, "Items.")
	}
	fmt.Println("Added", cnt, "values to", root)
	fmt.Println("Height", root.Height())
	// Get max value
	if l, err := root.Max(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Max is", l)
	}
	// Get min value
	if s, err := root.Min(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Min is ", s)
	}
	// Get number of elements in the BST
	fmt.Println("Count is", root.Count())
	// Get elements according to in-order traversal
	fmt.Println("InOrder traversal yields", root.InOrder())
}
