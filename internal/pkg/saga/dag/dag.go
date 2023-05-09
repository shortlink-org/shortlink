// Directed Acyclic Graph implementation in golang.

package dag

import (
	"errors"
	"fmt"
	"sync"
)

type Dag struct {
	vertices sync.Map
}

// New - creates a new Directed Acyclic Graph
func New() *Dag {
	return &Dag{}
}

// AddVertex adds a vertex to the graph
func (d *Dag) AddVertex(id string, value interface{}) (*Vertex, error) {
	// check on uniq key
	if _, ok := d.vertices.Load(id); ok {
		return nil, fmt.Errorf("dag already contains a vertex with the id: %s", id)
	}

	node := &Vertex{
		id:       id,
		value:    value,
		parents:  make([]*Vertex, 0),
		children: make([]*Vertex, 0),
	}
	d.vertices.Store(id, node)

	return node, nil
}

func (d *Dag) AddEdge(from, to string) error {
	fromVertexRaw, ok := d.vertices.Load(from)
	if !ok {
		return fmt.Errorf("not found vertex by id: %s", from)
	}

	fromVertex, ok := fromVertexRaw.(*Vertex)
	if !ok {
		return errors.New("incorrect type assertion")
	}
	toVertexRaw, ok := d.vertices.Load(to)

	if !ok {
		return fmt.Errorf("not found vertex by id: %s", to)
	}

	toVertex, ok := toVertexRaw.(*Vertex)
	if !ok {
		return errors.New("incorrect type assertion")
	}

	fromVertex.children = append(fromVertex.children, toVertex)
	toVertex.parents = append(toVertex.parents, fromVertex)

	// update map
	d.vertices.Store(from, fromVertex)
	d.vertices.Store(to, toVertex)

	return nil
}

func (d *Dag) GetVertex(id string) (*Vertex, error) {
	vertex, ok := d.vertices.Load(id)
	if !ok {
		return nil, fmt.Errorf("not found vertex with name: %s", id)
	}

	return vertex.(*Vertex), nil
}
