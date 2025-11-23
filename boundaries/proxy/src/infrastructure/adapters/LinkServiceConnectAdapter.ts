import {
  createClient as createPromiseClient,
  ConnectError,
  Code,
  type Client as PromiseClient,
} from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-node";
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
} from "./connect/interceptors/index.js";
import { LinkService } from "../proto/infrastructure/rpc/link/v1/link_connect.js";
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
      // Order matters: tracing first, then metrics, then retry, then logging
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

    // Create Connect transport with interceptors
    const transport = createConnectTransport({
      baseUrl: `http://${this.externalServicesConfig.linkServiceGrpcUrl}`,
      httpVersion: "1.1", // ConnectRPC 2.x requires HTTP/1.1 with Node transport
      interceptors,
    });

    // Create client using official ConnectRPC pattern
    this.client = createPromiseClient(LinkService, transport);
  }

  async getLinkByHash(hash: Hash): Promise<Link | null> {
    const hashValue = hash.value;

    const signal = AbortSignal.timeout(
      this.externalServicesConfig.requestTimeout
    );
    const timeoutMs = this.externalServicesConfig.requestTimeout;

    try {
      const res = await this.client.get(
        { hash: hashValue },
        { signal, timeoutMs }
      );

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
      // Interceptors don't transform or swallow NotFound - they only log/record metrics and propagate the error
      if (error instanceof ConnectError && error.code === Code.NotFound) {
        return null;
      }

      // All other errors are re-thrown
      // Interceptors have already logged and recorded metrics
      throw error;
    }
  }
}
