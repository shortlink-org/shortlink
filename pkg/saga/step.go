package saga

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"

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
	// start tracing
	span, newCtx := opentracing.StartSpanFromContext(*s.ctx, fmt.Sprintf("step: %s", s.name))
	span.SetTag("step", s.name)
	defer span.Finish()

	s.status = RUN
	err := s.then(newCtx)
	if err != nil {
		s.status = REJECT
		return err
	}
	s.status = DONE

	return nil
}

func (s *Step) Reject() error {
	// start tracing
	span, newCtx := opentracing.StartSpanFromContext(*s.ctx, fmt.Sprintf("step: %s", s.name))
	span.SetTag("step", s.name)
	defer span.Finish()

	s.status = REJECT

	// Check on compensation step
	if s.reject == nil {
		return nil
	}

	err := s.reject(newCtx)
	if err != nil {
		s.status = FAIL
		return err
	}
	s.status = ROLLBACK

	return nil
}
