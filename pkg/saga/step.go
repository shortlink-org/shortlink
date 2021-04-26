package saga

import (
	"context"
	"fmt"

	"github.com/batazor/shortlink/internal/pkg/logger/field"
	"github.com/batazor/shortlink/pkg/saga/dag"
)

type Step struct {
	ctx    *context.Context
	name   string
	status StepState
	then   func(ctx context.Context) error
	reject func(ctx context.Context) error
	dag    *dag.Dag

	// options
	Options
}

func (s *Step) Run() error {
	s.logger.Info(fmt.Sprintf("Run step by name: %s", s.name), field.Fields{"name": s.name})
	s.status = RUN
	err := s.then(*s.ctx)
	if err != nil {
		s.status = REJECT
		return err
	}
	s.status = DONE

	return nil
}

func (s *Step) Reject() error {
	s.logger.Info(fmt.Sprintf("Reject step by name: %s", s.name), field.Fields{"name": s.name})
	s.status = REJECT
	err := s.reject(*s.ctx)
	if err != nil {
		s.status = FAIL
		return err
	}
	s.status = ROLLBACK

	return nil
}
