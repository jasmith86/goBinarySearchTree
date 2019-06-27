package main

import (
	"errors"
	"fmt"
)

// Tree interface TODO?

// Binary Search Tree
type BST struct {
	head  *Node
	count int
}

// Node of Binary Search Tree
type Node struct {
	val   int
	left  *Node
	right *Node
}

// Add new Node to tree. Returns number of nodes added and error
func (b *BST) Insert(v int) (int, error) {
	if b.head == nil { // insert first/head Node
		b.head = &Node{val: v}
		b.count += 1
		return 1, nil
	}
	err := b.head.insert(v) // try to insert Node
	if err != nil {
		return 0, err
	}
	b.count += 1
	return 1, nil
}

// insert a new node starting at n.
// TODO this should be a method of BST, not node.
func (n *Node) insert(v int) error {
	if v == n.val {
		return errors.New("failed trying to insert duplicate value")
	}
	if v < n.val { // Go left
		if n.left == nil { // can insert value
			n.left = &Node{val: v}
			return nil
		}
		return n.left.insert(v)
	}
	// Go right. v must be > n.val
	if n.right == nil { // can insert value
		n.right = &Node{val: v}
		return nil
	}
	return n.right.insert(v)
}

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
func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (n *Node) Max() (int, error) {
	if n == nil {
		return 0, errors.New("cannot get max of empty tree")
	}
	if n.right != nil {
		return n.right.Max()
	}
	return n.val, errors.New("cannot get max of empty tree")
}

func (n *Node) Min() (int, error) {
	if n == nil {
		return 0, errors.New("cannot get min of empty tree")
	}
	if n.left != nil {
		return n.left.Min()
	}
	return n.val, nil
}

func main() {
	list := []int{1, 2, 0, -1, 0, 0, 0, 0, 0, 1000, 9999}

	t := &BST{}
	for _, v := range list {
		if _, err := t.Insert(v); err != nil {
			fmt.Println(err, v)
		} else {
			fmt.Println("added", v)
		}
	}

	fmt.Println("Height", t.head.Height())

	if l, err := t.head.Max(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Max", l)
	}

	if s, err := t.head.Min(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Min", s)
	}
}
