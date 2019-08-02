// Implements a simple Binary Search Tree. Can Insert, Remove, Search, find Min, find Max
package main

import "errors"

// https://github.com/google/btree/blob/master/btree.go
type Item interface {
	Less(than Item) bool
	Greater(than Item) bool
	Equals(to Item) bool
}

type Node struct {
	val     Item
	isReady bool
	left    *Node
	right   *Node
}

// Mark current Node as ready (n.val has been set -- not default value)
func (n *Node) setVal(v Item) {
	n.val = v
	n.isReady = true
}

// Mark current Node as not ready (equiv to &Node{})
func (n *Node) unsetVal() {
	//n.val = 0 # todo: problem??
	n.isReady = false
	n.left = nil
	n.right = nil
}

// Insert a new node starting at n.
func (n *Node) Insert(v Item) error {
	if !n.isReady { // make sure the root node is set
		n.setVal(v)
		return nil
	}
	if v.Equals(n.val) {
		return errors.New("failed trying to Insert duplicate value")
	}
	if v.Less(n.val) { // Go left
		if n.left == nil { // can Insert value
			n.left = &Node{}
			n.left.setVal(v)
			return nil
		}
		return n.left.Insert(v)
	}
	// Go right, since v must be > n.val
	if n.right == nil { // can Insert value
		n.right = &Node{}
		n.right.setVal(v)
		return nil
	}
	return n.right.Insert(v)
}

//// Insert multiple nodes. If duplicate values specified, will keep inserting and return a single Error
//func (n *Node) insertBulk(values []Item) (int, error) { // todo Remove?
//	var numInserted int
//	for _, v := range values {
//		err := n.Insert(v)
//		if err == nil {
//			numInserted += 1
//		}
//	}
//	if numInserted != len(values) {
//		return numInserted, errors.New("failed trying to Insert duplicate value(s)")
//	}
//	return numInserted, nil
//}

// Get number of nodes from Node n (inclusive)
func (n *Node) Count() int {
	if !n.isReady {
		return 0 // count of empty tree is 0
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

// Get height downwards from Node n (inclusive)
func (n *Node) Height() int {
	if !n.isReady {
		return 0 // height of empty tree is 0
	}
	leftHeight, rightHeight := 0, 0
	if n.left != nil {
		leftHeight = n.left.Height()
	}
	if n.right != nil {
		rightHeight = n.right.Height()
	}
	return 1 + maxInt(leftHeight, rightHeight)
}

// Get maximum value from Node n
func (n *Node) Max() (Item, error) {
	if !n.isReady {
		return nil, errors.New("cannot get Max() of empty tree")
	}
	if n.right != nil {
		return n.right.Max()
	}
	return n.val, nil
}

// Get minimum value from Node n
func (n *Node) Min() (Item, error) {
	if !n.isReady {
		return nil, errors.New("cannot get Min() of empty tree")
	}
	if n.left != nil {
		return n.left.Min()
	}
	return n.val, nil
}

// Walk from n in order. Returns values in ascending order.
func (n *Node) InOrder() []Item { // TODO take func as arg?
	var rv []Item
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

// Search for a value starting at Node n (inclusive)
func (n *Node) Search(searchVal Item) (bool, error) {
	if !n.isReady {
		return false, errors.New("cannot search empty tree")
	}
	if searchVal.Equals(n.val) {
		return true, nil
	}
	if searchVal.Less(n.val) {
		// search left
		if n.left != nil {
			return n.left.Search(searchVal)
		}
	}
	// search right -- removeVal > n.val
	if n.right != nil {
		return n.right.Search(searchVal)
	}
	//
	return false, nil
}

// Remove node with specified value. Returns a new root node.
func (n *Node) Remove(removeVal Item) *Node {
	if !n.isReady { // No nodes in tree
		return n
	}
	if removeVal.Less(n.val) { // Search left
		if n.left != nil {
			n.left = n.left.Remove(removeVal)
		}
	} else if removeVal.Greater(n.val) {
		if n.right != nil {
			n.right = n.right.Remove(removeVal)
		}
	} else {
		if n.left == nil && n.right != nil {
			return n.right
		} else if n.right == nil && n.left != nil {
			return n.left
		} else if n.right == nil && n.left == nil {
			n.unsetVal()
			return n
		}
		min, _ := n.right.Min()
		n.val = min
		n.right = n.right.Remove(min)
	}
	return n
}

// Utility max/min function for integers
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
