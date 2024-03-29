// Code generated by protoc-gen-go-orm. DO NOT EDIT.
// versions:
// - protoc-gen-go-orm v1.1.0
// - protoc             (unknown)
// source: domain/sitemap/v1/sitemap.proto

package v1

import (
	"strings"
	"github.com/Masterminds/squirrel"
	"go.mongodb.org/mongo-driver/bson"
)

type FilterUrl struct {
	Loc        *StringFilterInput `json:"loc"`
	LastMod    *StringFilterInput `json:"lastmod"`
	ChangeFreq *StringFilterInput `json:"changefreq"`
	Priority   *StringFilterInput `json:"priority"`
}

func (f *FilterUrl) BuildFilter(query squirrel.SelectBuilder) squirrel.SelectBuilder {
	if f.Loc != nil {
		if f.Loc.Eq != "" {
			query = query.Where("loc = ?", f.Loc.Eq)
		}
		if f.Loc.Ne != "" {
			query = query.Where("loc <> ?", f.Loc.Ne)
		}
		if f.Loc.Lt != "" {
			query = query.Where("loc < ?", f.Loc.Lt)
		}
		if f.Loc.Le != "" {
			query = query.Where("loc <= ?", f.Loc.Le)
		}
		if f.Loc.Gt != "" {
			query = query.Where("loc > ?", f.Loc.Gt)
		}
		if f.Loc.Ge != "" {
			query = query.Where("loc >= ?", f.Loc.Ge)
		}
		if f.Loc.StartsWith != "" {
			query = query.Where("loc LIKE '%' || ?", f.Loc.StartsWith)
		}
		if f.Loc.EndsWith != "" {
			query = query.Where("loc LIKE ? || '%'", f.Loc.EndsWith)
		}
		if len(f.Loc.Contains) > 0 {
			containsQueries := []string{}
			containsArgs := []interface{}{}
			for _, v := range f.Loc.Contains {
				if v != "" {
					containsQueries = append(containsQueries, "loc LIKE ?")
					containsArgs = append(containsArgs, "%"+v+"%")
				}
			}
			if len(containsQueries) > 0 {
				query = query.Where("("+strings.Join(containsQueries, " OR ")+")", containsArgs...)
			}
		}
		if len(f.Loc.NotContains) > 0 {
			notContainsQueries := []string{}
			notContainsArgs := []interface{}{}
			for _, v := range f.Loc.NotContains {
				if v != "" {
					notContainsQueries = append(notContainsQueries, "loc NOT LIKE ?")
					notContainsArgs = append(notContainsArgs, "%"+v+"%")
				}
			}
			if len(notContainsQueries) > 0 {
				query = query.Where("("+strings.Join(notContainsQueries, " OR ")+")", notContainsArgs...)
			}
		}
		if f.Loc.IsEmpty {
			query = query.Where("loc = '' OR loc IS NULL")
		}
		if f.Loc.IsNotEmpty {
			query = query.Where("loc <> '' AND loc IS NOT NULL")
		}
	}
	if f.LastMod != nil {
		if f.LastMod.Eq != "" {
			query = query.Where("lastmod = ?", f.LastMod.Eq)
		}
		if f.LastMod.Ne != "" {
			query = query.Where("lastmod <> ?", f.LastMod.Ne)
		}
		if f.LastMod.Lt != "" {
			query = query.Where("lastmod < ?", f.LastMod.Lt)
		}
		if f.LastMod.Le != "" {
			query = query.Where("lastmod <= ?", f.LastMod.Le)
		}
		if f.LastMod.Gt != "" {
			query = query.Where("lastmod > ?", f.LastMod.Gt)
		}
		if f.LastMod.Ge != "" {
			query = query.Where("lastmod >= ?", f.LastMod.Ge)
		}
		if f.LastMod.StartsWith != "" {
			query = query.Where("lastmod LIKE '%' || ?", f.LastMod.StartsWith)
		}
		if f.LastMod.EndsWith != "" {
			query = query.Where("lastmod LIKE ? || '%'", f.LastMod.EndsWith)
		}
		if len(f.LastMod.Contains) > 0 {
			containsQueries := []string{}
			containsArgs := []interface{}{}
			for _, v := range f.LastMod.Contains {
				if v != "" {
					containsQueries = append(containsQueries, "lastmod LIKE ?")
					containsArgs = append(containsArgs, "%"+v+"%")
				}
			}
			if len(containsQueries) > 0 {
				query = query.Where("("+strings.Join(containsQueries, " OR ")+")", containsArgs...)
			}
		}
		if len(f.LastMod.NotContains) > 0 {
			notContainsQueries := []string{}
			notContainsArgs := []interface{}{}
			for _, v := range f.LastMod.NotContains {
				if v != "" {
					notContainsQueries = append(notContainsQueries, "lastmod NOT LIKE ?")
					notContainsArgs = append(notContainsArgs, "%"+v+"%")
				}
			}
			if len(notContainsQueries) > 0 {
				query = query.Where("("+strings.Join(notContainsQueries, " OR ")+")", notContainsArgs...)
			}
		}
		if f.LastMod.IsEmpty {
			query = query.Where("lastmod = '' OR lastmod IS NULL")
		}
		if f.LastMod.IsNotEmpty {
			query = query.Where("lastmod <> '' AND lastmod IS NOT NULL")
		}
	}
	if f.ChangeFreq != nil {
		if f.ChangeFreq.Eq != "" {
			query = query.Where("changefreq = ?", f.ChangeFreq.Eq)
		}
		if f.ChangeFreq.Ne != "" {
			query = query.Where("changefreq <> ?", f.ChangeFreq.Ne)
		}
		if f.ChangeFreq.Lt != "" {
			query = query.Where("changefreq < ?", f.ChangeFreq.Lt)
		}
		if f.ChangeFreq.Le != "" {
			query = query.Where("changefreq <= ?", f.ChangeFreq.Le)
		}
		if f.ChangeFreq.Gt != "" {
			query = query.Where("changefreq > ?", f.ChangeFreq.Gt)
		}
		if f.ChangeFreq.Ge != "" {
			query = query.Where("changefreq >= ?", f.ChangeFreq.Ge)
		}
		if f.ChangeFreq.StartsWith != "" {
			query = query.Where("changefreq LIKE '%' || ?", f.ChangeFreq.StartsWith)
		}
		if f.ChangeFreq.EndsWith != "" {
			query = query.Where("changefreq LIKE ? || '%'", f.ChangeFreq.EndsWith)
		}
		if len(f.ChangeFreq.Contains) > 0 {
			containsQueries := []string{}
			containsArgs := []interface{}{}
			for _, v := range f.ChangeFreq.Contains {
				if v != "" {
					containsQueries = append(containsQueries, "changefreq LIKE ?")
					containsArgs = append(containsArgs, "%"+v+"%")
				}
			}
			if len(containsQueries) > 0 {
				query = query.Where("("+strings.Join(containsQueries, " OR ")+")", containsArgs...)
			}
		}
		if len(f.ChangeFreq.NotContains) > 0 {
			notContainsQueries := []string{}
			notContainsArgs := []interface{}{}
			for _, v := range f.ChangeFreq.NotContains {
				if v != "" {
					notContainsQueries = append(notContainsQueries, "changefreq NOT LIKE ?")
					notContainsArgs = append(notContainsArgs, "%"+v+"%")
				}
			}
			if len(notContainsQueries) > 0 {
				query = query.Where("("+strings.Join(notContainsQueries, " OR ")+")", notContainsArgs...)
			}
		}
		if f.ChangeFreq.IsEmpty {
			query = query.Where("changefreq = '' OR changefreq IS NULL")
		}
		if f.ChangeFreq.IsNotEmpty {
			query = query.Where("changefreq <> '' AND changefreq IS NOT NULL")
		}
	}
	if f.Priority != nil {
		if f.Priority.Eq != "" {
			query = query.Where("priority = ?", f.Priority.Eq)
		}
		if f.Priority.Ne != "" {
			query = query.Where("priority <> ?", f.Priority.Ne)
		}
		if f.Priority.Lt != "" {
			query = query.Where("priority < ?", f.Priority.Lt)
		}
		if f.Priority.Le != "" {
			query = query.Where("priority <= ?", f.Priority.Le)
		}
		if f.Priority.Gt != "" {
			query = query.Where("priority > ?", f.Priority.Gt)
		}
		if f.Priority.Ge != "" {
			query = query.Where("priority >= ?", f.Priority.Ge)
		}
		if f.Priority.StartsWith != "" {
			query = query.Where("priority LIKE '%' || ?", f.Priority.StartsWith)
		}
		if f.Priority.EndsWith != "" {
			query = query.Where("priority LIKE ? || '%'", f.Priority.EndsWith)
		}
		if len(f.Priority.Contains) > 0 {
			containsQueries := []string{}
			containsArgs := []interface{}{}
			for _, v := range f.Priority.Contains {
				if v != "" {
					containsQueries = append(containsQueries, "priority LIKE ?")
					containsArgs = append(containsArgs, "%"+v+"%")
				}
			}
			if len(containsQueries) > 0 {
				query = query.Where("("+strings.Join(containsQueries, " OR ")+")", containsArgs...)
			}
		}
		if len(f.Priority.NotContains) > 0 {
			notContainsQueries := []string{}
			notContainsArgs := []interface{}{}
			for _, v := range f.Priority.NotContains {
				if v != "" {
					notContainsQueries = append(notContainsQueries, "priority NOT LIKE ?")
					notContainsArgs = append(notContainsArgs, "%"+v+"%")
				}
			}
			if len(notContainsQueries) > 0 {
				query = query.Where("("+strings.Join(notContainsQueries, " OR ")+")", notContainsArgs...)
			}
		}
		if f.Priority.IsEmpty {
			query = query.Where("priority = '' OR priority IS NULL")
		}
		if f.Priority.IsNotEmpty {
			query = query.Where("priority <> '' AND priority IS NOT NULL")
		}
	}
	return query
}
func (f *FilterUrl) BuildMongoFilter() bson.M {
	if f == nil {
		return nil
	}
	filter := bson.M{}
	if f.Loc != nil {
		fieldFilter := bson.M{}
		if f.Loc.Eq != "" {
			fieldFilter["$eq"] = f.Loc.Eq
		}
		if f.Loc.Ne != "" {
			fieldFilter["$ne"] = f.Loc.Ne
		}
		if f.Loc.Lt != "" {
			fieldFilter["$lt"] = f.Loc.Lt
		}
		if f.Loc.Le != "" {
			fieldFilter["$lte"] = f.Loc.Le
		}
		if f.Loc.Gt != "" {
			fieldFilter["$gt"] = f.Loc.Gt
		}
		if f.Loc.Ge != "" {
			fieldFilter["$gte"] = f.Loc.Ge
		}
		if len(f.Loc.Contains) > 0 {
			fieldFilter["$in"] = f.Loc.Contains
		}
		if len(f.Loc.NotContains) > 0 {
			fieldFilter["$nin"] = f.Loc.NotContains
		}
		if len(fieldFilter) > 0 {
			filter["loc"] = fieldFilter
		}
	}
	if f.LastMod != nil {
		fieldFilter := bson.M{}
		if f.LastMod.Eq != "" {
			fieldFilter["$eq"] = f.LastMod.Eq
		}
		if f.LastMod.Ne != "" {
			fieldFilter["$ne"] = f.LastMod.Ne
		}
		if f.LastMod.Lt != "" {
			fieldFilter["$lt"] = f.LastMod.Lt
		}
		if f.LastMod.Le != "" {
			fieldFilter["$lte"] = f.LastMod.Le
		}
		if f.LastMod.Gt != "" {
			fieldFilter["$gt"] = f.LastMod.Gt
		}
		if f.LastMod.Ge != "" {
			fieldFilter["$gte"] = f.LastMod.Ge
		}
		if len(f.LastMod.Contains) > 0 {
			fieldFilter["$in"] = f.LastMod.Contains
		}
		if len(f.LastMod.NotContains) > 0 {
			fieldFilter["$nin"] = f.LastMod.NotContains
		}
		if len(fieldFilter) > 0 {
			filter["lastmod"] = fieldFilter
		}
	}
	if f.ChangeFreq != nil {
		fieldFilter := bson.M{}
		if f.ChangeFreq.Eq != "" {
			fieldFilter["$eq"] = f.ChangeFreq.Eq
		}
		if f.ChangeFreq.Ne != "" {
			fieldFilter["$ne"] = f.ChangeFreq.Ne
		}
		if f.ChangeFreq.Lt != "" {
			fieldFilter["$lt"] = f.ChangeFreq.Lt
		}
		if f.ChangeFreq.Le != "" {
			fieldFilter["$lte"] = f.ChangeFreq.Le
		}
		if f.ChangeFreq.Gt != "" {
			fieldFilter["$gt"] = f.ChangeFreq.Gt
		}
		if f.ChangeFreq.Ge != "" {
			fieldFilter["$gte"] = f.ChangeFreq.Ge
		}
		if len(f.ChangeFreq.Contains) > 0 {
			fieldFilter["$in"] = f.ChangeFreq.Contains
		}
		if len(f.ChangeFreq.NotContains) > 0 {
			fieldFilter["$nin"] = f.ChangeFreq.NotContains
		}
		if len(fieldFilter) > 0 {
			filter["changefreq"] = fieldFilter
		}
	}
	if f.Priority != nil {
		fieldFilter := bson.M{}
		if f.Priority.Eq != "" {
			fieldFilter["$eq"] = f.Priority.Eq
		}
		if f.Priority.Ne != "" {
			fieldFilter["$ne"] = f.Priority.Ne
		}
		if f.Priority.Lt != "" {
			fieldFilter["$lt"] = f.Priority.Lt
		}
		if f.Priority.Le != "" {
			fieldFilter["$lte"] = f.Priority.Le
		}
		if f.Priority.Gt != "" {
			fieldFilter["$gt"] = f.Priority.Gt
		}
		if f.Priority.Ge != "" {
			fieldFilter["$gte"] = f.Priority.Ge
		}
		if len(f.Priority.Contains) > 0 {
			fieldFilter["$in"] = f.Priority.Contains
		}
		if len(f.Priority.NotContains) > 0 {
			fieldFilter["$nin"] = f.Priority.NotContains
		}
		if len(fieldFilter) > 0 {
			filter["priority"] = fieldFilter
		}
	}
	return filter
}

type FilterSitemap struct {
}

func (f *FilterSitemap) BuildFilter(query squirrel.SelectBuilder) squirrel.SelectBuilder {
	return query
}
func (f *FilterSitemap) BuildMongoFilter() bson.M {
	if f == nil {
		return nil
	}
	filter := bson.M{}
	return filter
}
