// Directed Acyclic Graph implementation in golang.

package dag

import (
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
func (d *Dag) AddVertex(id string, value any) (*Vertex, error) {
	// check on an uniq key
	if _, ok := d.vertices.Load(id); ok {
		return nil, &VertexAlreadyExistsError{Id: id}
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
		return &VertexNotFoundError{Id: from}
	}

	fromVertex, ok := fromVertexRaw.(*Vertex)
	if !ok {
		return ErrIncorrectTypeAssertion
	}

	toVertexRaw, ok := d.vertices.Load(to)

	if !ok {
		return &VertexNotFoundError{Id: to}
	}

	toVertex, ok := toVertexRaw.(*Vertex)
	if !ok {
		return ErrIncorrectTypeAssertion
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
		return nil, &VertexNotFoundError{Id: id}
	}

	if v, okAsserion := vertex.(*Vertex); okAsserion {
		return v, nil
	}

	return nil, ErrIncorrectTypeAssertion
}
