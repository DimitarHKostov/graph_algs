package graph

type GraphEdge struct {
	source, destination GraphElement
	weight              int
}

func NewGraphEdge(source, destination GraphElement, weight int) *GraphEdge {
	return &GraphEdge{source: source, destination: destination, weight: weight}
}
