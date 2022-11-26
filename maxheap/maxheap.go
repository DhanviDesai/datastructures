package maxheap

import (
	"golang.org/x/exp/constraints"
)

type MaxHeap[T constraints.Ordered] struct {
	array []T
}

func (h *MaxHeap[T]) InsertElement(element T) {
	h.array = append(h.array, element)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap[T]) ExtractElement() T {
	element := h.array[0]
	h.array[0] = h.array[len(h.array)-1]
	temp := make([]T, len(h.array)-1)
	copy(temp, h.array[:len(h.array)-1])
	h.array = temp
	h.maxHeapifyDown(0)
	return element
}

func (h *MaxHeap[T]) maxHeapifyDown(index int) {
	var greaterChild int
	if h.hasLeftChild(index) {
		greaterChild = left(index)
		if h.hasRightChild(index) && h.array[right(index)] > h.array[greaterChild] {
			greaterChild = right(index)
		}
		if h.array[index] < h.array[greaterChild] {
			h.swap(index, greaterChild)
			index = greaterChild
			h.maxHeapifyDown(index)
		}
	}
}

func (h *MaxHeap[T]) maxHeapifyUp(index int) {
	for h.hasParent(index) && h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return 2*index + 1
}

func right(index int) int {
	return 2*index + 2
}

func (h *MaxHeap[T]) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func (h *MaxHeap[T]) hasLeftChild(index int) bool {
	return left(index) < len(h.array)
}

func (h *MaxHeap[T]) hasRightChild(index int) bool {
	return right(index) < len(h.array)
}

func (h *MaxHeap[T]) hasParent(index int) bool {
	return parent(index) >= 0
}

func (h *MaxHeap[T]) HasElements() bool {
	return len(h.array) > 0
}
