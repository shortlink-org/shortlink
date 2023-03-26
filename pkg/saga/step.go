package saga

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"

	"github.com/shortlink-org/shortlink/pkg/saga/dag"
)

type Step struct {
	Options
	ctx    *context.Context
	then   func(ctx context.Context) error
	reject func(ctx context.Context) error
	dag    *dag.Dag
	name   string
	status StepState
}

func (s *Step) Run() error {
	// start tracing
	newCtx, span := otel.Tracer(fmt.Sprintf("saga: %s", s.name)).Start(*s.ctx, fmt.Sprintf("saga: %s", s.name))
	span.SetAttributes(attribute.String("step", s.name))
	defer span.End()

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
	newCtx, span := otel.Tracer(fmt.Sprintf("saga: %s", s.name)).Start(*s.ctx, fmt.Sprintf("saga: %s", s.name))
	span.SetAttributes(attribute.String("step", s.name))
	defer span.End()

	s.status = REJECT

	// Check on compensation step
	if s.reject == nil {
		return nil
	}

	err := s.reject(newCtx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		s.status = FAIL
		return err
	}
	s.status = ROLLBACK

	return nil
}
