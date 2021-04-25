package saga

type StepState int

// Success chain: INIT -> WAIT -> RUN -> DONE
const (
	INIT StepState = iota
	WAIT
	RUN
	DONE
	REJECT
	FAIL
	ROLLBACK
)

type Store interface{}
