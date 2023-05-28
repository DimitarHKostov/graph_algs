package graph

import (
	"math"
	"sort"
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

func (g *WeightedGraph) Prim() int {
	start := GetRandomGraphElement(g.nodes)
	mst := make(map[GraphElement]bool)
	mst[start] = true
	total := 0

	for {
		if len(mst) == len(g.nodes) {
			break
		}

		minWeight := math.MaxInt64
		var minWeightTo GraphElement

		for node := range mst {
			for neighbor, edgeWeight := range g.nodes[node] {
				if !mst[neighbor] && edgeWeight < minWeight {
					minWeight = edgeWeight
					minWeightTo = neighbor
				}
			}
		}

		mst[minWeightTo] = true
		total += minWeight
	}

	return total
}

func (g *WeightedGraph) getVertices(nodes map[GraphElement]map[GraphElement]int) []GraphElement {
	vertices := make([]GraphElement, 0)

	for node := range nodes {
		vertices = append(vertices, node)
	}

	return vertices
}

func (g *WeightedGraph) getEdges(nodes map[GraphElement]map[GraphElement]int) []GraphEdge {
	edges := make([]GraphEdge, 0)
	addedEdges := make(map[GraphElement]map[GraphElement]bool)

	for node := range nodes {
		addedEdges[node] = make(map[GraphElement]bool)
	}

	for node, to := range nodes {
		for neighbor, edgeWeight := range to {
			add := true

			ok1 := addedEdges[node][neighbor]
			ok2 := addedEdges[neighbor][node]

			if ok1 || ok2 {
				add = false
			}

			if add {
				addedEdges[node][neighbor] = true
				addedEdges[neighbor][node] = true
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

func (ds *DisjointSet) Union(node1, node2 GraphElement) {
	root1 := ds.Find(node1)
	root2 := ds.Find(node2)

	if root1 != root2 {
		if ds.rank[root1].Less(ds.rank[root2]) {
			ds.parent[root1] = root2
		} else if ds.rank[root2].Less(ds.rank[root1]) {
			ds.parent[root2] = root1
		} else {
			ds.parent[root2] = root1
			ds.rank[root1] = root1
		}
	}
}

func (g *WeightedGraph) Kruskal() int {
	edges := g.getEdges(g.nodes)
	vertices := g.getVertices(g.nodes)
	ds := NewDisjointSet()
	total := 0

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
			ds.Union(srcRoot, destRoot)
		}
	}

	return total
}
