package link

import (
	"context"
	"log/slog"

	"github.com/shortlink-org/go-sdk/auth/session"
	"github.com/shortlink-org/go-sdk/saga"

	domain "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

// Get - get a link by hash
//
// According to ADR-42: SpiceDB is not used in GET (redirect use case).
// For public links: return immediately without permission check.
// For private links: check allowlist via email from Kratos Admin API.
//
// Saga:
// 1. Get a link from store
// 2. If public - return immediately
// 3. If private - check permission (via allowlist)
func (uc *UC) Get(ctx context.Context, hash string) (*domain.Link, error) {
	const (
		SAGA_NAME                = "GET_LINK"
		SAGA_STEP_GET_FROM_STORE = "SAGA_STEP_GET_FROM_STORE"
		SAGA_STEP_CHECK_ACCESS   = "SAGA_STEP_CHECK_ACCESS"
	)

	// Get user ID from session metadata (may be empty for anonymous requests)
	// If metadata is missing or empty, userID will be empty string (treated as anonymous)
	userID, err := session.GetUserID(ctx)
	if err != nil {
		// If GetUserID returns error (metadata missing), treat as anonymous request
		// This is normal for unauthenticated users accessing public links
		userID = ""
	}

	var resp *domain.Link

	// create a new saga for a get link by hash
	sagaGetLink, errs := saga.New(SAGA_NAME, saga.SetLogger(uc.log)).
		WithContext(ctx).
		Build()
	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, err
	}

	// Step 1: Get link from store first
	_, errs = sagaGetLink.AddStep(SAGA_STEP_GET_FROM_STORE).
		Then(func(ctx context.Context) error {
			var err error

			resp, err = uc.store.Get(ctx, hash)
			if err != nil {
				return err
			}

			return nil
		}).Reject(func(ctx context.Context, err error) error {
		return err
	}).Build()
	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, err
	}

		// Step 2: Check access (only for private links)
		_, errs = sagaGetLink.AddStep(SAGA_STEP_CHECK_ACCESS).
			Needs(SAGA_STEP_GET_FROM_STORE).
			Then(func(ctx context.Context) error {
				// If link is public, allow access without any checks
				if resp != nil && resp.IsPublic() {
					return nil
				}

				// For private links: check via allowlist
				// According to ADR-42:
				// - If user_id is empty (anonymous request) → ErrPermissionDenied
				// - Get email from Kratos Admin API
				// - Check email against allowlist using link.CanBeViewedByEmail(email)
				
				if userID == "" {
					return domain.ErrPermissionDenied(nil)
				}

				// Get email from Kratos Admin API
				userEmail, err := uc.kratos.GetUserEmail(ctx, userID)
				if err != nil {
					// According to ADR-42: any error should result in permission denied
					// to avoid revealing information about user existence
					uc.log.ErrorWithContext(ctx, "failed to get user email from Kratos",
						slog.String("user_id", userID),
						slog.String("error", err.Error()),
					)

					return domain.ErrPermissionDenied(err)
				}

				// Check if email is in allowlist
				if !resp.CanBeViewedByEmail(userEmail) {
					uc.log.ErrorWithContext(ctx, "email not in allowlist for private link",
						slog.String("user_id", userID),
						slog.String("link_hash", resp.GetHash()),
						// Don't log email for security reasons
					)

					return domain.ErrPermissionDenied(nil)
				}

				return nil
			}).Reject(func(ctx context.Context, thenErr error) error {
			return domain.ErrPermissionDenied(thenErr)
		}).Build()
	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, err
	}

	// Run saga
	err = sagaGetLink.Play(nil)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, domain.ErrNotFound(hash)
	}

	return resp, nil
}
