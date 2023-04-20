package main

import (
	"graph_algs/pkg/linked_list"
	"graph_algs/pkg/node"
	"log"
)

func getNodes() map[node.Node]map[node.Node]bool {
	// todo refactor this
	nodeOne := &node.Node{Id: 1}
	nodeTwo := &node.Node{Id: 2}
	nodeThree := &node.Node{Id: 3}
	nodeFour := &node.Node{Id: 4}
	nodeFive := &node.Node{Id: 5}

	// refactor this later
	nodes := make(map[node.Node]map[node.Node]bool)
	nodes[*nodeOne] = make(map[node.Node]bool)
	nodes[*nodeTwo] = make(map[node.Node]bool)
	nodes[*nodeThree] = make(map[node.Node]bool)
	nodes[*nodeFour] = make(map[node.Node]bool)
	nodes[*nodeFive] = make(map[node.Node]bool)

	nodes[*nodeOne][*nodeTwo] = true
	nodes[*nodeOne][*nodeThree] = true
	nodes[*nodeTwo][*nodeFour] = true
	nodes[*nodeFour][*nodeFive] = true

	return nodes
}

func main() {
	// graph := graph.Graph{Nodes: getNodes()}

	// graph.BFS()

	ll := &linked_list.LinkedList{}
	ll.AddFront(1)

	log.Println(ll.GetFront().Id)
}
