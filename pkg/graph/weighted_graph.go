package graph

import (
	"math"
)

type WeightedGraph struct {
	nodes map[GraphElement]map[GraphElement]int
}

func NewWeightedGraph() *WeightedGraph {
	graph := &WeightedGraph{nodes: make(map[GraphElement]map[GraphElement]int)}

	return graph
}

func (g *WeightedGraph) AddEdge(a, b GraphElement, weight int) {
	if _, ok := g.nodes[a]; !ok {
		g.nodes[a] = make(map[GraphElement]int)
	}

	if _, ok := g.nodes[b]; !ok {
		g.nodes[b] = make(map[GraphElement]int)
	}

	g.nodes[a][b] = weight
	g.nodes[b][a] = weight
}

func (g *WeightedGraph) GetWeight(a, b GraphElement) int {
	return g.nodes[a][b]
}

//todo: optimize using priority queue at some point
func (g *WeightedGraph) Dijkstra(from, to GraphElement) int {
	distanceTo := make(map[GraphElement]int)
	unvisited := make(map[GraphElement]bool)

	for node := range g.nodes {
		unvisited[node] = true

		if node != from {
			distanceTo[node] = math.MaxInt64
		} else {
			distanceTo[node] = 0
		}
	}

	for {
		if len(unvisited) == 0 {
			break
		}

		minDist := math.MaxInt64
		var minDistNode GraphElement

		for node := range unvisited {
			if distanceTo[node] < minDist {
				minDist = distanceTo[node]
				minDistNode = node
			}
		}

		delete(unvisited, minDistNode)

		for neighbor := range g.nodes[minDistNode] {
			if unvisited[neighbor] {
				if newDist := distanceTo[minDistNode] + g.nodes[minDistNode][neighbor]; newDist < distanceTo[neighbor] {
					distanceTo[neighbor] = newDist
				}
			}
		}
	}

	return distanceTo[to]
}
