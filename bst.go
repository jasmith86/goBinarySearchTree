package main

import (
	"errors"
)

// Tree interface TODO?

//// Binary Search Tree
//type BST struct {
//	head  *Node
//	count int
//}

// Node of Binary Search Tree
type Node struct {
	val   int
	left  *Node
	right *Node
}

//// Add new Node to tree. Returns number of nodes added and error
//func (b *BST) Insert(v int) (int, error) {
//	if b.head == nil { // Insert first/head Node
//		b.head = &Node{val: v}
//		b.count += 1
//		return 1, nil
//	}
//	err := b.head.Insert(v) // try to Insert Node
//	if err != nil {
//		return 0, err
//	}
//	b.count += 1
//	return 1, nil
//}

// Insert a new node starting at n.
func (n *Node) Insert(v int) error {
	if n == nil {
		n.val = v
		return nil
	}
	if v == n.val {
		return errors.New("failed trying to Insert duplicate value")
	}
	if v < n.val { // Go left
		if n.left == nil { // can Insert value
			n.left = &Node{val: v}
			return nil
		}
		return n.left.Insert(v)
	}
	// Go right. v must be > n.val
	if n.right == nil { // can Insert value
		n.right = &Node{val: v}
		return nil
	}
	return n.right.Insert(v)
}

// get height downwards from Node n
func (n *Node) Height(h ...int) int {
	height := 1
	if n == nil {
		return 0
	}
	if len(h) == 1 {
		height = h[0]
	}
	return height + maxInt(n.left.Height(), n.right.Height())
}

// Utility max function for integers
func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Get maximum value from Node n
func (n *Node) Max() (int, error) {
	if n == nil {
		return 0, errors.New("cannot get max of empty tree")
	}
	if n.right != nil {
		return n.right.Max()
	}
	return n.val, errors.New("cannot get max of empty tree")
}

// Get minimum value from Node n
func (n *Node) Min() (int, error) {
	if n == nil {
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
	rv = append(rv, n.val)
	if n.right != nil {
		rv = append(rv, n.right.InOrder()...)
	}
	return rv
}
