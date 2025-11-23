import type { Interceptor } from "@connectrpc/connect";

const USER_ID_HEADER = "user-id";

/**
 * Connect interceptor that ensures Link Service always receives a user-id metadata.
 * Link Service enforces `user-id` via go-sdk session interceptor, so proxy must send it.
 *
 * @param serviceUserId - stable identifier for proxy service account.
 */
export function createSessionInterceptor(serviceUserId: string): Interceptor {
  if (!serviceUserId || !serviceUserId.trim()) {
    throw new Error(
      "SERVICE_USER_ID is required to call Link Service. Provide non-empty value."
    );
  }

  const normalizedUserId = serviceUserId.trim();

  return (next) => async (req) => {
    // Setting header ensures go-sdk session interceptor accepts the request.
    req.header.set(USER_ID_HEADER, normalizedUserId);
    return next(req);
  };
}
