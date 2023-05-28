package graph

import (
	"math/rand"
	"time"
)

type GraphElement struct {
	Id int
}

func ProduceGraphElement(id int) *GraphElement {
	return &GraphElement{Id: id}
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

func (ge GraphElement) Equals(other GraphElement) bool {
	return ge.Id == other.Id
}

func (ge GraphElement) Less(other GraphElement) bool {
	return ge.Id < other.Id
}

func (ge GraphElement) Greater(other GraphElement) bool {
	return !ge.Less(other) && !ge.Equals(other)
}
