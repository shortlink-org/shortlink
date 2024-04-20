package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	v1 "github.com/shortlink-org/shortlink/pkg/protoc/protoc-gen-rich-model/options/v1"
)

const (
	version = "1.3.0"
)

var (
	filter = flag.String("filter", "", "Specify the filter type for generating the rich model")

	filterMap = map[string]struct{}{}
)

func main() {
	log.Println("protoc-rich-model version:", version, "is called")

	flag.Parse()

	// The main function runs the plugin.
	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(gen *protogen.Plugin) error {
		log.Println("Running with protoc version:", protocVersion(gen))

		// Convert the filter string into a map for quick lookup
		if *filter != "" {
			for _, name := range strings.Split(*filter, ";") {
				filterMap[name] = struct{}{}
			}
		}

		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}

			for _, message := range f.Messages {
				generateRichModel(gen, f, message) // Generate rich model for each message
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

func generateRichModel(gen *protogen.Plugin, file *protogen.File, message *protogen.Message) {
	// Filter by the specified type
	if *filter != "" {
		if _, ok := filterMap[message.GoIdent.GoName]; !ok {
			return
		}
	}

	// Adjust the filename to follow the '<typeName>.ddd.go' pattern
	filename := fmt.Sprintf("%s_ddd.go", strings.ToLower(message.GoIdent.GoName))

	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	printHeader(gen, g, file)

	importManager := newImportManager()

	// Preprocess to collect necessary imports
	for _, field := range message.Fields {
		if field.GoName == "" {
			continue
		}
		_, usedImports := protobufToGoType(field)
		importManager.addImports(usedImports)
	}

	importManager.writeImports(g)

	// Add comment for the message if available
	if message.Comments.Leading != "" {
		g.P(strings.TrimSpace(message.Comments.Leading.String()))
	}

	// Generate a rich model struct
	structName := strings.ToUpper(message.GoIdent.GoName[:1]) + message.GoIdent.GoName[1:] // Capitalize the first letter
	g.P("type ", structName, " struct {")

	for _, field := range message.Fields {
		if field.GoName == "" {
			continue
		}

		goType, usedImports := protobufToGoType(field)
		if goType == "" {
			continue
		}

		importManager.addImports(usedImports)

		fieldName := strings.ToLower(field.GoName)

		// Add comment for the field if available
		if field.Comments.Leading != "" {
			g.P(strings.TrimSpace(field.Comments.Leading.String()))
		}

		g.P(fieldName, " ", goType)
	}

	g.P("}")

	// Generate getters for each field
	generateGetters(g, message, structName)
}

// Generate getters for each field in the message
func generateGetters(g *protogen.GeneratedFile, message *protogen.Message, structName string) {
	for _, field := range message.Fields {
		if field.GoName == "" {
			continue
		}
		goType, _ := protobufToGoType(field)
		if goType == "" {
			continue
		}

		fieldName := strings.ToLower(field.GoName)
		getterName := "Get" + field.GoName // Capitalize the first letter for the public method

		g.P()
		g.P("// ", getterName, " returns the value of the ", fieldName, " field.")
		g.P("func (m *", structName, ") ", getterName, "() ", goType, " {")
		g.P("return m.", fieldName)
		g.P("}")
	}
}

// importManager manages the imports for the generated file
func protobufToGoType(field *protogen.Field) (string, map[string]bool) {
	if field.Desc.IsList() {
		innerType, innerImports := protobufToGoTypeSingle(field)
		if innerType != "" {
			return "[]" + innerType, innerImports
		}

		return "", nil
	}

	return protobufToGoTypeSingle(field)
}

// protobufToGoTypeSingle converts a protobuf field to a Go type
func protobufToGoTypeSingle(field *protogen.Field) (string, map[string]bool) {
	imports := make(map[string]bool)

	if opts := field.Desc.Options(); opts != nil {
		if proto.HasExtension(opts, v1.E_GoType) {
			customType, ok := proto.GetExtension(opts, v1.E_GoType).(string)
			if !ok {
				log.Println("invalid GoType extension for field:", field.GoName)
				return "", nil
			}

			// Determine if imports are needed for the custom type and add them.
			if strings.Contains(customType, "url.URL") {
				imports["net/url"] = true
			}

			return customType, imports
		}
	}

	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		return "bool", nil
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Uint32Kind:
		return "int", nil
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Uint64Kind:
		return "int64", nil
	case protoreflect.StringKind:
		return "string", nil
	case protoreflect.BytesKind:
		return "[]byte", nil
	case protoreflect.EnumKind:
		return "int32", nil // Enums are represented as int32 in Go
	case protoreflect.FloatKind:
		return "float32", nil
	case protoreflect.DoubleKind:
		return "float64", nil
	case protoreflect.MessageKind, protoreflect.GroupKind:
		if field.Message != nil && field.Message.GoIdent.GoImportPath != "" {
			if field.Message.GoIdent.GoName == "Timestamp" {
				imports["google.golang.org/protobuf/types/known/timestamppb"] = true
				return "*timestamppb.Timestamp", imports
			} else if field.Message.GoIdent.GoName == "FieldMask" {
				imports["google.golang.org/protobuf/types/known/fieldmaskpb"] = true
				return "fieldmaskpb.FieldMask", imports
			}
		}

		return "*" + strings.ToUpper(field.Message.GoIdent.GoName[:1]) + field.Message.GoIdent.GoName[1:], nil // Pointer to the type, but lowercase
	default:
		return "", nil
	}
}

// printHeader prints the header of the generated file
func printHeader(gen *protogen.Plugin, g *protogen.GeneratedFile, file *protogen.File) {
	g.P("// Code generated by protoc-gen-rich-model. DO NOT EDIT.")
	g.P("// versions:")
	g.P("// - protoc-gen-rich-model v" + version)
	g.P("// - protoc             ", protocVersion(gen))
	g.P("// source: ", file.Desc.Path())
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
}
