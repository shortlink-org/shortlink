package main

import (
	"go/token"
	"sync"
)

type ENV struct {
	key         string
	value       string
	kind        string
	describe    string
	fileName    string
	pos         token.Pos
	fromPackage string
}

type Config struct {
	mu   sync.Mutex
	envs []ENV
}

func (c *Config) appendEnv(env ENV) {
	c.mu.Lock()
	c.envs = append(c.envs, env)
	c.mu.Unlock()
}
