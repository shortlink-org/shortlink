package saga

import (
	"context"

	"github.com/shortlink-org/shortlink/pkg/saga/dag"
)

type SagaBuilder struct {
	*Saga

	errorList []error
}

func New(name string, setters ...Option) *SagaBuilder {
	newSaga := &Saga{
		name:  name,
		dag:   dag.New(),
		steps: make(map[string]*Step),
	}

	for _, setter := range setters {
		setter(&newSaga.Options)
	}

	return &SagaBuilder{
		Saga: newSaga,
	}
}

func (s *SagaBuilder) WithContext(ctx context.Context) *SagaBuilder {
	s.ctx = ctx
	return s
}

func (s *SagaBuilder) Build() (*Saga, []error) {
	return s.Saga, s.errorList
}
