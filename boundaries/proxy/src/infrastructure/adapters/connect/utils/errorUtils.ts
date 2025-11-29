import { ConnectError, Code } from "@connectrpc/connect";

/**
 * Domain of error semantics:
 * - Infrastructure errors: networking, timeouts, overload, aborted requests.
 * - Retryable errors: transient failures suitable for exponential backoff.
 * - User-facing messages: a small dictionary translating gRPC codes into gentle human language.
 */

export const INFRASTRUCTURE_ERROR_CODES = new Set<Code>([
  Code.Unavailable,
  Code.DeadlineExceeded,
  Code.ResourceExhausted,
  Code.Aborted,
  Code.Internal,
]);

export const RETRYABLE_ERROR_CODES = new Set<Code>([
  Code.Unavailable,
  Code.DeadlineExceeded,
  Code.ResourceExhausted,
  Code.Aborted,
]);

/**
 * Mapping from error codes to human-readable UI messages.
 * Designed to be minimal, neutral, emotionally calm.
 */
const USER_FRIENDLY_MESSAGES: Partial<Record<Code, string>> = {
  [Code.Unavailable]:
    "Service temporarily unavailable. Please try again later.",
  [Code.DeadlineExceeded]: "Request timeout. Please try again later.",
  [Code.ResourceExhausted]: "Service is overloaded. Please try again later.",
  [Code.Aborted]: "Request was aborted. Please try again.",
  [Code.Internal]: "Service temporarily unavailable. Please try again later.",
};

/**
 * Type guards and helpers
 */

export function isInfrastructureError(error: unknown): error is ConnectError {
  return (
    error instanceof ConnectError && INFRASTRUCTURE_ERROR_CODES.has(error.code)
  );
}

export function isRetryableError(error: unknown): boolean {
  return error instanceof ConnectError && RETRYABLE_ERROR_CODES.has(error.code);
}

/**
 * Resolves a ConnectError to a graceful, non-technical message.
 * Designed for UI surfaces where internal codes should remain hidden.
 */
export function getUserFriendlyMessage(error: ConnectError): string {
  return (
    USER_FRIENDLY_MESSAGES[error.code] ??
    "Service temporarily unavailable. Please try again later."
  );
}
