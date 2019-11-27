package query

type Filter struct { // nolint unused
	Url      *StringFilterInput
	Hash     *StringFilterInput
	Describe *StringFilterInput
}

type StringFilterInput struct { // nolint unused
	Ne          *string
	Eq          *string
	Le          *string
	Lt          *string
	Ge          *string
	Gt          *string
	Contains    *string
	NotContains *string
}
