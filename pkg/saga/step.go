package saga

import (
	"context"
	"fmt"

	"github.com/batazor/shortlink/pkg/saga/dag"
)

type Step struct {
	ctx    *context.Context
	name   string
	status StepState
	then   func(ctx context.Context) error
	reject func(ctx context.Context) error
	dag    *dag.Dag
}

func (s *Step) Run() error {
	fmt.Printf("Run step by name: %s\n", s.name)
	s.status = RUN
	err := s.then(*s.ctx)
	if err != nil {
		s.status = REJECT
		return err
	}
	s.status = DONE

	return nil
}
