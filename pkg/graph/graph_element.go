package graph

import (
	"math/rand"
	"time"
)

type GraphElement struct {
	Id int
}

func ProduceGraphElement(id int) *GraphElement {
	graphElement := &GraphElement{Id: id}

	return graphElement
}

func GetRandomGraphElement(nodes map[GraphElement]map[GraphElement]int) GraphElement {
	keys := make([]GraphElement, 0, len(nodes))

	for key := range nodes {
		keys = append(keys, key)
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(keys))

	return keys[randomIndex]
}
