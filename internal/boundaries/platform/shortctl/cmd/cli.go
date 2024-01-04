/*
CLI tooling

- generate ENV-docs
*/
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/internal/boundaries/platform/shortctl/internal/tool"
)

func init() {
	viper.SetDefault("SERVICE_NAME", "shortlink-cli")

	rootCmd := &cobra.Command{
		Use:   "shortctl",
		Short: "Shortlink it's sandbox for experiments",
		Long:  "Demo microservice architecture and best practices",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	rootCmd.Flags().String("o", "./docs/env.md", "Output file path")
	if err := viper.BindPFlag("o", rootCmd.Flags().Lookup("o")); err != nil {
		log.Fatal(err)
	}

	rootCmd.Flags().String("include-dir", "cmd,internal,pkg", "Include directories")
	if err := viper.BindPFlag("include-dir", rootCmd.Flags().Lookup("include-dir")); err != nil {
		log.Fatal(err)
	}

	rootCmd.Flags().String("exclude-dir", "vendor,node_modules,dist,ui", "Exclude directories")
	if err := viper.BindPFlag("exclude-dir", rootCmd.Flags().Lookup("exclude-dir")); err != nil {
		log.Fatal(err)
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

	// Retrieve the output file path using Viper
	outputPath := viper.GetString("o")

	// Extract the directory from the output file path
	outputDir := filepath.Dir(outputPath)

	// Ensure the directory exists or create it
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
			log.Fatalf("Failed to create directory: %v", err)
		}
	}
}

func main() {
	err := displayBanner()
	if err != nil {
		log.Fatalf("Error displaying banner: %v", err)
	}

	config := Config{}
	filePath := viper.GetString("o")

	dirs := findDirectories()

	for _, dir := range dirs {
		pterm.DefaultSection.Printf(`Search in directory %s`, dir)
		config.setConfigDocs(dir, &config)
	}

	payload := config.renderMDTable(config)

	if err := tool.SaveToFile(filePath, payload); err != nil {
		fmt.Println(err)
		return
	}
}

func displayBanner() error {
	err := pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("Short", pterm.NewStyle(pterm.FgCyan)),
		putils.LettersFromStringWithStyle("Link", pterm.NewStyle(pterm.FgLightMagenta))).
		Render()

	if err != nil {
		return err
	}

	return nil
}

func (*Config) setConfigDocs(basePath string, config *Config) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, basePath, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	for _, pkg := range pkgs {
		for fileName, file := range pkg.Files {
			wg.Add(1)
			go func(fileName string, file *ast.File) {
				defer wg.Done()
				pterm.Success.Printf("working on file %v\n", fileName)
				processFile(fset, fileName, file, config, basePath)
			}(fileName, file)
		}
	}
	wg.Wait()
}

func processFile(fset *token.FileSet, fileName string, file *ast.File, config *Config, basePath string) {
	ast.Inspect(file, func(n ast.Node) bool {
		return inspectASTNode(n, fileName, config, basePath)
	})

	processComments(fset, file, config, fileName)
}

func inspectASTNode(n ast.Node, fileName string, config *Config, basePath string) bool {
	if stmt, ok := n.(*ast.ExprStmt); ok {
		ast.Inspect(stmt.X, func(n ast.Node) bool {
			return extractEnvVariables(n, fileName, config, basePath)
		})
		return true
	}
	return true
}

func extractEnvVariables(n ast.Node, fileName string, config *Config, basePath string) bool {
	callExpr, ok := n.(*ast.CallExpr)
	if !ok {
		return true
	}

	selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok {
		return true
	}

	ident, ok := selectorExpr.X.(*ast.Ident)
	if !ok || ident.Name != "viper" || selectorExpr.Sel.Name != "SetDefault" {
		return true
	}

	// Ensure there are enough arguments for processing
	if len(callExpr.Args) < 2 {
		return true
	}

	env := ENV{
		pos:         callExpr.Args[0].(*ast.BasicLit).Pos(),
		fileName:    fileName,
		key:         callExpr.Args[0].(*ast.BasicLit).Value,
		fromPackage: filepath.Join(basePath, strings.TrimPrefix(fileName, basePath)),
	}

	switch arg := callExpr.Args[1].(type) {
	case *ast.BasicLit:
		env.value = tool.TrimQuotes(arg.Value)
		env.kind = arg.Kind.String()
	case *ast.Ident:
		if arg.Obj != nil {
			switch variable := arg.Obj.Decl.(type) {
			case *ast.AssignStmt:
				if len(variable.Rhs) > 0 {
					if call, ok := variable.Rhs[0].(*ast.CallExpr); ok {
						str := []any{}
						for i := range call.Args {
							if basicLit, ok := call.Args[i].(*ast.BasicLit); ok {
								str = append(str, tool.TrimQuotes(basicLit.Value))
							}
						}
						if len(str) > 0 {
							env.value = fmt.Sprintf(str[0].(string), str[1:]...)
						}
					}
				}
			}
		} else {
			env.value = tool.TrimQuotes(arg.Name)
		}
	case *ast.SelectorExpr:
		if xIdent, ok := arg.X.(*ast.Ident); ok {
			env.value = tool.TrimQuotes(fmt.Sprintf("%s.%s", xIdent.Name, arg.Sel.Name))
		}
	}

	config.appendEnv(env)
	return true
}

func processComments(fset *token.FileSet, file *ast.File, config *Config, fileName string) {
	for _, comment := range file.Comments {
		for _, item := range comment.List {
			line := fset.Position(item.Pos()).Line
			associateCommentsWithEnv(fset, line, item, config, fileName)
		}
	}
}

func associateCommentsWithEnv(fset *token.FileSet, line int, item *ast.Comment, config *Config, fileName string) {
	for index, conf := range config.envs {
		currentLine := fset.Position(conf.pos).Line
		if line == currentLine && fileName == conf.fileName {
			config.envs[index].describe = strings.TrimSpace(strings.TrimPrefix(item.Text, "//")) // remove comment symbols and trim spaces
		}
	}
}

func (*Config) renderMDTable(conf Config) string {
	str := `<!---
File generated by cli. DO NOT EDIT.
-->

# ENVIRONMENT

|Name | Default Value | Description | From Package |
|---|---|---|---|
`

	for _, env := range conf.envs {
		str += fmt.Sprintf("| %s | %s | %s | %s |\n", env.key, env.value, env.describe, env.fromPackage)
	}

	return str
}

func findDirectories() []string {
	var dirs []string
	includeDirs := viper.GetString("include-dir")
	findDirs := strings.Split(includeDirs, ",")

	excludeDirs := viper.GetString("exclude-dir")
	skipDirs := strings.Split(excludeDirs, ",")

	dirChan := make(chan []string)
	var wg sync.WaitGroup

	for _, dir := range findDirs {
		wg.Add(1)
		go func(dir string) {
			defer wg.Done()
			resp, err := tool.GetDirectories(dir, skipDirs)
			if err != nil {
				fmt.Println(err)
				dirChan <- nil
				return
			}
			dirChan <- resp
		}(dir)
	}

	go func() {
		wg.Wait()
		close(dirChan)
	}()

	for d := range dirChan {
		if d != nil {
			dirs = append(dirs, d...)
		}
	}

	return dirs
}
