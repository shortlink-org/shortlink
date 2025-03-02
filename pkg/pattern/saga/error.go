package saga

type DublicateStepError struct {
	Name string
}

func (e *DublicateStepError) Error() string {
	return "dublicate step: " + e.Name
}
