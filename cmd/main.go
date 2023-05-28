package main

import (
	"graph_algs/pkg/graph"
	"log"
)

var (
	one   = *graph.ProduceGraphElement(1)
	two   = *graph.ProduceGraphElement(2)
	three = *graph.ProduceGraphElement(3)
	four  = *graph.ProduceGraphElement(4)
	five  = *graph.ProduceGraphElement(5)
)

func notWeigtedGraphScenarios() {
	g := graph.NewNotWeightedGraph()

	g.AddEdge(one, two)
	g.AddEdge(one, three)

	g.BFS(one)
	g.DFS(one)
}

func weightedGraphScenarios() {
	wg := graph.NewWeightedGraph()

	wg.AddEdge(one, two, 2)
	wg.AddEdge(one, three, 1)
	wg.AddEdge(two, three, 4)
	wg.AddEdge(three, four, 2)
	wg.AddEdge(three, five, 3)
	wg.AddEdge(four, five, 5)

	//log.Println(wg.Dijkstra(one, five))
	//log.Println(wg.BellmanFord(one, five))
	//log.Println(wg.FloydWarshall(one, five))
	log.Println(wg.Prim())
	log.Println(wg.Kruskal())
}

func main() {
	//notWeigtedGraphScenarios()
	weightedGraphScenarios()
}
