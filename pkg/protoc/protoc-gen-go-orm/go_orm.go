package main

import (
	"flag"
	"fmt"
	"log"

	"google.golang.org/protobuf/compiler/protogen"
)

const (
	version        = "1.2.0"
	commonFilename = "common_types.orm.go" // Name of the file where common types are defined
)

var (
	dbType      = flag.String("orm", "postgres", "Specify the ORM type (postgres, mongo)")
	packageName = flag.String("pkg", "", "Specify the Go package name for the generated files")
)

func main() {
	flag.Parse()

	log.Println("protoc-go-orm version", version, "is called")

	// The main function runs the plugin.
	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(gen *protogen.Plugin) error {
		log.Println("Running with protoc version:", protocVersion(gen))
		log.Printf("Generating ORM for: %s", *dbType)

		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}

			generateCommonFile(gen, f) // Generate the common types file

			// Generate ORM
			if err := generateFile(gen, f); err != nil {
				log.Fatal(err)
			}
		}

		return nil
	})
}

func protocVersion(gen *protogen.Plugin) string {
	v := gen.Request.GetCompilerVersion()
	if v == nil {
		return "(unknown)"
	}
	var suffix string
	if s := v.GetSuffix(); s != "" {
		suffix = "-" + s
	}

	return fmt.Sprintf("v%d.%d.%d%s", v.GetMajor(), v.GetMinor(), v.GetPatch(), suffix)
}

func generateFile(gen *protogen.Plugin, file *protogen.File) error {
	switch *dbType {
	case "postgres":
		generatePostgresFile(gen, file)
	case "mongo":
		generateMongoFile(gen, file)
	default:
		return NotSupportDatabaseError{DbType: *dbType}
	}

	return nil
}
