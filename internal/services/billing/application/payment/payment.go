package payment_application

import (
	"context"
	"fmt"

	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/encoding/protojson"

	event_store "github.com/shortlink-org/shortlink/internal/pkg/eventsourcing/store"
	eventsourcing "github.com/shortlink-org/shortlink/internal/pkg/eventsourcing/v1"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
	"github.com/shortlink-org/shortlink/internal/pkg/notify"
	"github.com/shortlink-org/shortlink/internal/pkg/saga"
	billing "github.com/shortlink-org/shortlink/internal/services/billing/domain/billing/payment/v1"
)

type PaymentService struct {
	logger logger.Logger

	// EventSourcing
	eventsourcing.CommandHandle

	// Repositories
	paymentRepository event_store.EventStore
}

func New(logger logger.Logger, paymentRepository event_store.EventStore) (*PaymentService, error) {
	service := &PaymentService{
		logger: logger,

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
	if command.Version != 0 { // nolint:nestif
		snapshot, events, errLoad := p.paymentRepository.Load(ctx, command.AggregateId)
		if errLoad != nil {
			return errLoad
		}

		if snapshot.Payload != "" {
			aggregate.Version = snapshot.AggregateVersion
			err := protojson.Unmarshal([]byte(snapshot.Payload), aggregate.Payment)
			if err != nil {
				return err
			}
		}

		for _, event := range events {
			errApplyChange := aggregate.ApplyChangeHelper(aggregate, event, false)
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
		go notify.Publish(ctx, EventList[events[key].Type], events[key].GetPayload(), nil)
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

	if snapshot.Payload != "" {
		err = protojson.Unmarshal([]byte(snapshot.Payload), aggregate.Payment)
		if err != nil {
			return nil, err
		}
	}

	for _, event := range events {
		errApplyChange := aggregate.ApplyChangeHelper(aggregate, event, false)
		if errApplyChange != nil {
			return nil, errApplyChange
		}
	}

	return aggregate.Payment, nil
}

func (p *PaymentService) List(ctx context.Context, filter interface{}) ([]*billing.Payment, error) {
	panic("implement me")
	// return p.paymentRepository.List(ctx, filter)
}

func errorHelper(ctx context.Context, logger logger.Logger, errs []error) error {
	if len(errs) > 0 {
		errList := field.Fields{}
		for index := range errs {
			errList[fmt.Sprintf("stack error: %d", index)] = errs[index]
		}

		logger.ErrorWithContext(ctx, "Error create a new payment", errList)

		return fmt.Errorf("Error create a new payment")
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
	sagaAddPayment, errs := saga.New(SAGA_NAME, saga.SetLogger(p.logger)).
		WithContext(ctx).
		Build()
	if err := errorHelper(ctx, p.logger, errs); err != nil {
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

			// safe identity
			in.Id = command.AggregateId

			return nil
		}).
		Reject(func(ctx context.Context) error {
			return fmt.Errorf("Error create a new payment")
		}).
		Build()
	if err := errorHelper(ctx, p.logger, errs); err != nil {
		return nil, err
	}

	// add step: approve/reject payment
	_, errs = sagaAddPayment.AddStep(SAGA_STEP_PAYMENT_APPROVE).
		Needs([]string{SAGA_STEP_PAYMENT_CREATE}).
		Then(func(ctx context.Context) error {
			return p.Approve(ctx, in.Id)
		}).
		Reject(func(ctx context.Context) error {
			err := p.Reject(ctx, in.Id)
			return err
		}).
		Build()
	if err := errorHelper(ctx, p.logger, errs); err != nil {
		return nil, err
	}

	// add step: get actual state a payment
	_, errs = sagaAddPayment.AddStep(SAGA_STEP_PAYMENT_GET).
		Needs([]string{SAGA_STEP_PAYMENT_APPROVE}).
		Then(func(ctx context.Context) error {
			var err error
			in, err = p.Get(ctx, in.Id)
			if err != nil {
				return err
			}

			return nil
		}).
		Reject(func(ctx context.Context) error {
			return fmt.Errorf(`Payment was successfully created, but its status could not be received`)
		}).
		Build()
	if err := errorHelper(ctx, p.logger, errs); err != nil {
		return nil, err
	}

	// Run saga
	err := sagaAddPayment.Play(nil)
	if err != nil {
		return nil, err
	}

	return in, nil
}

func (p *PaymentService) Approve(ctx context.Context, id string) error {
	aggregate := &Payment{
		Payment:       &billing.Payment{},
		BaseAggregate: &eventsourcing.BaseAggregate{},
	}

	command, err := CommandPaymentApprove(ctx, &billing.Payment{
		Id: id,
	})
	if err != nil {
		return err
	}

	err = p.Handle(ctx, aggregate, command)
	if err != nil {
		return err
	}

	return nil
}

func (p *PaymentService) Reject(ctx context.Context, id string) error {
	aggregate := &Payment{
		Payment:       &billing.Payment{},
		BaseAggregate: &eventsourcing.BaseAggregate{},
	}

	command, err := CommandPaymentReject(ctx, &billing.Payment{
		Id: id,
	})
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

func (p *PaymentService) Close(ctx context.Context, id string) error {
	aggregate := &Payment{
		Payment:       &billing.Payment{},
		BaseAggregate: &eventsourcing.BaseAggregate{},
	}

	command, err := CommandPaymentClose(ctx, &billing.Payment{
		Id: id,
	})
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
	p.logger.InfoWithContext(ctx, "Run asyncUpdateSnapshot")

	aggregates, errGetAggregate := p.paymentRepository.GetAggregateWithoutSnapshot(ctx)
	if errGetAggregate != nil {
		p.logger.ErrorWithContext(ctx, errGetAggregate.Error())
		return
	}

	for key := range aggregates {
		payment, err := p.Get(ctx, aggregates[key].Id)
		if err != nil {
			p.logger.ErrorWithContext(ctx, err.Error())
			return
		}

		payload, err := protojson.Marshal(payment)
		if err != nil {
			p.logger.ErrorWithContext(ctx, err.Error())
			return
		}

		snapshot := &eventsourcing.Snapshot{
			AggregateId:      aggregates[key].Id,
			AggregateType:    aggregates[key].Type,
			AggregateVersion: aggregates[key].Version,
			Payload:          string(payload),
		}

		// save or update
		err = p.paymentRepository.SaveSnapshot(ctx, snapshot)
		if err != nil {
			p.logger.ErrorWithContext(ctx, err.Error())
			return
		}
	}
}
