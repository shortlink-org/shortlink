package v1

// StringFilterInput represents filter conditions for string fields
type StringFilterInput struct {
	Eq          string
	Ne          string
	Lt          string
	Le          string
	Gt          string
	Ge          string
	Contains    []string
	NotContains []string
	StartsWith  string
	EndsWith    string
	IsEmpty     bool
	IsNotEmpty  bool
}

// FilterLink represents filter conditions for Link queries
type FilterLink struct {
	URL       *StringFilterInput
	Hash      *StringFilterInput
	Describe  *StringFilterInput
	CreatedAt *StringFilterInput
	UpdatedAt *StringFilterInput
}

