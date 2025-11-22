import swaggerJsdoc from "swagger-jsdoc";
import { AppConfig } from "../config/AppConfig.js";

const appConfig = new AppConfig();

/**
 * Конфигурация Swagger/OpenAPI
 * Генерирует OpenAPI спецификацию из JSDoc комментариев
 */
const swaggerOptions: swaggerJsdoc.Options = {
  definition: {
    openapi: "3.0.0",
    info: {
      title: "Proxy Service API",
      version: "1.0.0",
      description: "API для редиректа коротких ссылок на оригинальные URL",
      contact: {
        name: "API Support",
      },
      license: {
        name: "MIT",
      },
    },
    servers: [
      {
        url: `http://localhost:${appConfig.port}`,
        description: "Development server",
      },
      {
        url: "https://api.example.com",
        description: "Production server",
      },
    ],
    tags: [
      {
        name: "Redirect",
        description: "Операции для редиректа коротких ссылок",
      },
      {
        name: "Health",
        description: "Health check endpoints",
      },
    ],
    components: {
      schemas: {
        Error: {
          type: "object",
          properties: {
            code: {
              type: "string",
              description: "Код ошибки",
              example: "LINK_NOT_FOUND",
            },
            message: {
              type: "string",
              description: "Сообщение об ошибке",
              example: "Link with hash 'abc123' not found",
            },
          },
          required: ["code", "message"],
        },
        ValidationError: {
          type: "object",
          properties: {
            code: {
              type: "string",
              description: "Код ошибки валидации",
              example: "VALIDATION_ERROR",
            },
            message: {
              type: "string",
              description: "Сообщение об ошибке валидации",
              example: "Invalid hash format",
            },
            field: {
              type: "string",
              description: "Поле с ошибкой",
              example: "hash",
            },
            details: {
              type: "object",
              description: "Дополнительные детали ошибки",
            },
          },
          required: ["code", "message"],
        },
      },
      responses: {
        NotFound: {
          description: "Ресурс не найден",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/schemas/Error",
              },
            },
          },
        },
        BadRequest: {
          description: "Невалидный запрос",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/schemas/ValidationError",
              },
            },
          },
        },
        TooManyRequests: {
          description:
            "Превышен лимит запросов (обрабатывается на уровне ingress/istio)",
          content: {
            "application/json": {
              schema: {
                type: "object",
                properties: {
                  message: {
                    type: "string",
                    example: "Too many requests",
                  },
                },
              },
            },
          },
        },
        InternalServerError: {
          description: "Внутренняя ошибка сервера",
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/schemas/Error",
              },
            },
          },
        },
      },
    },
  },
  apis: [
    "./src/proxy/interfaces/http/controllers/*.ts",
    "./src/infrastructure/health.ts",
  ],
};

export const swaggerSpec = swaggerJsdoc(swaggerOptions);
