package specification

import (
	"errors"
)

func Filter[T any](list []*T, spec Specification[T]) ([]*T, error) {
	var errs error
	result := make([]*T, 0, len(list))

	for _, item := range list {
		err := spec.IsSatisfiedBy(item)
		if err != nil {
			errs = errors.Join(errs, err)
			continue
		}

		result = append(result, item)
	}

	return result, errs
}
