package saga

import (
	"context"
)

type BuilderStep struct {
	*Step

	errorList []error
}

func (s *BuilderStep) Then(f func(ctx context.Context) error) *BuilderStep {
	s.then = f
	return s
}

func (s *BuilderStep) Reject(f func(ctx context.Context, thenError error) error) *BuilderStep {
	s.reject = f
	return s
}

func (s *BuilderStep) Needs(keys ...string) *BuilderStep {
	// set parents
	for _, key := range keys {
		errAddEdge := s.dag.AddEdge(key, s.name)
		if errAddEdge != nil {
			s.errorList = append(s.errorList, errAddEdge)
		}
	}

	return s
}

func (s *BuilderStep) Build() (*Step, []error) {
	return s.Step, s.errorList
}
