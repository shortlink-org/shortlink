package order_workflow

// import (
// 	"github.com/google/uuid"
// 	"go.temporal.io/sdk/workflow"
//
// 	v2 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/order/v1"
// 	v3 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/order/v1/model/v1"
// 	"github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/workers/order/workflow/dto"
// 	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/workers/order/workflow/model/order/v1"
// )
//
// // Workflow is a Temporal workflow that manages the order state.
// func Workflow(ctx workflow.Context, customerId uuid.UUID) error {
// 	state := v2.NewOrderState(customerId)
//
// 	// Set up query handler for getting order state
// 	err := workflow.SetQueryHandler(ctx, v2.Event_EVENT_GET.String(), func() (*v3.OrderState, error) {
// 		return dto.OrderStateToDomain(state), nil
// 	})
// 	if err != nil {
// 		return err
// 	}
//
// 	logger := workflow.GetLogger(ctx)
//
// 	createOrderChannel := workflow.GetSignalChannel(ctx, v2.Event_EVENT_CREATE.String())
// 	updateOrderChannel := workflow.GetSignalChannel(ctx, v2.Event_EVENT_UPDATE.String())
// 	cancelOrderChannel := workflow.GetSignalChannel(ctx, v2.Event_EVENT_CANCEL.String())
// 	completeOrderChannel := workflow.GetSignalChannel(ctx, v2.Event_EVENT_COMPLETE.String())
//
// 	selector := workflow.NewSelector(ctx)
//
// 	selector.AddReceive(createOrderChannel, func(c workflow.ReceiveChannel, _ bool) {
// 		var request v1.OrderEvent
// 		c.Receive(ctx, &request)
//
// 		ao := workflow.ActivityOptions{
// 			StartToCloseTimeout: workflow.DefaultActivityOptions(ctx).StartToCloseTimeout,
// 		}
// 		ctx = workflow.WithActivityOptions(ctx, ao)
//
// 		// Process payment
// 		err := workflow.ExecuteActivity(ctx, ProcessPaymentActivity, request.Order).Get(ctx, nil)
// 		if err != nil {
// 			logger.Error("Payment processing failed: %v", err)
// 			return
// 		}
//
// 		// Schedule delivery
// 		err = workflow.ExecuteActivity(ctx, ScheduleDeliveryActivity, request.Order).Get(ctx, nil)
// 		if err != nil {
// 			logger.Error("Delivery scheduling failed: %v", err)
// 			return
// 		}
//
// 		// Send order confirmation notification
// 		err = workflow.ExecuteActivity(ctx, SendNotificationActivity, "OrderConfirmation", request.Order).Get(ctx, nil)
// 		if err != nil {
// 			logger.Error("Notification sending failed: %v", err)
// 			return
// 		}
//
// 		state.CreateOrder(request.Order)
// 	})
//
// 	selector.AddReceive(updateOrderChannel, func(c workflow.ReceiveChannel, _ bool) {
// 		var request v1.OrderEvent
// 		c.Receive(ctx, &request)
//
// 		ao := workflow.ActivityOptions{
// 			StartToCloseTimeout: workflow.DefaultActivityOptions(ctx).StartToCloseTimeout,
// 		}
// 		ctx = workflow.WithActivityOptions(ctx, ao)
//
// 		// Update payment details
// 		err := workflow.ExecuteActivity(ctx, UpdatePaymentActivity, request.Order).Get(ctx, nil)
// 		if err != nil {
// 			logger.Error("Payment update failed: %v", err)
// 			return
// 		}
//
// 		// Reschedule delivery
// 		err = workflow.ExecuteActivity(ctx, RescheduleDeliveryActivity, request.Order).Get(ctx, nil)
// 		if err != nil {
// 			logger.Error("Delivery rescheduling failed: %v", err)
// 			return
// 		}
//
// 		// Send order update notification
// 		err = workflow.ExecuteActivity(ctx, SendNotificationActivity, "OrderUpdate", request.Order).Get(ctx, nil)
// 		if err != nil {
// 			logger.Error("Notification sending failed: %v", err)
// 			return
// 		}
//
// 		state.UpdateOrder(request.Order)
// 	})
//
// 	selector.AddReceive(cancelOrderChannel, func(c workflow.ReceiveChannel, _ bool) {
// 		var request v1.OrderEvent
// 		c.Receive(ctx, &request)
//
// 		ao := workflow.ActivityOptions{
// 			StartToCloseTimeout: workflow.DefaultActivityOptions(ctx).StartToCloseTimeout,
// 		}
// 		ctx = workflow.WithActivityOptions(ctx, ao)
//
// 		// Refund payment
// 		err := workflow.ExecuteActivity(ctx, RefundPaymentActivity, request.Order).Get(ctx, nil)
// 		if err != nil {
// 			logger.Error("Payment refund failed: %v", err)
// 			return
// 		}
//
// 		// Cancel delivery
// 		err = workflow.ExecuteActivity(ctx, CancelDeliveryActivity, request.Order).Get(ctx, nil)
// 		if err != nil {
// 			logger.Error("Delivery cancellation failed: %v", err)
// 			return
// 		}
//
// 		// Send order cancellation notification
// 		err = workflow.ExecuteActivity(ctx, SendNotificationActivity, "OrderCancellation", request.Order).Get(ctx, nil)
// 		if err != nil {
// 			logger.Error("Notification sending failed: %v", err)
// 			return
// 		}
//
// 		state.CancelOrder(request.Order)
// 	})
//
// 	selector.AddReceive(completeOrderChannel, func(c workflow.ReceiveChannel, _ bool) {
// 		var request v1.OrderEvent
// 		c.Receive(ctx, &request)
//
// 		ao := workflow.ActivityOptions{
// 			StartToCloseTimeout: workflow.DefaultActivityOptions(ctx).StartToCloseTimeout,
// 		}
// 		ctx = workflow.WithActivityOptions(ctx, ao)
//
// 		// Send order completion notification
// 		err := workflow.ExecuteActivity(ctx, SendNotificationActivity, "OrderCompletion", request.Order).Get(ctx, nil)
// 		if err != nil {
// 			logger.Error("Notification sending failed: %v", err)
// 			return
// 		}
//
// 		state.CompleteOrder(request.Order)
//
// 		// Break the loop to finish the workflow
// 		workflow.GetSignalChannel(ctx, "break").Send(ctx, true)
// 	})
//
// 	for {
// 		select {
// 		case <-workflow.GetSignalChannel(ctx, "break"):
// 			logger.Info("Order completed successfully. Finishing workflow.")
// 			return nil
// 		default:
// 			selector.Select(ctx)
// 		}
// 	}
// }
//
// // Activities ------------------------------------------------------------------
//
// func ProcessPaymentActivity(ctx context.Context, order v1.Order) error {
// 	// Logic to process payment
// 	return nil
// }
//
// func UpdatePaymentActivity(ctx context.Context, order v1.Order) error {
// 	// Logic to update payment
// 	return nil
// }
//
// func RefundPaymentActivity(ctx context.Context, order v1.Order) error {
// 	// Logic to refund payment
// 	return nil
// }
//
// func ScheduleDeliveryActivity(ctx context.Context, order v1.Order) error {
// 	// Logic to schedule delivery
// 	return nil
// }
//
// func RescheduleDeliveryActivity(ctx context.Context, order v1.Order) error {
// 	// Logic to reschedule delivery
// 	return nil
// }
//
// func CancelDeliveryActivity(ctx context.Context, order v1.Order) error {
// 	// Logic to cancel delivery
// 	return nil
// }
//
// func SendNotificationActivity(ctx context.Context, notificationType string, order v1.Order) error {
// 	// Logic to send notification (e.g., email or push)
// 	return nil
// }
