package main

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var TodoAnalyzer = &analysis.Analyzer{
	Name: "todo",
	Doc:  "ensures setter methods in types ending with 'Object' are private",
	Run:  run,
}

func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{TodoAnalyzer}, nil
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if comment, ok := n.(*ast.Comment); ok {
				if strings.HasPrefix(comment.Text, "// TODO:") || strings.HasPrefix(comment.Text, "// TODO():") {
					pass.Report(analysis.Diagnostic{
						Pos:            comment.Pos(),
						End:            0,
						Category:       "todo",
						Message:        "TODO comment has no author",
						SuggestedFixes: nil,
					})
				}
			}

			return true
		})
	}

	return nil, nil
}
