package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/compiler/protogen"
)

const (
	version         = "1.2.0"
	commonFilename  = "common_types.orm.go" // Name of the file where common types are defined
	defaultDatabase = "postgres"            // Default database type to generate ORM for (e.g. postgres, mongo)
)

var (
	// Optional parameter to generate ORM types, e.g. --orm=postgres
	generateFor string
)

func main() {
	log.Println("protoc-go-orm version", version, "is called")

	// The main function runs the plugin.
	protogen.Options{
		ParamFunc: func(name, value string) error {
			switch name {
			case "orm":
				generateFor = value
			default:
				log.Printf("unknown parameter %q", name)
			}
			return nil
		},
	}.Run(func(gen *protogen.Plugin) error {
		log.Println("Running with protoc version:", protocVersion(gen))

		if generateFor == "" {
			generateFor = defaultDatabase
		}
		log.Printf("Generating ORM for: %s", generateFor)

		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}

			generateCommonFile(gen, f)        // Generate the common types file
			generateFile(gen, f, generateFor) // Generate ORM
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

func generateFile(gen *protogen.Plugin, file *protogen.File, dbType string) {
	switch dbType {
	case "postgres":
		generatePostgresFile(gen, file)
	case "mongo":
		generateMongoFile(gen, file)
	default:
		log.Printf("Unsupported database type: %s", dbType)
	}
}
