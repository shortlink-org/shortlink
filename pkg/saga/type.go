package saga

import (
	"context"
)

type StepState int

const (
	WAIT StepState = iota
	START
	RUN
	READY
	REJECT
	FAIL
)

type EventState int

type Step struct {
	Name   string
	Status StepState
	Do     func(ctx context.Context) error
	Reject func(ctx context.Context) error
}

type Store interface{}

type Saga struct {
	Ctx   context.Context
	Name  string
	Store Store
	Steps []Step
}
