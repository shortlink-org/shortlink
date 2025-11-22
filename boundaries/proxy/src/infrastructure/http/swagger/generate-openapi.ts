import { writeFileSync } from "fs";
import { join } from "path";
import { swaggerSpec } from "./swagger.config.js";

/**
 * Генерирует OpenAPI спецификацию в JSON файл
 * Используется для генерации документации без запуска сервера
 */
export function generateOpenApiSpec(outputPath?: string): void {
  const defaultPath = join(process.cwd(), "openapi.json");
  const filePath = outputPath || defaultPath;

  writeFileSync(filePath, JSON.stringify(swaggerSpec, null, 2), "utf-8");
  console.log(`OpenAPI спецификация сохранена в: ${filePath}`);
}

// Если запущен напрямую через tsx, генерируем спецификацию
if (process.argv[1] && process.argv[1].includes("generate-openapi")) {
  generateOpenApiSpec();
}

