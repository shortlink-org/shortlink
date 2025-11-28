package rules

// EmailValidationData represents the data needed for email validation.
type EmailValidationData struct {
	Email      string
	SeenEmails map[string]bool
}

