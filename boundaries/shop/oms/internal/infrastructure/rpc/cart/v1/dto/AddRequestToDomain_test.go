package dto

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	domain "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/cart/v1"
	model "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/cart/v1/model/v1"
)

func TestAddRequestToDomain(t *testing.T) {
	tests := []struct {
		name          string
		request       *model.AddRequest
		expectedError error
		expectedState *domain.CartState
	}{
		{
			name: "Valid AddRequest",
			request: &model.AddRequest{
				CustomerId: "e2c8ba97-1a6b-4c5c-9a2a-3f4c9b9d65a1",
				Items: []*model.CartItem{
					{ProductId: "c5f5d6d6-98e6-4f57-b34a-48a3997f28d4", Quantity: 1},
					{ProductId: "da3f3a3e-784d-4a9a-8cfa-6321d555d6a3", Quantity: 2},
				},
			},
			expectedError: nil,
			expectedState: func() *domain.CartState {
				customerId, _ := uuid.Parse("e2c8ba97-1a6b-4c5c-9a2a-3f4c9b9d65a1")
				productId1, _ := uuid.Parse("c5f5d6d6-98e6-4f57-b34a-48a3997f28d4")
				productId2, _ := uuid.Parse("da3f3a3e-784d-4a9a-8cfa-6321d555d6a3")
				state := domain.NewCartState(customerId)
				state.AddItem(domain.NewCartItem(productId1, 1))
				state.AddItem(domain.NewCartItem(productId2, 2))
				return state
			}(),
		},
		{
			name: "Invalid Customer ID",
			request: &model.AddRequest{
				CustomerId: "invalid-uuid",
				Items: []*model.CartItem{
					{ProductId: uuid.New().String(), Quantity: 1},
				},
			},
			expectedError: ErrInvalidCustomerId,
			expectedState: nil,
		},
		{
			name: "Invalid Product ID",
			request: &model.AddRequest{
				CustomerId: uuid.New().String(),
				Items: []*model.CartItem{
					{ProductId: "invalid-uuid", Quantity: 1},
				},
			},
			expectedError: ParseItemError{
				Err:  errors.New("invalid UUID length: 12"),
				item: "invalid-uuid",
			},
			expectedState: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualState, err := AddRequestToDomain(tt.request)
			if tt.expectedError != nil {
				assert.NotNil(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, actualState)
				assert.Equal(t, tt.expectedState.GetCustomerId(), actualState.GetCustomerId())
				assert.Equal(t, len(tt.expectedState.GetItems()), len(actualState.GetItems()))
			}
		})
	}
}
