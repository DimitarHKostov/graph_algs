package queue

import (
	"graph_algs/pkg/linked_list"
)

type Queue struct {
	linkedList linked_list.LinkedList
}

func NewQueue() *Queue {
	linkedList := linked_list.NewLinkedList(nil, nil, 0)

	return &Queue{linkedList: *linkedList}
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
