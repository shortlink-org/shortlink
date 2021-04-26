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
		// If get error run rejectFunc
		fmt.Printf("Run REJECT after run step with error: %s\n", err)
		errReject := s.Reject(initSteps)
		return errReject
	}

	// Run children
	initChildrenStep := []*Step{}
	for _, rootStep := range initSteps {
		vertex, errGetVertex := s.dag.GetVertex(rootStep.name)
		if errGetVertex != nil {
			return errGetVertex
		}

		for _, child := range vertex.Children() {
			initChildrenStep = append(initChildrenStep, s.steps[child.GetId()])
		}
	}
	if len(initChildrenStep) == 0 {
		return nil
	}

	initChildrenStep, err = s.validateRun(initChildrenStep)
	if err != nil {
		return err
	}

	fmt.Println("===========================")

	return s.Play(initChildrenStep)
}

func (s *Saga) Reject(rejectSteps []*Step) error {
	fmt.Println("===========================")
	fmt.Println("Run REJECT")

	// Run root steps
	g := errgroup.Group{}

	for _, step := range rejectSteps {
		g.Go(step.Reject)
	}

	// ignore error and continue reject parent func
	err := g.Wait()
	fmt.Println(err)

	// get parents
	initParentStep := []*Step{}
	for _, rootStep := range rejectSteps {
		vertex, errGetVertex := s.dag.GetVertex(rootStep.name)
		if errGetVertex != nil {
			return errGetVertex
		}

		for _, child := range vertex.Parents() {
			initParentStep = append(initParentStep, s.steps[child.GetId()])
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

func (s *Saga) validateRun(steps []*Step) ([]*Step, error) {
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

func (s *Saga) validateReject(steps []*Step) ([]*Step, error) {
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
		for _, parent := range vertex.Parents() {
			if s.steps[parent.GetId()].status == ROLLBACK && step.status != DONE {
				isDone = false
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
