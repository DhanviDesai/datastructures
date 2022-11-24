package queue

var capacity = 10

type Queue[T any] struct {
	QueueList []T
	size      int
}

var queue = make([]int, capacity)

var queue2 = Queue[int]{
	QueueList: queue,
	size:      capacity,
}

var size = 0
var start = -1
var end = -1

func Add(value int) {
	queue = append(queue, value)
	size++
	end++
}

func Poll() int {
	start++
	size--
	return queue[start]
}
