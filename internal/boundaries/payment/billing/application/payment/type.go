package payment_application

import (
	billing "github.com/shortlink-org/shortlink/internal/boundaries/payment/billing/domain/billing/payment/v1"
	eventsourcing "github.com/shortlink-org/shortlink/internal/pkg/eventsourcing/v1"
	"github.com/shortlink-org/shortlink/internal/pkg/notify"
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
