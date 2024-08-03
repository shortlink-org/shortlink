package v1

import (
	"errors"
	"testing"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/bufbuild/protovalidate-go"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddRequestToDomain(t *testing.T) {
	validUUID := uuid.New().String()
	invalidUUID := "invalid-uuid"

	// Initialize the validator
	validator, err := protovalidate.New()
	require.NoError(t, err)

	tests := []struct {
		name          string
		request       *AddRequest
		expectedError error
	}{
		{
			name: "valid request",
			request: &AddRequest{
				CustomerId: validUUID,
				Items: map[string]int32{
					uuid.New().String(): 1,
				},
			},
			expectedError: nil,
		},
		{
			name: "invalid customer id",
			request: &AddRequest{
				CustomerId: invalidUUID,
				Items: map[string]int32{
					uuid.New().String(): 1,
				},
			},
			expectedError: ErrInvalidCustomerId,
		},
		{
			name: "invalid product id",
			request: &AddRequest{
				CustomerId: validUUID,
				Items: map[string]int32{
					invalidUUID: 1,
				},
			},
			expectedError: ParseItemError{Err: errors.New(""), item: invalidUUID},
		},
		{
			name: "negative quantity",
			request: &AddRequest{
				CustomerId: validUUID,
				Items: map[string]int32{
					"0ab35561-f7d9-4f70-967c-9b104e06866d": -1,
				},
			},
			expectedError: &protovalidate.ValidationError{
				Violations: []*validate.Violation{
					{
						FieldPath:    `items["0ab35561-f7d9-4f70-967c-9b104e06866d"]`,
						ConstraintId: "int32.gt",
						Message:      "value must be greater than 0",
						ForKey:       false,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cartState, err := AddRequestToDomain(tt.request, validator)

			// Check for error
			assert.Equal(t, tt.expectedError != nil, err != nil)

			// Check error type and content
			switch expectedErr := tt.expectedError.(type) {
			case nil:
				require.NoError(t, err)
				assert.NotNil(t, cartState)
				assert.Equal(t, tt.request.CustomerId, cartState.GetCustomerId().String())
			case ParseItemError:
				assert.IsType(t, expectedErr, err)
				assert.Equal(t, expectedErr.item, err.(ParseItemError).item)
			default:
				assert.EqualError(t, err, expectedErr.Error())
			}
		})
	}
}
