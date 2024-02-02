package resolver

type StringFilterInput struct {
	Ne          *string
	Eq          *string
	Le          *string
	Lt          *string
	Ge          *string
	Gt          *string
	Contains    *string
	NotContains *string
	Between     *[]*string
	BeginsWith  *string
}
