package v1

// NotFoundError - not found link
type NotFoundError struct {
	Hash string
}

func (e *NotFoundError) Error() string {
	if e.Hash != "" {
		return "Not found link: " + e.Hash
	}

	return "Not found link"
}

// NotFoundByHashError - not found link by hash
type NotFoundByHashError struct {
	Hash string
}

func (e *NotFoundByHashError) Error() string {
	return "Not found link by hash: " + e.Hash
}

// CreateLinkError - create link error
type CreateLinkError struct {
	Hash string
}

func (e *CreateLinkError) Error() string {
	if e.Hash != "" {
		return "Create link error: " + e.Hash
	}

	return "Create link error"
}
