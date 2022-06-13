package index

type TreeIndex interface {
	int
}

// Index - interface of index
type Index interface {
	Find(key any) []byte
	Insert(T any) error
	Delete(key any) error

	Marshal() ([]byte, error)
	UnMarshal([]byte, interface{}) error
}
