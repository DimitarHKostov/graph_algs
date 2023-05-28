package linked_list

import (
	"graph_algs/pkg/node"
	"log"
)

type LinkedList struct {
	head *node.Node
	end  *node.Node
	size int
}

func NewLinkedList(head, end *node.Node, size int) *LinkedList {
	return &LinkedList{head: head, end: end, size: size}
}

func (ll *LinkedList) AddFront(elementId *int) {
	if ll.size == 0 {
		ll.head = &node.Node{Id: *elementId}
		ll.end = ll.head
	} else {
		newNode := &node.Node{Id: *elementId}
		newNode.Next = ll.head
		ll.head.Prev = newNode
		ll.head = newNode
	}

	ll.size++
}

func (ll *LinkedList) AddEnd(element *int) {
	if ll.size == 0 {
		ll.head = &node.Node{Id: *element}
		ll.end = ll.head
	} else {
		newNode := &node.Node{Id: *element}
		newNode.Prev = ll.end
		ll.end.Next = newNode
		ll.end = newNode
	}

	ll.size++
}

func (ll *LinkedList) GetSize() int {
	return int(ll.size)
}

func (ll *LinkedList) GetFront() *node.Node {
	return ll.head
}

func (ll *LinkedList) GetEnd() *node.Node {
	return ll.end
}

func (ll *LinkedList) PopFront() {
	if ll.size > 0 {
		ll.size--
		ll.head = ll.head.Next
	}
}
func (ll *LinkedList) PopEnd() {
	if ll.size > 0 {
		ll.size--
		ll.end = ll.end.Prev
	}
}

func (ll *LinkedList) Print() {
	ll.printRecursive(ll.head)
}

func (ll *LinkedList) printRecursive(curr *node.Node) {
	if curr == nil {
		return
	}

	log.Println(curr.Id)
	ll.printRecursive(curr.Next)
}
