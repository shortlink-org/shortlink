package saga

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"

	"github.com/batazor/shortlink/pkg/saga/dag"
)

type Saga struct {
	ctx   context.Context
	name  string
	store Store
	steps map[string]*Step
	dag   *dag.Dag

	errorList []error
}

func (s *Saga) AddStep(name string) *BuilderStep {
	// create a new step
	step := &BuilderStep{
		Step: &Step{
			ctx:    &s.ctx,
			name:   name,
			dag:    s.dag,
			status: INIT,
		},
	}

	// check uniq
	if s.steps[name] != nil {
		step.errorList = append(s.errorList, fmt.Errorf("Dublicate step: %s", name))
	}
	s.steps[name] = step.Step

	// add vertex to DAG
	_, err := s.dag.AddVertex(name, nil)
	if err != nil {
		s.errorList = append(s.errorList, err)
	}

	return step
}

func (s *Saga) Play(initSteps []*Step) error {
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

	for _, step := range initSteps {
		g.Go(step.Run)
	}

	err = g.Wait()
	if err != nil {
		return err
	}

	// Run children
	initChildrenStep := []*Step{}
	for _, rootStep := range initSteps {
		vertex, err := s.dag.GetVertex(rootStep.name)
		if err != nil {
			return err
		}

		for _, child := range vertex.Children() {
			initChildrenStep = append(initChildrenStep, s.steps[child.GetId()])
		}
	}
	if len(initChildrenStep) == 0 {
		return nil
	}

	initChildrenStep, err = s.validate(initChildrenStep)
	if err != nil {
		return err
	}

	fmt.Println("===========================")

	return s.Play(initChildrenStep)
}

func (s *Saga) getRootSteps() ([]*Step, error) {
	initSteps := []*Step{}

	for _, step := range s.steps {
		// get steps with status: INIT
		if step.status == INIT {
			// get core steps
			vertex, err := s.dag.GetVertex(step.name)
			if err != nil {
				return nil, err
			}

			if len(vertex.Parents()) == 0 {
				initSteps = append(initSteps, step)
			}
		}
	}

	return initSteps, nil
}

func (s *Saga) validate(steps []*Step) ([]*Step, error) {
	var err error

	// drop double
	steps, err = s.uniq(steps)
	if err != nil {
		return nil, err
	}

	// drop if status of all parents step not DONE
	doneSteps := []*Step{}
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
			doneSteps = append(doneSteps, step)
		}
	}

	// drop double
	doneSteps, err = s.uniq(doneSteps)
	if err != nil {
		return nil, err
	}

	return doneSteps, nil
}

func (s *Saga) uniq(steps []*Step) ([]*Step, error) {
	// drop double
	mapStep := map[string]*Step{}
	for _, step := range steps {
		mapStep[step.name] = step
	}

	uniqSteps := []*Step{}
	for _, v := range mapStep {
		uniqSteps = append(uniqSteps, v)
	}

	return uniqSteps, nil
}
