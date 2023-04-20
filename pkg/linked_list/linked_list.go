package linked_list

import "graph_algs/pkg/node"

type LinkedList struct {
	head *node.Node
	end  *node.Node
	size uint32
}

func (ll *LinkedList) AddFront(element int) {
	if ll.size == 0 {
		ll.addFirstElement(element)
		return
	}
}

func (ll *LinkedList) AddEnd(element int) {
	if ll.size == 0 {
		ll.addFirstElement(element)
		return
	}
}

func (ll *LinkedList) addFirstElement(element int) {
	ll.head = &node.Node{Id: element}
	ll.end = ll.head
	ll.size++
}

func (ll *LinkedList) GetFront() node.Node {
	return *ll.head
}

func (ll *LinkedList) GetEnd() node.Node {
	return *ll.end
}

func (ll *LinkedList) PopFront() {

}
func (ll *LinkedList) PopEnd() {

}
