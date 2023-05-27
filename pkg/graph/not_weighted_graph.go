package graph

import (
	"graph_algs/pkg/queue"
	"log"
)

type NotWeightedGraph struct {
	nodes map[GraphElement]map[GraphElement]bool
}

func NewNotWeightedGraph() *NotWeightedGraph {
	graph := &NotWeightedGraph{nodes: make(map[GraphElement]map[GraphElement]bool)}

	return graph
}

func (g *NotWeightedGraph) AddEdge(a, b GraphElement) {
	if _, ok := g.nodes[a]; !ok {
		g.nodes[a] = make(map[GraphElement]bool)
	}

	if _, ok := g.nodes[b]; !ok {
		g.nodes[b] = make(map[GraphElement]bool)
	}

	g.nodes[a][b] = true
	g.nodes[b][a] = true
}

func (g *NotWeightedGraph) BFS(a GraphElement) {
	visited := make(map[GraphElement]bool)
	queue := queue.NewQueue()

	visited[a] = true
	queue.Push(a.Id)

	for {
		if queue.IsEmpty() {
			break
		}

		out := queue.Get()
		log.Println(out)

		for neighbor := range g.nodes[a] {
			if !visited[neighbor] {
				queue.Push(neighbor.Id)
				visited[neighbor] = true
			}
		}
	}
}

func (g *NotWeightedGraph) DFS(a GraphElement) {
	visited := make(map[GraphElement]bool)
	g.DFSRecursive(a, visited)
}

func (g *NotWeightedGraph) DFSRecursive(a GraphElement, visited map[GraphElement]bool) {
	visited[a] = true
	log.Println(a.Id)

	for neighbor := range g.nodes[a] {
		if !visited[neighbor] {
			g.DFSRecursive(neighbor, visited)
		}
	}
}
