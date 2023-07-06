package main

import "fmt"

func main() {

	t := tree[int]{
		value: 20,
		left:  nil,
		right: nil,
	}
	t.Insert(22)
	t.Insert(8)
	t.Insert(4)
	t.Insert(12)
	t.Insert(10)
	t.Insert(14)

	result := t.postOrderSuccessor(t.left) // tree with root 10

	fmt.Println(result)

}
