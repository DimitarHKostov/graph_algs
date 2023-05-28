package stack

import (
	"graph_algs/pkg/linked_list"
)

type Stack struct {
	linkedList linked_list.LinkedList
}

func NewStack() *Stack {
	linkedList := linked_list.NewLinkedList(nil, nil, 0)

	return &Stack{linkedList: *linkedList}
}

func (s *Stack) IsEmpty() bool {
	return s.linkedList.GetEnd().Id == 0
}

func (s *Stack) Add(element int) {
	s.linkedList.AddEnd(&element)
}

func (s *Stack) Get() int {
	element := s.linkedList.GetEnd()
	s.linkedList.PopEnd()
	return element.Id
}
