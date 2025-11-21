import { defineConfig } from "vitest/config";
import path from "path";

export default defineConfig({
  resolve: {
    // Modern TS + ESM resolution rules (best for Fastify + Testcontainers)
    conditions: ["import", "module", "default"],
    extensions: [".ts", ".tsx", ".js", ".jsx", ".json"],

    // Clean CA/DDD structure aliases
    alias: {
      "@/domain": path.resolve(__dirname, "./src/proxy/domain"),
      "@/application": path.resolve(__dirname, "./src/proxy/application"),
      "@/infrastructure": path.resolve(__dirname, "./src/proxy/infrastructure"),
    },
  },

  test: {
    globals: true,
    environment: "node",

    // 🟢 REQUIRED for Testcontainers stability
    // "forks" = Node.js child_process pool (NOT worker threads)
    // This avoids issues with Docker socket sharing.
    pool: "forks",

    // 🚫 Disable per-test-file isolation
    // Required when integration tests share container-based state
    isolate: false,

    // 🟢 Long timeouts for container startup (safe defaults)
    testTimeout: 60_000,
    hookTimeout: 60_000,

    // Load the protobuf loader + global test setup
    setupFiles: ["./src/__tests__/proto-init.ts", "./src/__tests__/setup.ts"],

    // Unit + Integration tests
    include: ["src/**/*.{test,spec}.ts", "tests/**/*.spec.ts"],

    // Vitest/Vite Edge case for BufBuild protobuf
    server: {
      deps: { inline: ["@bufbuild/protobuf"] },
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
