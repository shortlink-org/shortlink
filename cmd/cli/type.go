package main

import "go/token"

type ENV struct {
	key      string
	value    string
	kind     string
	describe string

	// for match comments
	pos      token.Pos
	fileName string
}

type Config struct {
	envs []ENV
}
