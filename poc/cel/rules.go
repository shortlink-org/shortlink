package main

import (
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	"github.com/shortlink-org/shortlink/pkg/fsroot"
)

func loadRules(path string) (map[string]string, error) {
	// Create a SafeFS rooted at the specified path to restrict file access
	fs, err := fsroot.NewSafeFS(path)
	if err != nil {
		return nil, err
	}
	defer fs.Close()

	// Read the directory contents using SafeFS
	files, err := fs.ReadDir(".")
	if err != nil {
		return nil, err
	}

	rules := make(map[string]string)
	for _, file := range files {
		if !file.IsDir() {
			// Read file content using SafeFS (no need for filepath.Join)
			content, errReadFile := fs.ReadFile(file.Name())
			if errReadFile != nil {
				return nil, errReadFile
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
