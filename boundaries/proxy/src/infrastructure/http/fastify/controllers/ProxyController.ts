import type { FastifyRequest, FastifyReply } from "fastify";
import { LinkApplicationService } from "../../../../application/services/LinkApplicationService.js";
import { Hash } from "../../../../domain/entities/Hash.js";
import type { ILogger } from "../../../logging/ILogger.js";

/**
 * HTTP Controller for redirect handling.
 * Thin adapter that converts HTTP requests to Use Case calls.
 * No decorators - plain class with explicit dependency injection.
 */
export class ProxyController {
  constructor(
    private readonly linkApplicationService: LinkApplicationService,
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

      // Call application service
      const result = await this.linkApplicationService.handleRedirect({ hash });

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

