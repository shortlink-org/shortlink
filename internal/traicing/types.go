package traicing

type key int

const (
	keyTraicer key = iota
)

// Config ...
type Config struct { // nolint unused
	ServiceName string
	URI         string
}
