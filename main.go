package main

import (
	"fmt"

	"github.com/DhanviDesai/datastructures/maxheap"
	"github.com/DhanviDesai/datastructures/queue"
)

func main() {
	m := maxheap.MaxHeap[int]{}
	m.InsertElement(1)
	m.InsertElement(2)
	m.InsertElement(3)
	m.InsertElement(4)
	m.InsertElement(5)
	for m.HasElements() {
		fmt.Printf("Extracted element or max element is %d\n", m.ExtractElement())
	}

	q := queue.Queue[int]{}
	q.Add(1)
	q.Add(2)
	q.Add(3)
	q.Add(4)
	q.Add(5)

	for q.HasElements() {
		fmt.Printf("Polled element from queue is %d\n", q.Poll())
	}

}
