import { ConnectError, Code } from "@connectrpc/connect";
import { ExternalServiceError } from "../../../../application/exceptions/index.js";
import { LinkNotFoundError } from "../../../../domain/exceptions/index.js";
import { getUserFriendlyMessage } from "./errorUtils.js";
import { Hash } from "../../../../domain/entities/Hash.js";

/**
 * Maps Connect errors to domain errors.
 *
 * Business logic:
 * - NotFound/PermissionDenied -> LinkNotFoundError (according to ADR-42, hide private links)
 * - Infrastructure errors -> ExternalServiceError (user-friendly message)
 * - Other errors -> re-thrown as-is
 *
 * Usage in adapter:
 * ```typescript
 * try {
 *   const res = await this.client.get(...);
 *   ...
 * } catch (err) {
 *   this.errorMapper.map(err, hash);
 * }
 * ```
 */
export class ConnectErrorMapper {
  /**
   * Service name for ExternalServiceError (e.g., "link-service")
   */
  constructor(private readonly serviceName: string) {}

  /**
   * Maps ConnectError to domain error or throws appropriate error.
   *
   * @param error - Error to map
   * @param hash - Hash of the requested link (for LinkNotFoundError)
   * @throws LinkNotFoundError for NotFound/PermissionDenied
   * @throws ExternalServiceError for infrastructure errors
   * @throws original error for other errors
   */
  map(error: unknown, hash: Hash): never {
    // If not a ConnectError, re-throw as-is
    if (!(error instanceof ConnectError)) {
      throw error;
    }

    // Handle ConnectError codes using switch for clarity and extensibility
    switch (error.code) {
      case Code.NotFound:
      case Code.PermissionDenied:
        // According to ADR-42: hide existence of private links
        throw new LinkNotFoundError(hash);

      case Code.Unavailable:
      case Code.DeadlineExceeded:
      case Code.ResourceExhausted:
      case Code.Aborted:
      case Code.Internal:
        // Infrastructure errors: wrap with user-friendly message
        const userMessage = getUserFriendlyMessage(error);
        throw new ExternalServiceError(
          userMessage,
          this.serviceName,
          503,
          error as Error
        );

      default:
        // All other errors: re-throw as-is
        // Interceptors have already handled logging, metrics, and tracing
        throw error;
    }
  }
}
