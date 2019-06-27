package main

import "fmt"

func main() {
	list := []int{1, 2, 0, -1, 0, 0, 0, 0, 0, 1000, 9999}

	root := &Node{}
	for _, v := range list {
		if err := root.Insert(v); err != nil {
			fmt.Println(err, v)
		} else {
			fmt.Println("added", v)
		}
	}

	fmt.Println("Height", root.Height())

	if l, err := root.Max(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Max", l)
	}

	if s, err := root.Min(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Min", s)
	}

	fmt.Println(root.InOrder())
}
