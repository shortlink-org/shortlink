package mongo

import (
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/batazor/shortlink/internal/store/query"
)

func getFilter(filter *query.Filter) bson.D {
	filterQuery := bson.D{}
	r := reflect.ValueOf(filter)

	for _, key := range filter.GetKeys() {
		val := reflect.Indirect(r).FieldByName(key).Interface().(*query.StringFilterInput)

		// Skip empty value
		if val == nil {
			continue
		}

		// Eq
		if val.Eq != nil {
			filterQuery = append(filterQuery, primitive.E{Key: strings.ToLower(key), Value: *val.Eq})
		}

		// Ne
		if val.Ne != nil {
			filterQuery = append(filterQuery, primitive.E{
				Key:   strings.ToLower(key),
				Value: bson.D{{"$ne", *val.Ne}},
			})
		}

		// Lt
		if val.Lt != nil {
			filterQuery = append(filterQuery, primitive.E{
				Key:   strings.ToLower(key),
				Value: bson.D{{"$lt", *val.Lt}},
			})
		}

		// Lte
		if val.Le != nil {
			filterQuery = append(filterQuery, primitive.E{
				Key:   strings.ToLower(key),
				Value: bson.D{{"$lte", *val.Le}},
			})
		}

		// Gt
		if val.Gt != nil {
			filterQuery = append(filterQuery, primitive.E{
				Key:   strings.ToLower(key),
				Value: bson.D{{"$gt", *val.Gt}},
			})
		}

		// Ge
		if val.Ge != nil {
			filterQuery = append(filterQuery, primitive.E{
				Key:   strings.ToLower(key),
				Value: bson.D{{"$gte", *val.Ge}},
			})
		}

		// Contains
		if val.Contains != nil {
			filterQuery = append(filterQuery, primitive.E{
				Key:   strings.ToLower(key),
				Value: bson.D{{"$in", *val.Contains}},
			})
		}

		// NotContains
		if val.NotContains != nil {
			filterQuery = append(filterQuery, primitive.E{
				Key:   strings.ToLower(key),
				Value: bson.D{{"$nin", *val.NotContains}},
			})
		}
	}

	return filterQuery
}
