package payment_application

import (
	billing "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/domain/payment/v1"
	"github.com/shortlink-org/shortlink/pkg/notify"
	eventsourcing "github.com/shortlink-org/shortlink/pkg/pattern/eventsourcing/domain/eventsourcing/v1"
)

type Payment struct {
	*eventsourcing.BaseAggregate
	*billing.Payment
}

// EventList - event notify list
var EventList map[string]uint32

func init() {
	EventList = make(map[string]uint32)

	for event := range billing.Event_name {
		EventList[billing.Event_name[event]] = notify.NewEventID()
	}
}
