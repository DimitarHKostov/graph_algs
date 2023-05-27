package queue

import (
	"graph_algs/pkg/linked_list"
)

const (
	emptyQueueError = "queue is empty"
)

type Queue struct {
	linkedList linked_list.LinkedList
}

func NewQueue() *Queue {
	queue := &Queue{linkedList: linked_list.LinkedList{}}

	return queue
}

func (q *Queue) Push(element int) {
	q.linkedList.AddFront(&element)
}

func (q *Queue) IsEmpty() bool {
	return q.linkedList.GetSize() == 0
}

func (q *Queue) Get() int {
	element := q.linkedList.GetEnd().Id
	q.linkedList.PopEnd()
	return element
}
