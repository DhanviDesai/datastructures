package queue

import (
	"golang.org/x/exp/constraints"
)

type Queue[T constraints.Ordered] struct {
	array []T
}

func (q *Queue[T]) Add(value T) {
	q.array = append(q.array, value)
}

func (q *Queue[T]) Poll() T {
	element := q.array[0]
	temp := make([]T, len(q.array)-1)
	copy(temp, q.array[1:])
	q.array = temp
	return element
}

func (q *Queue[T]) HasElements() bool {
	return len(q.array) > 0
}
