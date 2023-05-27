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

// todo: optimize using priority queue at some point
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

func (g *WeightedGraph) BellmanFord(from, to GraphElement) int {
	distanceTo := make(map[GraphElement]int)

	for node := range g.nodes {
		if node != from {
			distanceTo[node] = math.MaxInt64
		} else {
			distanceTo[node] = 0
		}
	}

	for node := range g.nodes {
		for neighbor := range g.nodes[node] {
			if newDist := distanceTo[node] + g.nodes[node][neighbor]; newDist < distanceTo[neighbor] {
				distanceTo[neighbor] = newDist
			}
		}
	}

	//pass through all edges once more and if any gets relaxed => negative cycle

	return distanceTo[to]
}

func (g *WeightedGraph) FloydWarshall(from, to GraphElement) int {
	distanceBetween := make(map[GraphElement]map[GraphElement]int)

	for node := range g.nodes {
		if _, ok := distanceBetween[node]; !ok {
			distanceBetween[node] = make(map[GraphElement]int)
		}
	}

	for u := range distanceBetween {
		distanceBetween[u][u] = 0
	}

	for node := range g.nodes {
		for neighbor, edgeWeight := range g.nodes[node] {
			distanceBetween[node][neighbor] = edgeWeight
		}
	}

	for i := range g.nodes {
		for j := range g.nodes {
			if i != j && distanceBetween[i][j] == 0 {
				distanceBetween[i][j] = math.MaxInt64
				distanceBetween[j][i] = math.MaxInt64
			}
		}
	}

	for i := range distanceBetween {
		for j := range distanceBetween {
			for k := range distanceBetween {
				if newDist := distanceBetween[i][k] + distanceBetween[k][j]; newDist < distanceBetween[i][j] {
					distanceBetween[i][j] = newDist
				}
			}
		}
	}

	return distanceBetween[from][to]
}
