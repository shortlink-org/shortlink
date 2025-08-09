package v1

import (
	"net/url"
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

type Url url.URL

func (u Url) GetUrl() url.URL {
	return url.URL(u)
}

func (u Url) String() string {
	tmp := url.URL(u)

	return tmp.String()
}
