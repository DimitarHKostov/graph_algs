package graph

type GraphElement struct {
	Id int
}

func ProduceGraphElement(id int) *GraphElement {
	graphElement := &GraphElement{Id: id}

	return graphElement
}