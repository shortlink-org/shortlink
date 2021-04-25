package saga

import (
	"context"

	"github.com/batazor/shortlink/pkg/saga/dag"
)

type SagaBuilder struct {
	*Saga

	errorList []error
}

func New(name string) *SagaBuilder {
	return &SagaBuilder{
		Saga: &Saga{
			name:  name,
			dag:   dag.New(),
			steps: make(map[string]*Step),
		},
	}
}

func (s *SagaBuilder) WithContext(ctx context.Context) *SagaBuilder {
	s.ctx = ctx
	return s
}

func (s *SagaBuilder) SetStore(store Store) *SagaBuilder {
	s.store = store
	return s
}

func (s *SagaBuilder) Build() (*Saga, []error) {
	return s.Saga, s.errorList
}
