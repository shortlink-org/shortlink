package resolver

type LinkResolver struct {
	// *link.Link
}

// func (r *LinkResolver) Url() string {
// 	return r.Link.Url
// }
//
// func (r *LinkResolver) Hash() string {
// 	return r.Link.Hash
// }
//
// func (r *LinkResolver) Describe() string {
// 	return r.Link.Describe
// }
//
// func (r *LinkResolver) Created_at() graphql.Time {
// 	var cr time.Time
// 	if r.Link.CreatedAt != nil {
// 		cr = time.Unix(r.Link.CreatedAt.Seconds, int64(r.Link.CreatedAt.Nanos))
// 	}
// 	return graphql.Time{Time: cr}
// }
//
// func (r *LinkResolver) Updated_at() graphql.Time {
// 	var cr time.Time
// 	if r.Link.UpdatedAt != nil {
// 		cr = time.Unix(r.Link.UpdatedAt.Seconds, int64(r.Link.UpdatedAt.Nanos))
// 	}
// 	return graphql.Time{Time: cr}
// }

type LinkFilterInput struct {
	Url        *StringFilterInput
	Hash       *StringFilterInput
	Describe   *StringFilterInput
	Created_at *StringFilterInput
	Updated_at *StringFilterInput
}
