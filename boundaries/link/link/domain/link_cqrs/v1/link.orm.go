// Code generated by protoc-gen-go-orm. DO NOT EDIT.
// versions:
// - protoc-gen-go-orm v1.1.0
// - protoc             (unknown)
// source: domain/link_cqrs/v1/link.proto

package v1

import (
	"strings"
	"github.com/Masterminds/squirrel"
	"go.mongodb.org/mongo-driver/bson"
)

type FilterLinkView struct {
	FieldMask       *StringFilterInput `json:"fieldmask"`
	Url             *StringFilterInput `json:"url"`
	Hash            *StringFilterInput `json:"hash"`
	Describe        *StringFilterInput `json:"describe"`
	ImageUrl        *StringFilterInput `json:"imageurl"`
	MetaDescription *StringFilterInput `json:"metadescription"`
	MetaKeywords    *StringFilterInput `json:"metakeywords"`
	CreatedAt       *StringFilterInput `json:"createdat"`
	UpdatedAt       *StringFilterInput `json:"updatedat"`
}

func (f *FilterLinkView) BuildFilter(query squirrel.SelectBuilder) squirrel.SelectBuilder {
	if f.FieldMask != nil {
		if f.FieldMask.Eq != "" {
			query = query.Where("fieldmask = ?", f.FieldMask.Eq)
		}
		if f.FieldMask.Ne != "" {
			query = query.Where("fieldmask <> ?", f.FieldMask.Ne)
		}
		if f.FieldMask.Lt != "" {
			query = query.Where("fieldmask < ?", f.FieldMask.Lt)
		}
		if f.FieldMask.Le != "" {
			query = query.Where("fieldmask <= ?", f.FieldMask.Le)
		}
		if f.FieldMask.Gt != "" {
			query = query.Where("fieldmask > ?", f.FieldMask.Gt)
		}
		if f.FieldMask.Ge != "" {
			query = query.Where("fieldmask >= ?", f.FieldMask.Ge)
		}
		if f.FieldMask.StartsWith != "" {
			query = query.Where("fieldmask LIKE '%' || ?", f.FieldMask.StartsWith)
		}
		if f.FieldMask.EndsWith != "" {
			query = query.Where("fieldmask LIKE ? || '%'", f.FieldMask.EndsWith)
		}
		if len(f.FieldMask.Contains) > 0 {
			containsQueries := []string{}
			containsArgs := []interface{}{}
			for _, v := range f.FieldMask.Contains {
				if v != "" {
					containsQueries = append(containsQueries, "fieldmask LIKE ?")
					containsArgs = append(containsArgs, "%"+v+"%")
				}
			}
			if len(containsQueries) > 0 {
				query = query.Where("("+strings.Join(containsQueries, " OR ")+")", containsArgs...)
			}
		}
		if len(f.FieldMask.NotContains) > 0 {
			notContainsQueries := []string{}
			notContainsArgs := []interface{}{}
			for _, v := range f.FieldMask.NotContains {
				if v != "" {
					notContainsQueries = append(notContainsQueries, "fieldmask NOT LIKE ?")
					notContainsArgs = append(notContainsArgs, "%"+v+"%")
				}
			}
			if len(notContainsQueries) > 0 {
				query = query.Where("("+strings.Join(notContainsQueries, " OR ")+")", notContainsArgs...)
			}
		}
		if f.FieldMask.IsEmpty {
			query = query.Where("fieldmask = '' OR fieldmask IS NULL")
		}
		if f.FieldMask.IsNotEmpty {
			query = query.Where("fieldmask <> '' AND fieldmask IS NOT NULL")
		}
	}
	if f.Url != nil {
		if f.Url.Eq != "" {
			query = query.Where("url = ?", f.Url.Eq)
		}
		if f.Url.Ne != "" {
			query = query.Where("url <> ?", f.Url.Ne)
		}
		if f.Url.Lt != "" {
			query = query.Where("url < ?", f.Url.Lt)
		}
		if f.Url.Le != "" {
			query = query.Where("url <= ?", f.Url.Le)
		}
		if f.Url.Gt != "" {
			query = query.Where("url > ?", f.Url.Gt)
		}
		if f.Url.Ge != "" {
			query = query.Where("url >= ?", f.Url.Ge)
		}
		if f.Url.StartsWith != "" {
			query = query.Where("url LIKE '%' || ?", f.Url.StartsWith)
		}
		if f.Url.EndsWith != "" {
			query = query.Where("url LIKE ? || '%'", f.Url.EndsWith)
		}
		if len(f.Url.Contains) > 0 {
			containsQueries := []string{}
			containsArgs := []interface{}{}
			for _, v := range f.Url.Contains {
				if v != "" {
					containsQueries = append(containsQueries, "url LIKE ?")
					containsArgs = append(containsArgs, "%"+v+"%")
				}
			}
			if len(containsQueries) > 0 {
				query = query.Where("("+strings.Join(containsQueries, " OR ")+")", containsArgs...)
			}
		}
		if len(f.Url.NotContains) > 0 {
			notContainsQueries := []string{}
			notContainsArgs := []interface{}{}
			for _, v := range f.Url.NotContains {
				if v != "" {
					notContainsQueries = append(notContainsQueries, "url NOT LIKE ?")
					notContainsArgs = append(notContainsArgs, "%"+v+"%")
				}
			}
			if len(notContainsQueries) > 0 {
				query = query.Where("("+strings.Join(notContainsQueries, " OR ")+")", notContainsArgs...)
			}
		}
		if f.Url.IsEmpty {
			query = query.Where("url = '' OR url IS NULL")
		}
		if f.Url.IsNotEmpty {
			query = query.Where("url <> '' AND url IS NOT NULL")
		}
	}
	if f.Hash != nil {
		if f.Hash.Eq != "" {
			query = query.Where("hash = ?", f.Hash.Eq)
		}
		if f.Hash.Ne != "" {
			query = query.Where("hash <> ?", f.Hash.Ne)
		}
		if f.Hash.Lt != "" {
			query = query.Where("hash < ?", f.Hash.Lt)
		}
		if f.Hash.Le != "" {
			query = query.Where("hash <= ?", f.Hash.Le)
		}
		if f.Hash.Gt != "" {
			query = query.Where("hash > ?", f.Hash.Gt)
		}
		if f.Hash.Ge != "" {
			query = query.Where("hash >= ?", f.Hash.Ge)
		}
		if f.Hash.StartsWith != "" {
			query = query.Where("hash LIKE '%' || ?", f.Hash.StartsWith)
		}
		if f.Hash.EndsWith != "" {
			query = query.Where("hash LIKE ? || '%'", f.Hash.EndsWith)
		}
		if len(f.Hash.Contains) > 0 {
			containsQueries := []string{}
			containsArgs := []interface{}{}
			for _, v := range f.Hash.Contains {
				if v != "" {
					containsQueries = append(containsQueries, "hash LIKE ?")
					containsArgs = append(containsArgs, "%"+v+"%")
				}
			}
			if len(containsQueries) > 0 {
				query = query.Where("("+strings.Join(containsQueries, " OR ")+")", containsArgs...)
			}
		}
		if len(f.Hash.NotContains) > 0 {
			notContainsQueries := []string{}
			notContainsArgs := []interface{}{}
			for _, v := range f.Hash.NotContains {
				if v != "" {
					notContainsQueries = append(notContainsQueries, "hash NOT LIKE ?")
					notContainsArgs = append(notContainsArgs, "%"+v+"%")
				}
			}
			if len(notContainsQueries) > 0 {
				query = query.Where("("+strings.Join(notContainsQueries, " OR ")+")", notContainsArgs...)
			}
		}
		if f.Hash.IsEmpty {
			query = query.Where("hash = '' OR hash IS NULL")
		}
		if f.Hash.IsNotEmpty {
			query = query.Where("hash <> '' AND hash IS NOT NULL")
		}
	}
	if f.Describe != nil {
		if f.Describe.Eq != "" {
			query = query.Where("describe = ?", f.Describe.Eq)
		}
		if f.Describe.Ne != "" {
			query = query.Where("describe <> ?", f.Describe.Ne)
		}
		if f.Describe.Lt != "" {
			query = query.Where("describe < ?", f.Describe.Lt)
		}
		if f.Describe.Le != "" {
			query = query.Where("describe <= ?", f.Describe.Le)
		}
		if f.Describe.Gt != "" {
			query = query.Where("describe > ?", f.Describe.Gt)
		}
		if f.Describe.Ge != "" {
			query = query.Where("describe >= ?", f.Describe.Ge)
		}
		if f.Describe.StartsWith != "" {
			query = query.Where("describe LIKE '%' || ?", f.Describe.StartsWith)
		}
		if f.Describe.EndsWith != "" {
			query = query.Where("describe LIKE ? || '%'", f.Describe.EndsWith)
		}
		if len(f.Describe.Contains) > 0 {
			containsQueries := []string{}
			containsArgs := []interface{}{}
			for _, v := range f.Describe.Contains {
				if v != "" {
					containsQueries = append(containsQueries, "describe LIKE ?")
					containsArgs = append(containsArgs, "%"+v+"%")
				}
			}
			if len(containsQueries) > 0 {
				query = query.Where("("+strings.Join(containsQueries, " OR ")+")", containsArgs...)
			}
		}
		if len(f.Describe.NotContains) > 0 {
			notContainsQueries := []string{}
			notContainsArgs := []interface{}{}
			for _, v := range f.Describe.NotContains {
				if v != "" {
					notContainsQueries = append(notContainsQueries, "describe NOT LIKE ?")
					notContainsArgs = append(notContainsArgs, "%"+v+"%")
				}
			}
			if len(notContainsQueries) > 0 {
				query = query.Where("("+strings.Join(notContainsQueries, " OR ")+")", notContainsArgs...)
			}
		}
		if f.Describe.IsEmpty {
			query = query.Where("describe = '' OR describe IS NULL")
		}
		if f.Describe.IsNotEmpty {
			query = query.Where("describe <> '' AND describe IS NOT NULL")
		}
	}
	if f.ImageUrl != nil {
		if f.ImageUrl.Eq != "" {
			query = query.Where("imageurl = ?", f.ImageUrl.Eq)
		}
		if f.ImageUrl.Ne != "" {
			query = query.Where("imageurl <> ?", f.ImageUrl.Ne)
		}
		if f.ImageUrl.Lt != "" {
			query = query.Where("imageurl < ?", f.ImageUrl.Lt)
		}
		if f.ImageUrl.Le != "" {
			query = query.Where("imageurl <= ?", f.ImageUrl.Le)
		}
		if f.ImageUrl.Gt != "" {
			query = query.Where("imageurl > ?", f.ImageUrl.Gt)
		}
		if f.ImageUrl.Ge != "" {
			query = query.Where("imageurl >= ?", f.ImageUrl.Ge)
		}
		if f.ImageUrl.StartsWith != "" {
			query = query.Where("imageurl LIKE '%' || ?", f.ImageUrl.StartsWith)
		}
		if f.ImageUrl.EndsWith != "" {
			query = query.Where("imageurl LIKE ? || '%'", f.ImageUrl.EndsWith)
		}
		if len(f.ImageUrl.Contains) > 0 {
			containsQueries := []string{}
			containsArgs := []interface{}{}
			for _, v := range f.ImageUrl.Contains {
				if v != "" {
					containsQueries = append(containsQueries, "imageurl LIKE ?")
					containsArgs = append(containsArgs, "%"+v+"%")
				}
			}
			if len(containsQueries) > 0 {
				query = query.Where("("+strings.Join(containsQueries, " OR ")+")", containsArgs...)
			}
		}
		if len(f.ImageUrl.NotContains) > 0 {
			notContainsQueries := []string{}
			notContainsArgs := []interface{}{}
			for _, v := range f.ImageUrl.NotContains {
				if v != "" {
					notContainsQueries = append(notContainsQueries, "imageurl NOT LIKE ?")
					notContainsArgs = append(notContainsArgs, "%"+v+"%")
				}
			}
			if len(notContainsQueries) > 0 {
				query = query.Where("("+strings.Join(notContainsQueries, " OR ")+")", notContainsArgs...)
			}
		}
		if f.ImageUrl.IsEmpty {
			query = query.Where("imageurl = '' OR imageurl IS NULL")
		}
		if f.ImageUrl.IsNotEmpty {
			query = query.Where("imageurl <> '' AND imageurl IS NOT NULL")
		}
	}
	if f.MetaDescription != nil {
		if f.MetaDescription.Eq != "" {
			query = query.Where("metadescription = ?", f.MetaDescription.Eq)
		}
		if f.MetaDescription.Ne != "" {
			query = query.Where("metadescription <> ?", f.MetaDescription.Ne)
		}
		if f.MetaDescription.Lt != "" {
			query = query.Where("metadescription < ?", f.MetaDescription.Lt)
		}
		if f.MetaDescription.Le != "" {
			query = query.Where("metadescription <= ?", f.MetaDescription.Le)
		}
		if f.MetaDescription.Gt != "" {
			query = query.Where("metadescription > ?", f.MetaDescription.Gt)
		}
		if f.MetaDescription.Ge != "" {
			query = query.Where("metadescription >= ?", f.MetaDescription.Ge)
		}
		if f.MetaDescription.StartsWith != "" {
			query = query.Where("metadescription LIKE '%' || ?", f.MetaDescription.StartsWith)
		}
		if f.MetaDescription.EndsWith != "" {
			query = query.Where("metadescription LIKE ? || '%'", f.MetaDescription.EndsWith)
		}
		if len(f.MetaDescription.Contains) > 0 {
			containsQueries := []string{}
			containsArgs := []interface{}{}
			for _, v := range f.MetaDescription.Contains {
				if v != "" {
					containsQueries = append(containsQueries, "metadescription LIKE ?")
					containsArgs = append(containsArgs, "%"+v+"%")
				}
			}
			if len(containsQueries) > 0 {
				query = query.Where("("+strings.Join(containsQueries, " OR ")+")", containsArgs...)
			}
		}
		if len(f.MetaDescription.NotContains) > 0 {
			notContainsQueries := []string{}
			notContainsArgs := []interface{}{}
			for _, v := range f.MetaDescription.NotContains {
				if v != "" {
					notContainsQueries = append(notContainsQueries, "metadescription NOT LIKE ?")
					notContainsArgs = append(notContainsArgs, "%"+v+"%")
				}
			}
			if len(notContainsQueries) > 0 {
				query = query.Where("("+strings.Join(notContainsQueries, " OR ")+")", notContainsArgs...)
			}
		}
		if f.MetaDescription.IsEmpty {
			query = query.Where("metadescription = '' OR metadescription IS NULL")
		}
		if f.MetaDescription.IsNotEmpty {
			query = query.Where("metadescription <> '' AND metadescription IS NOT NULL")
		}
	}
	if f.MetaKeywords != nil {
		if f.MetaKeywords.Eq != "" {
			query = query.Where("metakeywords = ?", f.MetaKeywords.Eq)
		}
		if f.MetaKeywords.Ne != "" {
			query = query.Where("metakeywords <> ?", f.MetaKeywords.Ne)
		}
		if f.MetaKeywords.Lt != "" {
			query = query.Where("metakeywords < ?", f.MetaKeywords.Lt)
		}
		if f.MetaKeywords.Le != "" {
			query = query.Where("metakeywords <= ?", f.MetaKeywords.Le)
		}
		if f.MetaKeywords.Gt != "" {
			query = query.Where("metakeywords > ?", f.MetaKeywords.Gt)
		}
		if f.MetaKeywords.Ge != "" {
			query = query.Where("metakeywords >= ?", f.MetaKeywords.Ge)
		}
		if f.MetaKeywords.StartsWith != "" {
			query = query.Where("metakeywords LIKE '%' || ?", f.MetaKeywords.StartsWith)
		}
		if f.MetaKeywords.EndsWith != "" {
			query = query.Where("metakeywords LIKE ? || '%'", f.MetaKeywords.EndsWith)
		}
		if len(f.MetaKeywords.Contains) > 0 {
			containsQueries := []string{}
			containsArgs := []interface{}{}
			for _, v := range f.MetaKeywords.Contains {
				if v != "" {
					containsQueries = append(containsQueries, "metakeywords LIKE ?")
					containsArgs = append(containsArgs, "%"+v+"%")
				}
			}
			if len(containsQueries) > 0 {
				query = query.Where("("+strings.Join(containsQueries, " OR ")+")", containsArgs...)
			}
		}
		if len(f.MetaKeywords.NotContains) > 0 {
			notContainsQueries := []string{}
			notContainsArgs := []interface{}{}
			for _, v := range f.MetaKeywords.NotContains {
				if v != "" {
					notContainsQueries = append(notContainsQueries, "metakeywords NOT LIKE ?")
					notContainsArgs = append(notContainsArgs, "%"+v+"%")
				}
			}
			if len(notContainsQueries) > 0 {
				query = query.Where("("+strings.Join(notContainsQueries, " OR ")+")", notContainsArgs...)
			}
		}
		if f.MetaKeywords.IsEmpty {
			query = query.Where("metakeywords = '' OR metakeywords IS NULL")
		}
		if f.MetaKeywords.IsNotEmpty {
			query = query.Where("metakeywords <> '' AND metakeywords IS NOT NULL")
		}
	}
	if f.CreatedAt != nil {
		if f.CreatedAt.Eq != "" {
			query = query.Where("createdat = ?", f.CreatedAt.Eq)
		}
		if f.CreatedAt.Ne != "" {
			query = query.Where("createdat <> ?", f.CreatedAt.Ne)
		}
		if f.CreatedAt.Lt != "" {
			query = query.Where("createdat < ?", f.CreatedAt.Lt)
		}
		if f.CreatedAt.Le != "" {
			query = query.Where("createdat <= ?", f.CreatedAt.Le)
		}
		if f.CreatedAt.Gt != "" {
			query = query.Where("createdat > ?", f.CreatedAt.Gt)
		}
		if f.CreatedAt.Ge != "" {
			query = query.Where("createdat >= ?", f.CreatedAt.Ge)
		}
		if f.CreatedAt.StartsWith != "" {
			query = query.Where("createdat LIKE '%' || ?", f.CreatedAt.StartsWith)
		}
		if f.CreatedAt.EndsWith != "" {
			query = query.Where("createdat LIKE ? || '%'", f.CreatedAt.EndsWith)
		}
		if len(f.CreatedAt.Contains) > 0 {
			containsQueries := []string{}
			containsArgs := []interface{}{}
			for _, v := range f.CreatedAt.Contains {
				if v != "" {
					containsQueries = append(containsQueries, "createdat LIKE ?")
					containsArgs = append(containsArgs, "%"+v+"%")
				}
			}
			if len(containsQueries) > 0 {
				query = query.Where("("+strings.Join(containsQueries, " OR ")+")", containsArgs...)
			}
		}
		if len(f.CreatedAt.NotContains) > 0 {
			notContainsQueries := []string{}
			notContainsArgs := []interface{}{}
			for _, v := range f.CreatedAt.NotContains {
				if v != "" {
					notContainsQueries = append(notContainsQueries, "createdat NOT LIKE ?")
					notContainsArgs = append(notContainsArgs, "%"+v+"%")
				}
			}
			if len(notContainsQueries) > 0 {
				query = query.Where("("+strings.Join(notContainsQueries, " OR ")+")", notContainsArgs...)
			}
		}
		if f.CreatedAt.IsEmpty {
			query = query.Where("createdat = '' OR createdat IS NULL")
		}
		if f.CreatedAt.IsNotEmpty {
			query = query.Where("createdat <> '' AND createdat IS NOT NULL")
		}
	}
	if f.UpdatedAt != nil {
		if f.UpdatedAt.Eq != "" {
			query = query.Where("updatedat = ?", f.UpdatedAt.Eq)
		}
		if f.UpdatedAt.Ne != "" {
			query = query.Where("updatedat <> ?", f.UpdatedAt.Ne)
		}
		if f.UpdatedAt.Lt != "" {
			query = query.Where("updatedat < ?", f.UpdatedAt.Lt)
		}
		if f.UpdatedAt.Le != "" {
			query = query.Where("updatedat <= ?", f.UpdatedAt.Le)
		}
		if f.UpdatedAt.Gt != "" {
			query = query.Where("updatedat > ?", f.UpdatedAt.Gt)
		}
		if f.UpdatedAt.Ge != "" {
			query = query.Where("updatedat >= ?", f.UpdatedAt.Ge)
		}
		if f.UpdatedAt.StartsWith != "" {
			query = query.Where("updatedat LIKE '%' || ?", f.UpdatedAt.StartsWith)
		}
		if f.UpdatedAt.EndsWith != "" {
			query = query.Where("updatedat LIKE ? || '%'", f.UpdatedAt.EndsWith)
		}
		if len(f.UpdatedAt.Contains) > 0 {
			containsQueries := []string{}
			containsArgs := []interface{}{}
			for _, v := range f.UpdatedAt.Contains {
				if v != "" {
					containsQueries = append(containsQueries, "updatedat LIKE ?")
					containsArgs = append(containsArgs, "%"+v+"%")
				}
			}
			if len(containsQueries) > 0 {
				query = query.Where("("+strings.Join(containsQueries, " OR ")+")", containsArgs...)
			}
		}
		if len(f.UpdatedAt.NotContains) > 0 {
			notContainsQueries := []string{}
			notContainsArgs := []interface{}{}
			for _, v := range f.UpdatedAt.NotContains {
				if v != "" {
					notContainsQueries = append(notContainsQueries, "updatedat NOT LIKE ?")
					notContainsArgs = append(notContainsArgs, "%"+v+"%")
				}
			}
			if len(notContainsQueries) > 0 {
				query = query.Where("("+strings.Join(notContainsQueries, " OR ")+")", notContainsArgs...)
			}
		}
		if f.UpdatedAt.IsEmpty {
			query = query.Where("updatedat = '' OR updatedat IS NULL")
		}
		if f.UpdatedAt.IsNotEmpty {
			query = query.Where("updatedat <> '' AND updatedat IS NOT NULL")
		}
	}
	return query
}
func (f *FilterLinkView) BuildMongoFilter() bson.M {
	if f == nil {
		return nil
	}
	filter := bson.M{}
	if f.FieldMask != nil {
		fieldFilter := bson.M{}
		if f.FieldMask.Eq != "" {
			fieldFilter["$eq"] = f.FieldMask.Eq
		}
		if f.FieldMask.Ne != "" {
			fieldFilter["$ne"] = f.FieldMask.Ne
		}
		if f.FieldMask.Lt != "" {
			fieldFilter["$lt"] = f.FieldMask.Lt
		}
		if f.FieldMask.Le != "" {
			fieldFilter["$lte"] = f.FieldMask.Le
		}
		if f.FieldMask.Gt != "" {
			fieldFilter["$gt"] = f.FieldMask.Gt
		}
		if f.FieldMask.Ge != "" {
			fieldFilter["$gte"] = f.FieldMask.Ge
		}
		if len(f.FieldMask.Contains) > 0 {
			fieldFilter["$in"] = f.FieldMask.Contains
		}
		if len(f.FieldMask.NotContains) > 0 {
			fieldFilter["$nin"] = f.FieldMask.NotContains
		}
		if len(fieldFilter) > 0 {
			filter["fieldmask"] = fieldFilter
		}
	}
	if f.Url != nil {
		fieldFilter := bson.M{}
		if f.Url.Eq != "" {
			fieldFilter["$eq"] = f.Url.Eq
		}
		if f.Url.Ne != "" {
			fieldFilter["$ne"] = f.Url.Ne
		}
		if f.Url.Lt != "" {
			fieldFilter["$lt"] = f.Url.Lt
		}
		if f.Url.Le != "" {
			fieldFilter["$lte"] = f.Url.Le
		}
		if f.Url.Gt != "" {
			fieldFilter["$gt"] = f.Url.Gt
		}
		if f.Url.Ge != "" {
			fieldFilter["$gte"] = f.Url.Ge
		}
		if len(f.Url.Contains) > 0 {
			fieldFilter["$in"] = f.Url.Contains
		}
		if len(f.Url.NotContains) > 0 {
			fieldFilter["$nin"] = f.Url.NotContains
		}
		if len(fieldFilter) > 0 {
			filter["url"] = fieldFilter
		}
	}
	if f.Hash != nil {
		fieldFilter := bson.M{}
		if f.Hash.Eq != "" {
			fieldFilter["$eq"] = f.Hash.Eq
		}
		if f.Hash.Ne != "" {
			fieldFilter["$ne"] = f.Hash.Ne
		}
		if f.Hash.Lt != "" {
			fieldFilter["$lt"] = f.Hash.Lt
		}
		if f.Hash.Le != "" {
			fieldFilter["$lte"] = f.Hash.Le
		}
		if f.Hash.Gt != "" {
			fieldFilter["$gt"] = f.Hash.Gt
		}
		if f.Hash.Ge != "" {
			fieldFilter["$gte"] = f.Hash.Ge
		}
		if len(f.Hash.Contains) > 0 {
			fieldFilter["$in"] = f.Hash.Contains
		}
		if len(f.Hash.NotContains) > 0 {
			fieldFilter["$nin"] = f.Hash.NotContains
		}
		if len(fieldFilter) > 0 {
			filter["hash"] = fieldFilter
		}
	}
	if f.Describe != nil {
		fieldFilter := bson.M{}
		if f.Describe.Eq != "" {
			fieldFilter["$eq"] = f.Describe.Eq
		}
		if f.Describe.Ne != "" {
			fieldFilter["$ne"] = f.Describe.Ne
		}
		if f.Describe.Lt != "" {
			fieldFilter["$lt"] = f.Describe.Lt
		}
		if f.Describe.Le != "" {
			fieldFilter["$lte"] = f.Describe.Le
		}
		if f.Describe.Gt != "" {
			fieldFilter["$gt"] = f.Describe.Gt
		}
		if f.Describe.Ge != "" {
			fieldFilter["$gte"] = f.Describe.Ge
		}
		if len(f.Describe.Contains) > 0 {
			fieldFilter["$in"] = f.Describe.Contains
		}
		if len(f.Describe.NotContains) > 0 {
			fieldFilter["$nin"] = f.Describe.NotContains
		}
		if len(fieldFilter) > 0 {
			filter["describe"] = fieldFilter
		}
	}
	if f.ImageUrl != nil {
		fieldFilter := bson.M{}
		if f.ImageUrl.Eq != "" {
			fieldFilter["$eq"] = f.ImageUrl.Eq
		}
		if f.ImageUrl.Ne != "" {
			fieldFilter["$ne"] = f.ImageUrl.Ne
		}
		if f.ImageUrl.Lt != "" {
			fieldFilter["$lt"] = f.ImageUrl.Lt
		}
		if f.ImageUrl.Le != "" {
			fieldFilter["$lte"] = f.ImageUrl.Le
		}
		if f.ImageUrl.Gt != "" {
			fieldFilter["$gt"] = f.ImageUrl.Gt
		}
		if f.ImageUrl.Ge != "" {
			fieldFilter["$gte"] = f.ImageUrl.Ge
		}
		if len(f.ImageUrl.Contains) > 0 {
			fieldFilter["$in"] = f.ImageUrl.Contains
		}
		if len(f.ImageUrl.NotContains) > 0 {
			fieldFilter["$nin"] = f.ImageUrl.NotContains
		}
		if len(fieldFilter) > 0 {
			filter["imageurl"] = fieldFilter
		}
	}
	if f.MetaDescription != nil {
		fieldFilter := bson.M{}
		if f.MetaDescription.Eq != "" {
			fieldFilter["$eq"] = f.MetaDescription.Eq
		}
		if f.MetaDescription.Ne != "" {
			fieldFilter["$ne"] = f.MetaDescription.Ne
		}
		if f.MetaDescription.Lt != "" {
			fieldFilter["$lt"] = f.MetaDescription.Lt
		}
		if f.MetaDescription.Le != "" {
			fieldFilter["$lte"] = f.MetaDescription.Le
		}
		if f.MetaDescription.Gt != "" {
			fieldFilter["$gt"] = f.MetaDescription.Gt
		}
		if f.MetaDescription.Ge != "" {
			fieldFilter["$gte"] = f.MetaDescription.Ge
		}
		if len(f.MetaDescription.Contains) > 0 {
			fieldFilter["$in"] = f.MetaDescription.Contains
		}
		if len(f.MetaDescription.NotContains) > 0 {
			fieldFilter["$nin"] = f.MetaDescription.NotContains
		}
		if len(fieldFilter) > 0 {
			filter["metadescription"] = fieldFilter
		}
	}
	if f.MetaKeywords != nil {
		fieldFilter := bson.M{}
		if f.MetaKeywords.Eq != "" {
			fieldFilter["$eq"] = f.MetaKeywords.Eq
		}
		if f.MetaKeywords.Ne != "" {
			fieldFilter["$ne"] = f.MetaKeywords.Ne
		}
		if f.MetaKeywords.Lt != "" {
			fieldFilter["$lt"] = f.MetaKeywords.Lt
		}
		if f.MetaKeywords.Le != "" {
			fieldFilter["$lte"] = f.MetaKeywords.Le
		}
		if f.MetaKeywords.Gt != "" {
			fieldFilter["$gt"] = f.MetaKeywords.Gt
		}
		if f.MetaKeywords.Ge != "" {
			fieldFilter["$gte"] = f.MetaKeywords.Ge
		}
		if len(f.MetaKeywords.Contains) > 0 {
			fieldFilter["$in"] = f.MetaKeywords.Contains
		}
		if len(f.MetaKeywords.NotContains) > 0 {
			fieldFilter["$nin"] = f.MetaKeywords.NotContains
		}
		if len(fieldFilter) > 0 {
			filter["metakeywords"] = fieldFilter
		}
	}
	if f.CreatedAt != nil {
		fieldFilter := bson.M{}
		if f.CreatedAt.Eq != "" {
			fieldFilter["$eq"] = f.CreatedAt.Eq
		}
		if f.CreatedAt.Ne != "" {
			fieldFilter["$ne"] = f.CreatedAt.Ne
		}
		if f.CreatedAt.Lt != "" {
			fieldFilter["$lt"] = f.CreatedAt.Lt
		}
		if f.CreatedAt.Le != "" {
			fieldFilter["$lte"] = f.CreatedAt.Le
		}
		if f.CreatedAt.Gt != "" {
			fieldFilter["$gt"] = f.CreatedAt.Gt
		}
		if f.CreatedAt.Ge != "" {
			fieldFilter["$gte"] = f.CreatedAt.Ge
		}
		if len(f.CreatedAt.Contains) > 0 {
			fieldFilter["$in"] = f.CreatedAt.Contains
		}
		if len(f.CreatedAt.NotContains) > 0 {
			fieldFilter["$nin"] = f.CreatedAt.NotContains
		}
		if len(fieldFilter) > 0 {
			filter["createdat"] = fieldFilter
		}
	}
	if f.UpdatedAt != nil {
		fieldFilter := bson.M{}
		if f.UpdatedAt.Eq != "" {
			fieldFilter["$eq"] = f.UpdatedAt.Eq
		}
		if f.UpdatedAt.Ne != "" {
			fieldFilter["$ne"] = f.UpdatedAt.Ne
		}
		if f.UpdatedAt.Lt != "" {
			fieldFilter["$lt"] = f.UpdatedAt.Lt
		}
		if f.UpdatedAt.Le != "" {
			fieldFilter["$lte"] = f.UpdatedAt.Le
		}
		if f.UpdatedAt.Gt != "" {
			fieldFilter["$gt"] = f.UpdatedAt.Gt
		}
		if f.UpdatedAt.Ge != "" {
			fieldFilter["$gte"] = f.UpdatedAt.Ge
		}
		if len(f.UpdatedAt.Contains) > 0 {
			fieldFilter["$in"] = f.UpdatedAt.Contains
		}
		if len(f.UpdatedAt.NotContains) > 0 {
			fieldFilter["$nin"] = f.UpdatedAt.NotContains
		}
		if len(fieldFilter) > 0 {
			filter["updatedat"] = fieldFilter
		}
	}
	return filter
}

type FilterLinksView struct {
}

func (f *FilterLinksView) BuildFilter(query squirrel.SelectBuilder) squirrel.SelectBuilder {
	return query
}
func (f *FilterLinksView) BuildMongoFilter() bson.M {
	if f == nil {
		return nil
	}
	filter := bson.M{}
	return filter
}
