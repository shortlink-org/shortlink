package options

func (o OptionError) Error() string {
	return "Options[T] has no value"
}

func New[T any]() *Option[T] {
	return &Option[T]{}
}

func (o Option[T]) Take() (T, error) {
	if o.IsNone() {
		var zero T
		return zero, OptionError{}
	}

	return *o.val, nil
}

func (o *Option[T]) Set(val T) {
	o.val = &val
}

func (o *Option[T]) Clear() {
	o.val = nil
}

func (o Option[T]) IsSome() bool {
	return o.val != nil
}

func (o Option[T]) IsNone() bool {
	return !o.IsSome()
}

func (o Option[T]) Apply() T {
	if o.IsNone() {
		panic("gonads: Yank on None Option")
	}

	return *o.val
}
