package saga

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"golang.org/x/sync/errgroup"

	"github.com/shortlink-org/shortlink/pkg/pattern/saga/dag"
)

type Saga struct {
	ctx context.Context
	Options
	steps     map[string]*Step
	dag       *dag.Dag
	name      string
	errorList []error
}

func (s *Saga) AddStep(name string, setters ...Option) *BuilderStep {
	// create a new step
	step := &BuilderStep{
		Step: &Step{
			ctx:    s.ctx,
			name:   name,
			dag:    s.dag,
			status: INIT,

			Options: Options{log: s.log},
		},
	}

	for _, setter := range setters {
		setter(&step.Options)
	}

	// check uniq
	if s.steps[name] != nil {
		s.errorList = append(s.errorList, &DublicateStepError{Name: name})
	}

	s.steps[name] = step.Step

	// add vertex to DAG
	_, err := s.dag.AddVertex(name, nil)
	if err != nil {
		s.errorList = append(s.errorList, err)
	}

	return step
}

func (s *Saga) Play(initSteps map[string]*Step) error {
	var err error

	// Get steps for RUN
	if len(initSteps) == 0 {
		initSteps, err = s.getRootSteps()
		if err != nil {
			return err
		}
	}

	// Run root steps
	g := errgroup.Group{}
	// set limiter for goroutines if it's set
	if s.limiter > 0 {
		g.SetLimit(s.limiter)
	}

	// start tracing
	newCtx, span := otel.Tracer(fmt.Sprintf("saga: %s", s.name)).Start(s.ctx, fmt.Sprintf("saga: %s", s.name))
	defer span.End()

	span.SetAttributes(attribute.String("saga", s.name))

	for _, step := range initSteps {
		step.ctx = newCtx //nolint:fatcontext // false positive?
		g.Go(step.Run)
	}

	err = g.Wait()
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		errReject := s.Reject(initSteps)

		return errReject
	}

	// Run children
	initChildrenStep := make(map[string]*Step)

	for _, rootStep := range initSteps {
		vertex, errGetVertex := s.dag.GetVertex(rootStep.name)
		if errGetVertex != nil {
			span.RecordError(errGetVertex)
			span.SetStatus(codes.Error, errGetVertex.Error())

			return errGetVertex
		}

		for _, child := range vertex.Children() {
			step := s.steps[child.GetId()]
			initChildrenStep[step.name] = step
		}
	}

	if len(initChildrenStep) == 0 {
		return nil
	}

	initChildrenStep, err = s.validateRun(initChildrenStep)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return err
	}

	return s.Play(initChildrenStep)
}

func (s *Saga) getRootSteps() (map[string]*Step, error) {
	initSteps := make(map[string]*Step)

	for _, step := range s.steps {
		// get steps with status: INIT
		if step.status == INIT {
			// get core steps
			vertex, err := s.dag.GetVertex(step.name)
			if err != nil {
				return nil, err
			}

			if len(vertex.Parents()) == 0 {
				initSteps[step.name] = step
			}
		}
	}

	return initSteps, nil
}

func (s *Saga) validateRun(steps map[string]*Step) (map[string]*Step, error) {
	// skip if status of all parents steps not DONE
	doneSteps := make(map[string]*Step)

	for _, step := range steps {
		vertex, errGetVertex := s.dag.GetVertex(step.name)
		if errGetVertex != nil {
			return nil, errGetVertex
		}

		isDone := true

		for _, child := range vertex.Parents() {
			if s.steps[child.GetId()].status != DONE {
				isDone = false
				step.status = WAIT
			}
		}

		if isDone {
			doneSteps[step.name] = step
		}
	}

	return doneSteps, nil
}

func (s *Saga) Reject(rejectSteps map[string]*Step) error {
	// Run root steps
	g := errgroup.Group{}

	for _, step := range rejectSteps {
		g.Go(step.Reject)
	}

	// ignore error and continue reject parent func
	err := g.Wait()
	if err != nil {
		return err
	}

	// get parents
	initParentStep := make(map[string]*Step)

	for _, rootStep := range rejectSteps {
		vertex, errGetVertex := s.dag.GetVertex(rootStep.name)
		if errGetVertex != nil {
			return errGetVertex
		}

		for _, child := range vertex.Parents() {
			step := s.steps[child.GetId()]
			initParentStep[step.name] = step
		}
	}

	if len(initParentStep) == 0 {
		return nil
	}

	initParentStep, err = s.validateReject(initParentStep)
	if err != nil {
		return err
	}

	return s.Reject(initParentStep)
}

func (s *Saga) validateReject(steps map[string]*Step) (map[string]*Step, error) {
	// skip the status of all parents step ROLLBACK's
	doneSteps := make(map[string]*Step)

	for _, step := range steps {
		vertex, errGetVertex := s.dag.GetVertex(step.name)
		if errGetVertex != nil {
			return nil, errGetVertex
		}

		isDone := true

		for _, parent := range vertex.Parents() {
			if s.steps[parent.GetId()].status == ROLLBACK && step.status != DONE {
				isDone = false
			}
		}

		if isDone {
			doneSteps[step.name] = step
		}
	}

	return doneSteps, nil
}
