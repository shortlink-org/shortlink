package saga

import (
	"context"

	"github.com/shortlink-org/shortlink/internal/pkg/saga/dag"
)

type Builder struct {
	*Saga

	errorList []error
}

func New(name string, setters ...Option) *Builder {
	newSaga := &Saga{
		name:  name,
		dag:   dag.New(),
		steps: make(map[string]*Step),
	}

	for _, setter := range setters {
		setter(&newSaga.Options)
	}

	return &Builder{
		Saga: newSaga,
	}
}

func (s *Builder) WithContext(ctx context.Context) *Builder {
	s.ctx = ctx
	return s
}

func (s *Builder) Build() (*Saga, []error) {
	return s.Saga, s.errorList
}
