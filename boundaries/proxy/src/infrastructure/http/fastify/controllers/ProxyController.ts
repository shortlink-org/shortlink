import type { FastifyRequest, FastifyReply } from "fastify";
import { LinkApplicationService } from "../../../../application/services/LinkApplicationService.js";
import { Hash } from "../../../../domain/entities/Hash.js";
import type { ILogger } from "../../../logging/ILogger.js";
import { KratosSessionExtractor } from "../../../auth/index.js";

/**
 * HTTP Controller for redirect handling.
 * Thin adapter that converts HTTP requests to Use Case calls.
 * No decorators - plain class with explicit dependency injection.
 */
export class ProxyController {
  constructor(
    private readonly linkApplicationService: LinkApplicationService,
    private readonly kratosSessionExtractor: KratosSessionExtractor,
    private readonly logger: ILogger
  ) {}

  /**
   * Handles redirect request
   * GET /s/:hash
   *
   * @param request - Fastify request with hash parameter
   * @param reply - Fastify reply object
   */
  async redirect(
    request: FastifyRequest<{
      Params: { hash: string };
    }>,
    reply: FastifyReply
  ): Promise<void> {
    try {
      // Extract and validate hash from params
      // Validation is already done in route preValidation hook
      const hash = new Hash(request.params.hash);

      // Extract Kratos session to get user_id for private link access
      // According to ADR 42: if no valid session, pass "anonymous"
      const session = await this.kratosSessionExtractor.extractSession(request);
      const userId = session.isAuthenticated && session.userId
        ? session.userId
        : "anonymous";

      // Call application service with user_id
      const result = await this.linkApplicationService.handleRedirect({
        hash,
        userId,
      });

      if (result.isErr()) {
        // Error is handled by error handler middleware
        throw result.error;
      }

      const { link } = result.value;

      // Perform redirect (301 Moved Permanently)
      // Fastify redirect API: reply.code(statusCode).redirect(url)
      return reply.code(301).redirect(link.url);
    } catch (error) {
      // Re-throw error to be handled by error handler middleware
      throw error;
    }
  }
}
