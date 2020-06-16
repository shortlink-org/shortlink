package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
)

type ENV struct {
	key   string
	value string
	kind  string
}

type Config struct {
	envs []ENV
}

var (
	config   = Config{}
	findDirs = []string{"cmd", "internal", "pkg"}
	skipDirs = []string{"vendor", "node_modules", "dist", "ui"}
)

func main() {
	var dirs []string

	for _, dir := range findDirs {
		resp, err := getDirectories(dir)
		if err != nil {
			fmt.Println(err)
			return
		}

		dirs = append(dirs, resp...)
	}

	for _, dir := range dirs {
		setConfigDocs(dir)
	}

	fmt.Println(config)
}

func getDirectories(root string) ([]string, error) {
	dirs := []string{}

	err := filepath.Walk(
		root,
		func(path string, f os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if f.IsDir() && isExist(skipDirs, f.Name()) {
				return filepath.SkipDir
			}

			if f.IsDir() {
				dirs = append(dirs, path)
			}

			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	return dirs, nil
}

func isExist(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func setConfigDocs(path string) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	for _, pkg := range pkgs {
		for fileName, file := range pkg.Files {
			fmt.Printf("working on file %v\n", fileName)
			ast.Inspect(file, func(n ast.Node) bool {
				if stmt, ok := n.(*ast.ExprStmt); ok {

					ast.Inspect(stmt.X, func(n ast.Node) bool {
						if x, ok := n.(*ast.CallExpr); ok {
							if fun, ok := x.Fun.(*ast.SelectorExpr); ok {
								if ident, ok := fun.X.(*ast.Ident); ok {
									if ident.Name == "viper" && fun.Sel.Name == "SetDefault" {
										env := ENV{
											key: x.Args[0].(*ast.BasicLit).Value,
										}

										switch arg := x.Args[1].(type) {
										case *ast.BasicLit:
											env.value = arg.Value
											env.kind = arg.Kind.String()
										case *ast.Ident:
											env.value = arg.Name
										case *ast.SelectorExpr:
											env.value = fmt.Sprintf("%s.%s", arg.X.(*ast.Ident).Name, arg.Sel.Name)
										}

										config.envs = append(config.envs, env)
									}
								}
							}
						}

						return true
					})

					return true
				}

				return true
			})
		}
	}
}
