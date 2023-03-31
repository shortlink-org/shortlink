package main

import (
	"io/ioutil"
	"path/filepath"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
)

func loadRules(path string) (map[string]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	rules := make(map[string]string)
	for _, file := range files {
		if !file.IsDir() {
			filename := filepath.Join(path, file.Name())
			content, err := ioutil.ReadFile(filename)
			if err != nil {
				return nil, err
			}
			rules[file.Name()] = string(content)
		}
	}
	return rules, nil
}

func compileRules(rules map[string]string) (map[string]*cel.Program, error) {
	env, err := cel.NewEnv(
		cel.Declarations(
			decls.NewVar(
				"claims",
				decls.NewMapType(decls.String, decls.Dyn)),
			decls.NewVar("expected_audience", decls.String),
			decls.NewVar("now", decls.Timestamp),
			decls.NewVar("aud", decls.String),
		),
	)
	if err != nil {
		return nil, err
	}

	compiledRules := make(map[string]*cel.Program)
	for name, rule := range rules {
		ast := compile(env, rule, cel.BoolType)
		program, errProgram := env.Program(ast)
		if errProgram != nil {
			return nil, errProgram
		}
		compiledRules[name] = &program
	}

	return compiledRules, nil
}
