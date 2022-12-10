package raft

import (
	"time"
)

type Raft interface {
	Connect(config Config) error
}

type Vote struct {
	Timeout  time.Duration
	NextVote time.Time
}

type Config struct {
	Name   string
	Weight int
	URI    string

	Vote Vote

	Status string
}
