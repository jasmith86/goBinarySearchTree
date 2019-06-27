package main

import "fmt"

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

	fmt.Println(t.head.InOrder())
}
