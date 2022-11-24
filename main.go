package main

import (
	"fmt"

	"github.com/DhanviDesai/datastructures/binarysearchtree"
)

func main() {
	k := 1
	binarysearchtree.AddNode(3)
	binarysearchtree.AddNode(1)
	binarysearchtree.AddNode(2)
	binarysearchtree.AddNode(5)
	binarysearchtree.AddNode(4)
	fmt.Printf("kth smallest element is %d\n", binarysearchtree.InorderTraversal()[k-1])
	fmt.Printf("Number of nodes is %d\n", binarysearchtree.CountNodes())

}
