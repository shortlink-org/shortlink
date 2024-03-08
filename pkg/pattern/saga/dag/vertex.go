package dag

// Vertex type implements a vertex of a Directed Acyclic graph or DAG.
type Vertex struct {
	id       string
	value    any
	parents  []*Vertex
	children []*Vertex
}

func (v *Vertex) GetId() string {
	return v.id
}

func (v *Vertex) Parents() []*Vertex {
	return v.parents
}

func (v *Vertex) Children() []*Vertex {
	return v.children
}
