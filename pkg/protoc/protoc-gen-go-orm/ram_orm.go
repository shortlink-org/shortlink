package main

import (
	"strconv"

	"google.golang.org/protobuf/compiler/protogen"
)

func generateRamFile(gen *protogen.Plugin, file *protogen.File) {
	for _, message := range file.Messages {
		filename := message.GoIdent.GoName + ".ram.orm.go"

		// if we have filter and after checking the filter map, we don't have the message name, skip generating ORM
		if *filter != "" {
			if _, ok := filterMap[message.GoIdent.GoName]; !ok {
				return
			}
		}

		g := gen.NewGeneratedFile(filename, file.GoImportPath)
		printHeader(gen, g, file)

		g.P("import (")
		g.P("\"reflect\"")
		g.P("\"strings\"")
		g.P(")")
		g.P()

		// Generate the BuildFilter method
		structName := "Filter" + message.GoIdent.GoName

		generateBuildRamFilterMethod(g, structName, message.Fields)
	}
}

func generateBuildRamFilterMethod(g *protogen.GeneratedFile, structName string, fields []*protogen.Field) {
	g.P("func (f *", structName, ") BuildRAMFilter(item any) bool {")
	g.P("var fieldVal *StringFilterInput")
	g.P("var ok bool")
	g.P("var v reflect.Value")
	g.P()

	for _, field := range fields {
		if field.Desc.IsList() || field.Desc.IsMap() {
			continue
		}
		fieldName := field.GoName
		generateBuildRamFieldFilterConditions(g, fieldName)
	}

	g.P("return true")
	g.P("}")
}

func generateBuildRamFieldFilterConditions(g *protogen.GeneratedFile, fieldName string) {
	g.P("fieldVal, ok = reflect.ValueOf(f).Elem().FieldByName(", strconv.Quote(fieldName), ").Interface().(*StringFilterInput)")
	g.P("if !ok || fieldVal == nil { return true } // If field is not found or nil, no filtering is applied")
	g.P("v = reflect.ValueOf(item).Elem().FieldByName(", strconv.Quote(fieldName), ")")
	g.P("if !v.IsValid() { return false } // If the link does not have this field, fail the filter")

	// Generate condition checks
	g.P("if fieldVal.Eq != \"\" && v.String() != fieldVal.Eq { return false }")
	g.P("if fieldVal.Ne != \"\" && v.String() == fieldVal.Ne { return false }")
	g.P("if fieldVal.Lt != \"\" && !(v.String() < fieldVal.Lt) { return false }")
	g.P("if fieldVal.Le != \"\" && !(v.String() <= fieldVal.Le) { return false }")
	g.P("if fieldVal.Gt != \"\" && !(v.String() > fieldVal.Gt) { return false }")
	g.P("if fieldVal.Ge != \"\" && !(v.String() >= fieldVal.Ge) { return false }")
	g.P("if fieldVal.StartsWith != \"\" && !strings.HasPrefix(v.String(), fieldVal.StartsWith) { return false }")
	g.P("if fieldVal.EndsWith != \"\" && !strings.HasSuffix(v.String(), fieldVal.EndsWith) { return false }")

	// Handle Contains and NotContains for slice of strings
	g.P("for _, contain := range fieldVal.Contains {")
	g.P("    if !strings.Contains(v.String(), contain) { return false }")
	g.P("}")
	g.P("for _, notContain := range fieldVal.NotContains {")
	g.P("    if strings.Contains(v.String(), notContain) { return false }")
	g.P("}")
	g.P()

	// Handle IsEmpty and IsNotEmpty
	g.P("if fieldVal.IsEmpty && v.String() != \"\" { return false }")
	g.P("if fieldVal.IsNotEmpty && v.String() == \"\" { return false }")
}
