import {
  createClient as createPromiseClient,
  ConnectError,
  Code,
  type Client as PromiseClient,
} from "@connectrpc/connect";
import type { DescService } from "@bufbuild/protobuf";
import { createGrpcTransport } from "@connectrpc/connect-node";
import { LinkServiceACL } from "../anti-corruption/LinkServiceACL.js";
import { ILinkServiceAdapter } from "./ILinkServiceAdapter.js";
import { Link } from "../../domain/entities/Link.js";
import { Hash } from "../../domain/entities/Hash.js";
import { ExternalServicesConfig } from "../config/ExternalServicesConfig.js";
import { ILogger } from "../logging/ILogger.js";
import { IGrpcMetrics } from "../metrics/IGrpcMetrics.js";
import { IGrpcTracing } from "../tracing/IGrpcTracing.js";
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

  constructor(
    private readonly linkServiceACL: LinkServiceACL,
    private readonly externalServicesConfig: ExternalServicesConfig,
    private readonly logger: ILogger,
    private readonly grpcMetrics: IGrpcMetrics,
    private readonly grpcTracing: IGrpcTracing
  ) {
    // Create Connect interceptors
    const interceptors = [
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

    // Create gRPC transport with interceptors (legacy gRPC server)
    const transport = createGrpcTransport({
      baseUrl: `http://${this.externalServicesConfig.linkServiceGrpcUrl}`,
      interceptors,
    });

    // Create client using official ConnectRPC pattern
    this.client = createPromiseClient(
      linkServiceDescriptor,
      transport
    ) as PromiseClient<typeof LinkService>;
  }

  async getLinkByHash(
    hash: Hash,
    userId?: string | null
  ): Promise<Link | null> {
    const hashValue = hash.value;

    const signal = AbortSignal.timeout(
      this.externalServicesConfig.requestTimeout
    );
    const timeoutMs = this.externalServicesConfig.requestTimeout;

    // Prepare call options
    const callOptions: {
      signal: AbortSignal;
      timeoutMs: number;
      header?: Record<string, string>;
    } = {
      signal,
      timeoutMs,
    };

    // According to ADR 42: pass user_id via x-user-id header
    // If userId is provided, use it; otherwise interceptor will use serviceUserId or "anonymous"
    if (userId) {
      callOptions.header = {
        "x-user-id": userId === "anonymous" ? "anonymous" : userId,
      };
    }

    try {
      const res = await this.client.get({ hash: hashValue }, callOptions);

      if (!res || !res.link) {
        // Successful response with empty link (unexpected server behavior)
        // NOT_FOUND errors are handled in the catch block below
        return null;
      }

      // Convert protobuf Link to domain entity via ACL
      const domainLink = this.linkServiceACL.toDomainEntityFromProto(res.link);

      return domainLink;
    } catch (error: unknown) {
      // Handle Connect errors
      // Interceptors have already handled logging, metrics, and tracing
      // Here we only handle business logic for error processing

      // Check for NOT_FOUND (gRPC status = NotFound)
      if (error instanceof ConnectError && error.code === Code.NotFound) {
        return null;
      }

      // Check for PERMISSION_DENIED (gRPC status = PermissionDenied)
      // According to ADR 42: PermissionDenied should return 404 Not Found
      // We return null here, which will be converted to LinkNotFoundError and then 404
      if (error instanceof ConnectError && error.code === Code.PermissionDenied) {
        return null;
      }

      // All other errors are re-thrown
      // Interceptors have already logged and recorded metrics
      throw error;
    }
  }
}
