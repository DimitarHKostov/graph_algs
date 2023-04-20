package graph

import (
	"graph_algs/pkg/node"
	"log"
)

type Graph struct {
	Nodes map[node.Node]map[node.Node]bool
}

func (g *Graph) BFS() {
	log.Println("run bfs")
}
