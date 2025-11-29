import {
  createClient as createPromiseClient,
  type Client as PromiseClient,
  type Transport,
} from "@connectrpc/connect";
import type { DescService } from "@bufbuild/protobuf";
import { createGrpcTransport } from "@connectrpc/connect-node";
import type { Interceptor } from "@connectrpc/connect";
import { LinkServiceACL } from "../anti-corruption/LinkServiceACL.js";
import { ILinkServiceAdapter } from "./ILinkServiceAdapter.js";
import { Link } from "../../domain/entities/Link.js";
import { Hash } from "../../domain/entities/Hash.js";
import { LinkNotFoundError } from "../../domain/exceptions/index.js";
import { ExternalServicesConfig } from "../config/ExternalServicesConfig.js";
import { ILogger } from "../logging/ILogger.js";
import { IGrpcMetrics } from "../metrics/IGrpcMetrics.js";
import { IGrpcTracing } from "../tracing/IGrpcTracing.js";
import { ConnectErrorMapper } from "./connect/utils/ConnectErrorMapper.js";
import {
  createLoggingInterceptor,
  createRetryInterceptor,
  createMetricsInterceptor,
  createTracingInterceptor,
  createSessionInterceptor,
} from "./connect/interceptors/index.js";
import { LinkService } from "@buf/shortlink-org_shortlink-link-link.bufbuild_es/infrastructure/rpc/link/v1/link_rpc_pb.js";

const linkServiceDescriptor = LinkService as unknown as DescService &
  typeof LinkService;

/**
 * Connect adapter for retrieving links from Link Service
 * Uses official ConnectRPC client via createPromiseClient
 */
export class LinkServiceConnectAdapter implements ILinkServiceAdapter {
  private readonly client: PromiseClient<typeof LinkService>;
  private readonly errorMapper: ConnectErrorMapper;

  constructor(
    private readonly linkServiceACL: LinkServiceACL,
    private readonly externalServicesConfig: ExternalServicesConfig,
    private readonly logger: ILogger,
    private readonly grpcMetrics: IGrpcMetrics,
    private readonly grpcTracing: IGrpcTracing
  ) {
    this.errorMapper = new ConnectErrorMapper("link-service");

    // Create interceptors and transport using protected factory methods
    // This allows test classes to override createTransport() for mocking
    const interceptors = this.createInterceptors();
    const transport = this.createTransport(interceptors);

    // Create client using official ConnectRPC pattern
    this.client = createPromiseClient(
      linkServiceDescriptor,
      transport
    ) as PromiseClient<typeof LinkService>;
  }

  /**
   * Creates Connect interceptors for the transport.
   * Protected method allows test classes to override if needed.
   *
   * @returns Array of interceptors in execution order
   */
  protected createInterceptors(): Interceptor[] {
    // Retry logic is handled at transport level via interceptor
    return [
      // Order matters: session metadata first so downstream interceptors/transport see it
      createSessionInterceptor(this.externalServicesConfig.serviceUserId),
      // Tracing/metrics order preserved after metadata enrichment
      createTracingInterceptor(this.grpcTracing),
      createMetricsInterceptor(this.grpcMetrics),
      createRetryInterceptor(this.logger, {
        maxAttempts: this.externalServicesConfig.retryCount,
        initialDelayMs: 100,
        maxDelayMs: 5000,
        backoffMultiplier: 2,
      }),
      createLoggingInterceptor(this.logger),
    ];
  }

  /**
   * Creates gRPC transport with interceptors.
   * Protected method allows test classes to override and provide mock transport.
   *
   * @param interceptors - Array of interceptors to use
   * @returns Transport instance
   */
  protected createTransport(interceptors: Interceptor[]): Transport {
    return createGrpcTransport({
      baseUrl: `http://${this.externalServicesConfig.linkServiceGrpcUrl}`,
      interceptors,
    });
  }

  /**
   * Builds call options for ConnectRPC request.
   * Encapsulates timeout and header configuration logic.
   *
   * @param userId - Optional user ID for authorization header (from Kratos session)
   * @returns Call options with timeout signal and headers
   */
  private buildCallOptions(userId?: string | null): {
    signal: AbortSignal;
    header?: Record<string, string>;
  } {
    const options: {
      signal: AbortSignal;
      header?: Record<string, string>;
    } = {
      // Use AbortSignal.timeout for request cancellation
      signal: AbortSignal.timeout(this.externalServicesConfig.requestTimeout),
    };

    // According to ADR 42: pass user_id via x-user-id header
    // If userId is provided, use it; otherwise interceptor will use serviceUserId or "anonymous"
    if (userId) {
      options.header = {
        "x-user-id": userId === "anonymous" ? "anonymous" : userId,
      };
    }

    return options;
  }

  async getLinkByHash(hash: Hash, userId?: string | null): Promise<Link> {
    const hashValue = hash.value;

    try {
      const res = await this.client.get(
        { hash: hashValue },
        this.buildCallOptions(userId)
      );

      if (!res || !res.link) {
        // Successful response with empty link (unexpected server behavior)
        // Throw domain error instead of returning null
        throw new LinkNotFoundError(hash);
      }

      // Convert protobuf Link to domain entity via ACL
      const domainLink = this.linkServiceACL.toDomainEntityFromProto(res.link);

      return domainLink;
    } catch (error: unknown) {
      // Map Connect errors to domain errors
      // Interceptors have already handled logging, metrics, tracing, and retry
      this.errorMapper.map(error, hash);
      // Unreachable, but TypeScript needs this
      throw error;
    }
  }
}
