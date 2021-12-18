package v1

func New() (*Session, error) {
	return &Session{
		CurrentDatabase: "public",
		Raw:             "",
		Exec:            false,
	}, nil
}
