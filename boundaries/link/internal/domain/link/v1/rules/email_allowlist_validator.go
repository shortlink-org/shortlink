package rules

import "github.com/shortlink-org/go-sdk/specification"

// ValidateEmailAllowlist validates a list of emails using specification pattern.
// Returns normalized emails and any validation errors.
func ValidateEmailAllowlist(emails []string) ([]string, error) {
	if len(emails) == 0 {
		return []string{}, nil
	}

	// Check size limit
	sizeSpec := NewAllowlistSizeSpecification()
	if err := sizeSpec.IsSatisfiedBy(&emails); err != nil {
		return nil, err
	}

	// Validate each email
	normalizedEmails := make([]string, 0, len(emails))
	seenEmails := make(map[string]bool, len(emails))

	for _, email := range emails {
		validationData := &EmailValidationData{
			Email:      email,
			SeenEmails: seenEmails,
		}

		// Combine specifications for email validation
		emailSpecs := specification.NewAndSpecification[EmailValidationData](
			NewEmailNonEmptySpecification(email),
			NewEmailFormatSpecification(email),
			NewEmailDuplicateSpecification(email),
		)

		if err := emailSpecs.IsSatisfiedBy(validationData); err != nil {
			return nil, err
		}

		// Normalize and add to seen emails
		normalized := NormalizeEmail(email)
		seenEmails[normalized] = true
		normalizedEmails = append(normalizedEmails, normalized)
	}

	return normalizedEmails, nil
}

