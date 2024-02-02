package main

import (
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func generateBuildMongoFilterMethod(g *protogen.GeneratedFile, structName string, fields []*protogen.Field) {
	g.P("func (f *", structName, ") BuildMongoFilter() bson.M {")
	g.P("if f == nil {")
	g.P("return nil")
	g.P("}")
	g.P("filter := bson.M{}")

	for _, field := range fields {
		if field.Desc.IsList() || field.Desc.IsMap() {
			continue
		}
		fieldName := field.GoName
		bsonFieldName := strings.ToLower(fieldName) // Adjust as per your field naming conventions
		g.P("if f.", fieldName, " != nil {")
		g.P("fieldFilter := bson.M{}")

		// Logic for each condition in the StringFilterInput struct
		g.P("if f.", fieldName, ".Eq != \"\" {")
		g.P("fieldFilter[\"$eq\"] = f.", fieldName, ".Eq")
		g.P("}")
		g.P("if f.", fieldName, ".Ne != \"\" {")
		g.P("fieldFilter[\"$ne\"] = f.", fieldName, ".Ne")
		g.P("}")
		g.P("if f.", fieldName, ".Lt != \"\" {")
		g.P("fieldFilter[\"$lt\"] = f.", fieldName, ".Lt")
		g.P("}")
		g.P("if f.", fieldName, ".Le != \"\" {")
		g.P("fieldFilter[\"$lte\"] = f.", fieldName, ".Le")
		g.P("}")
		g.P("if f.", fieldName, ".Gt != \"\" {")
		g.P("fieldFilter[\"$gt\"] = f.", fieldName, ".Gt")
		g.P("}")
		g.P("if f.", fieldName, ".Ge != \"\" {")
		g.P("fieldFilter[\"$gte\"] = f.", fieldName, ".Ge")
		g.P("}")

		// Handle Contains as an array
		g.P("if len(f.", fieldName, ".Contains) > 0 {")
		g.P("fieldFilter[\"$in\"] = f.", fieldName, ".Contains")
		g.P("}")

		// Handle NotContains as an array
		g.P("if len(f.", fieldName, ".NotContains) > 0 {")
		g.P("fieldFilter[\"$nin\"] = f.", fieldName, ".NotContains")
		g.P("}")

		g.P("if len(fieldFilter) > 0 {")
		g.P("filter[\"", bsonFieldName, "\"] = fieldFilter")
		g.P("}")
		g.P("}")
	}

	g.P("return filter")
	g.P("}")
}
