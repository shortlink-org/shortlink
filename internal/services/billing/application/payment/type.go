package payment_application

import (
	eventsourcing "github.com/batazor/shortlink/internal/pkg/eventsourcing/v1"
	"github.com/batazor/shortlink/internal/pkg/notify"
	billing "github.com/batazor/shortlink/internal/services/billing/domain/billing/payment/v1"
)

type Payment struct {
	*eventsourcing.BaseAggregate
	*billing.Payment
}

var (
	// EventList - event notify list
	EventList map[string]uint32
)

func init() {
	EventList = make(map[string]uint32, 0)

	for event := range billing.Event_name {
		EventList[billing.Event_name[event]] = notify.NewEventID()
	}
}
