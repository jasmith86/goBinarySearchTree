package main

import (
	"errors"
)

// Tree interface TODO?

// Node of Binary Search Tree
type Node struct {
	val     int
	isReady bool
	left    *Node
	right   *Node
}

// Mark current Node as ready (n.val has been set -- not default value)
func (n *Node) setVal(v int) {
	n.val = v
	n.isReady = true
}

// Insert a new node starting at n.
func (n *Node) Insert(v int) error {
	if !n.isReady { // make sure the root node is set
		n.setVal(v)
		return nil
	}
	if v == n.val {
		return errors.New("failed trying to Insert duplicate value")
	}
	if v < n.val { // Go left
		if n.left == nil { // can Insert value
			n.left = &Node{}
			n.left.setVal(v)
			return nil
		}
		return n.left.Insert(v)
	}
	// Go right. v must be > n.val
	if n.right == nil { // can Insert value
		n.right = &Node{}
		n.right.setVal(v)
		return nil
	}
	return n.right.Insert(v)
}

// get number of nodes from n (inclusive)
func (n *Node) Count(c ...int) int {
	if !n.isReady {
		return 0
	}
	count := 1
	if n.left != nil {
		count += n.left.Count()
	}
	if n.right != nil {
		count += n.right.Count()
	}
	return count
}

// get height downwards from Node n
func (n *Node) Height(h ...int) int {
	if !n.isReady {
		return 0
	}
	lh, rh := 0, 0
	if n.left != nil {
		lh = n.left.Height()
	}
	if n.right != nil {
		rh = n.right.Height()
	}
	//return height + maxInt(n.left.Height(), n.right.Height())
	return 1 + maxInt(lh, rh)
}

// Get maximum value from Node n
func (n *Node) Max() (int, error) {
	if !n.isReady {
		return 0, errors.New("cannot get max of empty tree")
	}
	if n.right != nil {
		return n.right.Max()
	}
	return n.val, nil
}

// Get minimum value from Node n
func (n *Node) Min() (int, error) {
	if !n.isReady {
		return 0, errors.New("cannot get min of empty tree")
	}
	if n.left != nil {
		return n.left.Min()
	}
	return n.val, nil
}

// Walk from n in order. Returns values in ascending order.
func (n *Node) InOrder(prev ...int) []int {
	var rv []int
	if n.left != nil {
		rv = append(rv, n.left.InOrder()...)
	}
	if n.isReady {
		rv = append(rv, n.val)
	}
	if n.right != nil {
		rv = append(rv, n.right.InOrder()...)
	}
	return rv
}

// Search for a value

// Utility max function for integers
func maxInt(x ...int) int {
	max := x[0]
	for _, v := range x {
		if v > max {
			max = v
		}
	}
	return max
}
func minInt(x ...int) int {
	min := x[0]
	for _, v := range x {
		if v < min {
			min = v
		}
	}
	return min
}
