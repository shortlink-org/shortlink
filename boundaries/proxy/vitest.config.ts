import { defineConfig } from "vitest/config";
import path from "path";
import fs from "fs";

// Функция для создания alias, который разрешает .js → .ts
function createJsToTsAlias() {
  const srcDir = path.resolve(__dirname, "./src");
  const aliases: Record<string, string> = {};
  
  // Рекурсивно находим все .ts файлы в src/
  function findTsFiles(dir: string, basePath: string = "") {
    const entries = fs.readdirSync(dir, { withFileTypes: true });
    
    for (const entry of entries) {
      const fullPath = path.join(dir, entry.name);
      const relativePath = path.join(basePath, entry.name);
      
      if (entry.isDirectory()) {
        findTsFiles(fullPath, relativePath);
      } else if (entry.isFile() && entry.name.endsWith(".ts")) {
        // Создаем alias для .js версии файла
        const jsPath = relativePath.replace(/\.ts$/, ".js");
        aliases[`${jsPath}`] = fullPath;
      }
    }
  }
  
  findTsFiles(srcDir);
  return aliases;
}

export default defineConfig({
  test: {
    globals: true,
    environment: "node",
    setupFiles: ["./src/__tests__/proto-init.ts", "./src/__tests__/setup.ts"],
    include: ["src/**/*.{test,spec}.ts"],
    testTimeout: 30000, // Увеличиваем таймаут для e2e тестов с Testcontainers
    hookTimeout: 120000, // Увеличиваем таймаут для beforeAll/afterAll хуков
    resolve: {
      // Разрешаем импорты с .js расширением к .ts файлам (ESM стандарт TypeScript)
      // Порядок важен: сначала проверяем .ts файлы
      extensions: [".ts", ".tsx", ".js", ".jsx", ".json"],
      alias: {
        // Настройка для path aliases из tsconfig.json
        "@/domain": path.resolve(__dirname, "./src/proxy/domain"),
        "@/application": path.resolve(__dirname, "./src/proxy/application"),
        "@/infrastructure": path.resolve(__dirname, "./src/proxy/infrastructure"),
        // Добавляем динамические aliases для .js → .ts
        ...createJsToTsAlias(),
      },
    },
    server: {
      deps: {
        inline: ["@bufbuild/protobuf"],
      },
    },
    coverage: {
      provider: "v8",
      reporter: ["text", "json", "html"],
      exclude: [
        "node_modules/",
        "src/**/*.d.ts",
        "src/**/__tests__/**",
        "src/**/*.test.ts",
        "src/**/*.spec.ts",
      ],
    },
  },
});

