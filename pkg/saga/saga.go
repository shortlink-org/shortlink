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
	steps map[string]*Step

	// A Directed acyclic graph or DAG  describes the workflow processes and
	// how are they related each other.
	dag *dag.Dag

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
	initChildrenStep := make(map[string]*Step)
	for _, rootStep := range initSteps {
		vertex, errGetVertex := s.dag.GetVertex(rootStep.name)
		if errGetVertex != nil {
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
		return err
	}

	fmt.Println("===========================")

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
	// skip if status of all parents step not DONE
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
	// skip if status of all parents step ROLLBACK
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
