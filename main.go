package main

import "fmt"

func main() {
	list := []int{1, 2, 0, -1, 0, 0, 0, 0, 0, 1000, 9999}

	root := &Node{}

	cnt, err := root.InsertBulk(list)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("added", cnt, "values to", root)
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
	fmt.Println("Count", root.Count())
	fmt.Println("InOrder", root.InOrder())

}
