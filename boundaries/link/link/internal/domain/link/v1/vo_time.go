package v1

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Time time.Time

func (t Time) GetTime() time.Time {
	return time.Time(t)
}

func (t Time) GetTimestamp() *timestamppb.Timestamp {
	return timestamppb.New(time.Time(t))
}
