import type { Interceptor } from "@connectrpc/connect";

const USER_ID_HEADER = "user-id";

/**
 * Connect interceptor that ensures Link Service always receives a user-id metadata.
 * Link Service expects "user-id" in metadata (session_interceptor checks for it).
 * Link Service uses this to check private link access via Kratos Admin API.
 *
 * This interceptor sets a default value (serviceUserId) if user-id is not already set
 * in the request headers. The actual userId from Kratos session is set via callOptions.header
 * in LinkServiceConnectAdapter.getLinkByHash().
 *
 * @param serviceUserId - stable identifier for proxy service account (fallback).
 */
export function createSessionInterceptor(serviceUserId: string): Interceptor {
  if (!serviceUserId || !serviceUserId.trim()) {
    throw new Error(
      "SERVICE_USER_ID is required to call Link Service. Provide non-empty value."
    );
  }

  const normalizedServiceUserId = serviceUserId.trim();

  return (next) => async (req) => {
    // Set default user-id if not already set (from callOptions.header)
    // This ensures backward compatibility and provides a fallback
    if (!req.header.get(USER_ID_HEADER)) {
      req.header.set(USER_ID_HEADER, normalizedServiceUserId);
    }

    return next(req);
  };
}
