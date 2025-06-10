package models

type queue[T any] struct {
	list []T
}

func NewQueue[T any]() *queue[T] {
	return &queue[T]{}
}

func (q *queue[T]) Enqueue(value T) {
	q.list = append(q.list, value)
}

func (q *queue[T]) Dequeue() T {
	var ZERO T
	if len(q.list) == 0 {
		return ZERO
	}
	temp := q.list[0]
	q.list = q.list[1:]
	return temp
}

func (q *queue[T]) Len() int {
	return len(q.list)
}
