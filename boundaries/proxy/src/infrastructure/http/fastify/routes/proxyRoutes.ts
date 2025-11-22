import type { FastifyInstance, FastifyPluginOptions } from "fastify";
import type { AwilixContainer } from "awilix";
import type { ContainerDependencies } from "../../../../di/container.js";
import { z } from "zod";

/**
 * Zod schema for redirect request parameters
 */
const RedirectParamsSchema = z.object({
  hash: z
    .string()
    .min(1, "Hash cannot be empty")
    .regex(/^[a-zA-Z0-9]+$/, "Hash must contain only alphanumeric characters"),
});

/**
 * Fastify JSON Schema for redirect endpoint (auto-generated from Zod)
 */
const redirectSchema = {
  description: "Redirect to original URL by short link hash",
  tags: ["Redirect"],
  params: {
    type: "object",
    required: ["hash"],
    properties: {
      hash: {
        type: "string",
        pattern: "^[a-zA-Z0-9]+$",
        minLength: 1,
        description: "Short link hash",
        example: "abc123",
      },
    },
  },
  response: {
    301: {
      description: "Successful redirect",
      headers: {
        type: "object",
        properties: {
          Location: {
            type: "string",
            description: "Redirect URL",
          },
        },
      },
    },
    400: {
      description: "Invalid hash format",
      type: "object",
      properties: {
        error: {
          type: "object",
          properties: {
            code: { type: "string" },
            message: { type: "string" },
            field: { type: "string" },
            timestamp: { type: "string" },
          },
        },
      },
    },
    404: {
      description: "Link not found",
      type: "object",
      properties: {
        error: {
          type: "object",
          properties: {
            code: { type: "string" },
            message: { type: "string" },
            timestamp: { type: "string" },
          },
        },
      },
    },
    429: {
      description: "Rate limit exceeded",
    },
    500: {
      description: "Internal server error",
    },
  },
};

/**
 * Registers proxy redirect routes
 */
export async function registerProxyRoutes(
  fastify: FastifyInstance,
  opts: FastifyPluginOptions & { container: AwilixContainer<ContainerDependencies> }
): Promise<void> {
  const container = opts.container;
  const proxyController = container.resolve("proxyController");

  // GET /s/:hash - Redirect to original URL
  fastify.get<{
    Params: { hash: string };
  }>(
    "/s/:hash",
    {
      schema: redirectSchema,
      preValidation: async (request, reply) => {
        // Validate params using Zod
        const result = RedirectParamsSchema.safeParse(request.params);
        if (!result.success) {
          const firstError = result.error.issues[0];
          reply.code(400).send({
            error: {
              code: "VALIDATION_ERROR",
              message: firstError?.message || "Invalid request parameters",
              field: firstError?.path.join(".") || "hash",
              timestamp: new Date().toISOString(),
            },
          });
          return;
        }
      },
    },
    async (request, reply) => {
      // Get controller from request-scoped container
      const requestContainer = (request as any).container as AwilixContainer<ContainerDependencies>;
      const controller = requestContainer.resolve("proxyController");

      await controller.redirect(request, reply);
    }
  );
}

