package payment_application

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
	"github.com/segmentio/encoding/json"
	"github.com/spf13/viper"

	billing "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/domain/payment/v1"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/logger/field"
	"github.com/shortlink-org/shortlink/pkg/notify"
	es "github.com/shortlink-org/shortlink/pkg/pattern/eventsourcing"
	eventsourcing "github.com/shortlink-org/shortlink/pkg/pattern/eventsourcing/domain/eventsourcing/v1"
	"github.com/shortlink-org/shortlink/pkg/pattern/saga"
)

type PaymentService struct {
	log logger.Logger

	// EventSourcing
	eventsourcing.CommandHandle

	// Repositories
	paymentRepository es.EventSourcing
}

func New(log logger.Logger, paymentRepository es.EventSourcing) (*PaymentService, error) {
	service := &PaymentService{
		log: log,

		// Repositories
		paymentRepository: paymentRepository,
	}

	err := service.initTask()
	if err != nil {
		return nil, err
	}

	return service, nil
}

func (p *PaymentService) Handle(ctx context.Context, aggregate *Payment, command *eventsourcing.BaseCommand) error {
	// Check update or create
	if command.GetVersion() != 0 { //nolint:nestif // ignore
		snapshot, events, errLoad := p.paymentRepository.Load(ctx, command.GetAggregateId())
		if errLoad != nil {
			return errLoad
		}

		if snapshot.GetPayload() != "" {
			aggregate.Version = snapshot.GetAggregateVersion()
			err := json.Unmarshal([]byte(snapshot.GetPayload()), aggregate.Payment)
			if err != nil {
				return err
			}
		}

		for _, event := range events {
			errApplyChange := aggregate.ApplyChangeHelper(ctx, aggregate, event, false)
			if errApplyChange != nil {
				return errApplyChange
			}
		}
	}

	err := aggregate.HandleCommand(ctx, command)
	if err != nil {
		return err
	}

	err = p.paymentRepository.Save(ctx, aggregate.Uncommitted())
	if err != nil {
		return err
	}

	err = p.PublishEvents(ctx, aggregate.Uncommitted())
	if err != nil {
		return err
	}

	return nil
}

// PublishEvents - send message about a new events
func (p *PaymentService) PublishEvents(ctx context.Context, events []*eventsourcing.Event) error {
	for key := range events {
		go notify.Publish(ctx, EventList[events[key].GetType()], events[key].GetPayload(), nil)
	}

	return nil
}

func (p *PaymentService) Get(ctx context.Context, aggregateId string) (*billing.Payment, error) {
	aggregate := &Payment{
		Payment:       &billing.Payment{},
		BaseAggregate: &eventsourcing.BaseAggregate{},
	}

	snapshot, events, err := p.paymentRepository.Load(ctx, aggregateId)
	if err != nil {
		return nil, err
	}

	if snapshot.GetPayload() != "" {
		err = json.Unmarshal([]byte(snapshot.GetPayload()), aggregate.Payment)
		if err != nil {
			return nil, err
		}
	}

	for _, event := range events {
		errApplyChange := aggregate.ApplyChangeHelper(ctx, aggregate, event, false)
		if errApplyChange != nil {
			return nil, errApplyChange
		}
	}

	return aggregate.Payment, nil
}

func (p *PaymentService) List(ctx context.Context, filter any) ([]*billing.Payment, error) {
	panic("implement me")
	// return p.paymentRepository.list(ctx, filter)
}

func errorHelper(ctx context.Context, log logger.Logger, errs []error) error {
	if len(errs) > 0 {
		errList := field.Fields{}
		for index := range errs {
			errList[fmt.Sprintf("stack error: %d", index)] = errs[index]
		}

		log.ErrorWithContext(ctx, "Error create a new payment", errList)

		return ErrCreatePayment
	}

	return nil
}

