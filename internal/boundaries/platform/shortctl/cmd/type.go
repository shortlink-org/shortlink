package main

import "go/token"

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
	envs []ENV
}
