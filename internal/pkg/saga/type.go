package saga

type StepState int

// Success chain: INIT -> WAIT -> RUN -> DONE
// Fail chain: INIT -> WAIT -> RUN -> REJECT -> FAIL or ROLLBACK
const (
	INIT StepState = iota + 1
	WAIT
	RUN
	DONE
	REJECT
	FAIL
	ROLLBACK
)
