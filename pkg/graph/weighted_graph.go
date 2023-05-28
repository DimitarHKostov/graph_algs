package graph

import (
	"math"
	"sort"
)

type WeightedGraph struct {
	nodes map[GraphElement]map[GraphElement]int
}

func NewWeightedGraph() *WeightedGraph {
	return &WeightedGraph{nodes: make(map[GraphElement]map[GraphElement]int)}
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

	for i := range g.nodes {
		for j := range g.nodes {
			for k := range g.nodes {
				if newDist := distanceBetween[i][k] + distanceBetween[k][j]; newDist < distanceBetween[i][j] {
					distanceBetween[i][j] = newDist
				}
			}
		}
	}

	return distanceBetween[from][to]
}

func (g *WeightedGraph) Prim() (int, []GraphEdge) {
	start := GetRandomGraphElement(g.nodes)
	mstVertices := make(map[GraphElement]bool)
	mstVertices[start] = true
	total := 0
	mstEdges := make([]GraphEdge, 0)

	for {
		if len(mstVertices) == len(g.nodes) {
			break
		}

		minWeight := math.MaxInt64
		var minWeightTo GraphElement
		var minWeightFrom GraphElement

		for node := range mstVertices {
			for neighbor, edgeWeight := range g.nodes[node] {
				if !mstVertices[neighbor] && edgeWeight < minWeight {
					minWeight = edgeWeight
					minWeightTo = neighbor
					minWeightFrom = node
				}
			}
		}

		mstVertices[minWeightTo] = true
		total += minWeight
		mstEdge := NewGraphEdge(minWeightFrom, minWeightTo, minWeight)
		mstEdges = append(mstEdges, *mstEdge)
	}

	return total, mstEdges
}

func (g *WeightedGraph) getVertices(nodes map[GraphElement]map[GraphElement]int) []GraphElement {
	vertices := make([]GraphElement, 0)
	verticesSet := make(map[GraphElement]bool)

	for node := range nodes {
		if !verticesSet[node] {
			verticesSet[node] = true
			vertices = append(vertices, node)
		}

	}

	return vertices
}

func (g *WeightedGraph) getEdges(nodes map[GraphElement]map[GraphElement]int) []GraphEdge {
	edges := make([]GraphEdge, 0)
	edgesSet := make(map[GraphElement]map[GraphElement]bool)

	for node := range nodes {
		edgesSet[node] = make(map[GraphElement]bool)
	}

	for node, to := range nodes {
		for neighbor, edgeWeight := range to {
			add := true

			ok1 := edgesSet[node][neighbor]
			ok2 := edgesSet[neighbor][node]

			if ok1 || ok2 {
				add = false
			}

			if add {
				edgesSet[node][neighbor] = true
				edgesSet[neighbor][node] = true
				edge := NewGraphEdge(node, neighbor, edgeWeight)
				edges = append(edges, *edge)
			}
		}
	}

	return edges
}

type DisjointSet struct {
	parent, rank map[GraphElement]GraphElement
}

func NewDisjointSet() *DisjointSet {
	return &DisjointSet{parent: make(map[GraphElement]GraphElement), rank: make(map[GraphElement]GraphElement)}
}

func (ds *DisjointSet) MakeSet(node GraphElement) {
	ds.parent[node] = node
	ds.rank[node] = node
}

func (ds *DisjointSet) Find(node GraphElement) GraphElement {
	if ds.parent[node] != node {
		ds.parent[node] = ds.Find(ds.parent[node])
	}
	return ds.parent[node]
}

func (ds *DisjointSet) Union(first, second GraphElement) {
	rootFirst := ds.Find(first)
	rootSecond := ds.Find(second)

	if rootFirst != rootSecond {
		if ds.rank[rootFirst].Less(ds.rank[rootSecond]) {
			ds.parent[rootFirst] = rootSecond
		} else if ds.rank[rootSecond].Less(ds.rank[rootFirst]) {
			ds.parent[rootSecond] = rootFirst
		} else {
			ds.parent[rootSecond] = rootFirst
			ds.rank[rootFirst] = rootFirst
		}
	}
}

func (g *WeightedGraph) Kruskal() (int, []GraphEdge) {
	edges := g.getEdges(g.nodes)
	vertices := g.getVertices(g.nodes)
	ds := NewDisjointSet()
	total := 0
	mstEdges := make([]GraphEdge, 0)

	for _, vertex := range vertices {
		ds.MakeSet(vertex)
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	for _, edge := range edges {
		srcRoot := ds.Find(edge.source)
		destRoot := ds.Find(edge.destination)

		if srcRoot != destRoot {
			total += edge.weight
			mstEdges = append(mstEdges, edge)
			ds.Union(srcRoot, destRoot)
		}
	}

	return total, mstEdges
}
