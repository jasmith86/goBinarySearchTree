package main

import "fmt"

// Fulfil Item interface specified in bst.go
type MyInt int

func (a MyInt) Less(than Item) bool {
	return a < than.(MyInt)
}
func (a MyInt) Greater(than Item) bool {
	return a > than.(MyInt)
}
func (a MyInt) Equals(to Item) bool {
	return a == to.(MyInt)
}

func main() {
	var list []MyInt
	list = []MyInt{1, 2, 0, -1, 0, 0, 0, 0, 0, 1000, 9999}

	root := &Node{}

	//var cnt int
	//for _, v := range list {
	//	err := root.Insert(v)
	//	if err != nil {
	//		fmt.Println(err)
	//	} else {
	//		cnt++
	//	}
	//}
	cnt, err := root.InsertBulk(list)
	_ = err
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
