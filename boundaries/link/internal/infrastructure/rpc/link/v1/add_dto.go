package v1

import (
	domain "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

func (in *AddRequest) ToEntity() (*domain.Link, error) {
	builder := domain.NewLinkBuilder().
		SetURL(in.GetLink().GetUrl()).
		SetDescribe(in.GetLink().GetDescribe())

	// Set allowed_emails if provided
	if allowedEmails := in.GetLink().GetAllowedEmails(); len(allowedEmails) > 0 {
		builder = builder.SetAllowedEmails(allowedEmails)
	}

	link, err := builder.Build()
	if err != nil {
		return nil, domain.NewInvalidInputError(err.Error())
	}

	return link, nil
}

// ToAddResponse converts domain Link to AddResponse
func ToAddResponse(link *domain.Link) *AddResponse {
	allowedEmails := link.GetAllowedEmails()
	if allowedEmails == nil {
		allowedEmails = []string{}
	}

	return &AddResponse{
		Link: &Link{
			Url:          link.GetUrl().String(),
			Hash:         link.GetHash(),
			Describe:     link.GetDescribe(),
			CreatedAt:    link.GetCreatedAt().GetTimestamp(),
			UpdatedAt:    link.GetUpdatedAt().GetTimestamp(),
			AllowedEmails: allowedEmails,
		},
	}
}
