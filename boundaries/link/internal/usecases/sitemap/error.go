package sitemap

import (
	"fmt"
)

// IncorrectResponseCodeError is an error returned when the response code is not 200.
type IncorrectResponseCodeError struct {
	StatusCode int
	URL        string
}

func (e *IncorrectResponseCodeError) Error() string {
	return fmt.Sprintf(`incorrect response code: %d for %s`, e.StatusCode, e.URL)
}
