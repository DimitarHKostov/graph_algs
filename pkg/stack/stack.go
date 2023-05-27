package stack

import (
	"graph_algs/pkg/linked_list"
)

const (
	emptyStackMessage = "stack is empty"
)

type Stack struct {
	linkedList linked_list.LinkedList
}

func NewStack() *Stack {
	stack := &Stack{linkedList: linked_list.LinkedList{}}

	return stack
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
