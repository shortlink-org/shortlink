import type { FastifyRequest } from "fastify";
import { Configuration, FrontendApi } from "@ory/kratos-client";
import { ILogger } from "../logging/ILogger.js";

/**
 * User session information extracted from Kratos
 */
export interface UserSession {
  userId: string;
  email: string | null;
  isAuthenticated: boolean;
}

/**
 * Extracts and validates Ory Kratos session from HTTP requests
 * Used to identify authenticated users for private link access
 */
export class KratosSessionExtractor {
  private readonly kratosClient: FrontendApi;

  constructor(kratosPublicUrl: string, private readonly logger: ILogger) {
    const config = new Configuration({
      basePath: kratosPublicUrl,
    });
    this.kratosClient = new FrontendApi(config);
  }

  /**
   * Extracts user session from Kratos cookie in the request
   * @param request - Fastify request object
   * @returns User session information or unauthenticated session
   */
  async extractSession(request: FastifyRequest): Promise<UserSession> {
    // Extract session cookie
    const sessionCookie = request.cookies?.["ory_kratos_session"];

    if (!sessionCookie) {
      return {
        userId: "",
        email: null,
        isAuthenticated: false,
      };
    }

    try {
      // Validate session with Kratos
      const { data: session } = await this.kratosClient.toSession({
        cookie: `ory_kratos_session=${sessionCookie}`,
      });

      if (!session || !session.identity) {
        return {
          userId: "",
          email: null,
          isAuthenticated: false,
        };
      }

      // Extract email from identity traits
      const email = session.identity.traits?.email as string | undefined;

      return {
        userId: session.identity.id,
        email: email || null,
        isAuthenticated: true,
      };
    } catch (error) {
      // Session invalid or expired - log but don't throw
      this.logger.debug("Kratos session validation failed", {
        error: error instanceof Error ? error.message : String(error),
      });

      return {
        userId: "",
        email: null,
        isAuthenticated: false,
      };
    }
  }
}
