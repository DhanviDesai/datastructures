package linkedlist

import (
	"fmt"
)

var head *Node

func AddNode(key int) {
	newNode := createNode(key)
	if head == nil {
		head = &newNode
	} else {
		temp := head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = &newNode
	}
}

func createNode(key int) Node {
	return Node{
		key:  key,
		next: nil,
	}
}

func TraverseList() {
	temp := head
	for temp != nil {
		fmt.Printf("%d\t", temp.key)
		temp = temp.next
	}
	fmt.Println()
}
