package resolver

import "github.com/batazor/shortlink/pkg/link"

type LinkResolver struct {
	link.Link
}

func (_ *Resolver) Link() (*LinkResolver, error) {
	return &LinkResolver{}, nil
}

func (r *LinkResolver) Url() string {
	return ""
}

func (r *LinkResolver) Hash() string {
	return ""
}

func (r *LinkResolver) Describe() string {
	return ""
}
