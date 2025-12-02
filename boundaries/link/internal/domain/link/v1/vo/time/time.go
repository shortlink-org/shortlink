package time

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// Time is a Value Object representing a time value.
type Time time.Time

// NewTime creates a new Time Value Object from a time.Time.
func NewTime(t time.Time) Time {
	return Time(t)
}

// GetTime returns the time.Time value.
func (t Time) GetTime() time.Time {
	return time.Time(t)
}

// GetTimestamp returns the protobuf Timestamp representation.
func (t Time) GetTimestamp() *timestamppb.Timestamp {
	return timestamppb.New(time.Time(t))
}

// IsZero checks if the time is zero.
func (t Time) IsZero() bool {
	return time.Time(t).IsZero()
}