// Add - Create a payment
func (p *PaymentService) Add(ctx context.Context, in *billing.Payment) (*billing.Payment, error) {
	const (
		SAGA_NAME                 = "ADD_PAYMENT"
		SAGA_STEP_PAYMENT_CREATE  = "SAGA_STEP_PAYMENT_CREATE"
		SAGA_STEP_PAYMENT_APPROVE = "SAGA_STEP_PAYMENT_APPROVE"
		SAGA_STEP_PAYMENT_GET     = "SAGA_STEP_PAYMENT_GET"
	)

	// saga for create a new payment
	sagaAddPayment, errs := saga.New(SAGA_NAME, saga.SetLogger(p.log)).
		WithContext(ctx).
		Build()
	if err := errorHelper(ctx, p.log, errs); err != nil {
		return nil, err
	}

	// add step: create a payment
	_, errs = sagaAddPayment.AddStep(SAGA_STEP_PAYMENT_CREATE).
		Then(func(ctx context.Context) error {
			aggregate := &Payment{
				Payment:       &billing.Payment{},
				BaseAggregate: &eventsourcing.BaseAggregate{},
			}

			command, err := CommandPaymentCreate(ctx, in)
			if err != nil {
				return err
			}

			err = p.Handle(ctx, aggregate, command)
			if err != nil {
				return err
			}

			return nil
		}).
		Reject(func(ctx context.Context, thenErr error) error {
			return ErrCreatePayment
		}).
		Build()
	if err := errorHelper(ctx, p.log, errs); err != nil {
		return nil, err
	}

	// add step: approve/reject payment
	_, errs = sagaAddPayment.AddStep(SAGA_STEP_PAYMENT_APPROVE).
		Needs(SAGA_STEP_PAYMENT_CREATE).
		Then(func(ctx context.Context) error {
			return p.Approve(ctx, in.GetId())
		}).
		Reject(func(ctx context.Context, thenErr error) error {
			err := p.Reject(ctx, in.GetId())
			return err
		}).
		Build()
	if err := errorHelper(ctx, p.log, errs); err != nil {
		return nil, err
	}

	// add step: get actual state a payment
	_, errs = sagaAddPayment.AddStep(SAGA_STEP_PAYMENT_GET).
		Needs(SAGA_STEP_PAYMENT_APPROVE).
		Then(func(ctx context.Context) error {
			var err error
			in, err = p.Get(ctx, in.GetId().String())
			if err != nil {
				return err
			}

			return nil
		}).
		Reject(func(ctx context.Context, thenErr error) error {
			return ErrApprovePayment
		}).
		Build()
	if err := errorHelper(ctx, p.log, errs); err != nil {
		return nil, err
	}

	// Run saga
	err := sagaAddPayment.Play(nil)
	if err != nil {
		return nil, err
	}

	return in, nil
}

func (p *PaymentService) Approve(ctx context.Context, id uuid.UUID) error {
	aggregate := &Payment{
		Payment:       &billing.Payment{},
		BaseAggregate: &eventsourcing.BaseAggregate{},
	}

	command, err := CommandPaymentApprove(ctx, id)
	if err != nil {
		return err
	}

	err = p.Handle(ctx, aggregate, command)
	if err != nil {
		return err
	}

	return nil
}

func (p *PaymentService) Reject(ctx context.Context, id uuid.UUID) error {
	aggregate := &Payment{
		Payment:       &billing.Payment{},
		BaseAggregate: &eventsourcing.BaseAggregate{},
	}

	command, err := CommandPaymentReject(ctx, id)
	if err != nil {
		return err
	}

	// set version `0` for do insert
	err = p.Handle(ctx, aggregate, command)
	if err != nil {
		return err
	}

	return nil
}

func (p *PaymentService) Close(ctx context.Context, id uuid.UUID) error {
	aggregate := &Payment{
		Payment:       &billing.Payment{},
		BaseAggregate: &eventsourcing.BaseAggregate{},
	}

	command, err := CommandPaymentClose(ctx, id)
	if err != nil {
		return err
	}

	err = p.Handle(ctx, aggregate, command)
	if err != nil {
		return err
	}

	return nil
}

func (p *PaymentService) UpdateBalance(ctx context.Context, in *billing.Payment) (*billing.Payment, error) {
	aggregate := &Payment{
		Payment:       &billing.Payment{},
		BaseAggregate: &eventsourcing.BaseAggregate{},
	}

	command, err := CommandPaymentUpdateBalance(ctx, in)
	if err != nil {
		return nil, err
	}

	err = p.Handle(ctx, aggregate, command)
	if err != nil {
		return nil, err
	}

	return aggregate.Payment, nil
}

func (p *PaymentService) initTask() error {
	viper.AutomaticEnv()
	viper.SetDefault("PAYMENT_SNAPSHOT_CRON", "* * * * *") // check snapshot by timeout

	c := cron.New()
	// CRON Expression Format
	// https://pkg.go.dev/github.com/robfig/cron#hdr-CRON_Expression_Format
	_, err := c.AddFunc(viper.GetString("PAYMENT_SNAPSHOT_CRON"), func() {
		p.asyncUpdateSnapshot()
	})
	if err != nil {
		return err
	}
	c.Start()

	return nil
}

func (p *PaymentService) asyncUpdateSnapshot() {
	ctx := context.Background()
	p.log.InfoWithContext(ctx, "Run asyncUpdateSnapshot")

	aggregates, errGetAggregate := p.paymentRepository.GetAggregateWithoutSnapshot(ctx)
	if errGetAggregate != nil {
		p.log.ErrorWithContext(ctx, errGetAggregate.Error())
		return
	}

	for key := range aggregates {
		payment, err := p.Get(ctx, aggregates[key].GetId())
		if err != nil {
			p.log.ErrorWithContext(ctx, err.Error())
			return
		}

		payload, err := json.Marshal(payment)
		if err != nil {
			p.log.ErrorWithContext(ctx, err.Error())
			return
		}

		snapshot := &eventsourcing.Snapshot{
			AggregateId:      aggregates[key].GetId(),
			AggregateType:    aggregates[key].GetType(),
			AggregateVersion: aggregates[key].GetVersion(),
			Payload:          string(payload),
		}

		// save or update
		err = p.paymentRepository.SaveSnapshot(ctx, snapshot)
		if err != nil {
			p.log.ErrorWithContext(ctx, err.Error())
			return
		}
	}
}
