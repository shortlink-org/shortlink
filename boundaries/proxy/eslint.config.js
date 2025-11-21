import js from "@eslint/js";
import tseslint from "typescript-eslint";
import importPlugin from "eslint-plugin-import";
import boundariesPlugin from "eslint-plugin-boundaries";
import globals from "globals";

export default [
  js.configs.recommended,
  ...tseslint.configs.recommended,
  ...tseslint.configs["recommended-type-checked"],
  ...tseslint.configs["stylistic-type-checked"],
  {
    files: ["src/**/*.ts"],
    languageOptions: {
      parserOptions: {
        ecmaVersion: 2022,
        sourceType: "module",
        project: "./tsconfig.json",
      },
      globals: {
        ...globals.node,
        ...globals.es2021,
      },
    },
    plugins: {
      import: importPlugin,
      boundaries: boundariesPlugin,
    },
    settings: {
      "import/resolver": {
        typescript: {
          alwaysTryTypes: true,
          project: "./tsconfig.json",
        },
      },
      // Clean Architecture boundaries configuration
      "boundaries/include": ["src"],
      "boundaries/elements": [
        {
          type: "domain",
          pattern: "src/proxy/domain/**",
        },
        {
          type: "application",
          pattern: "src/proxy/application/**",
        },
        {
          type: "infrastructure",
          pattern: "src/proxy/infrastructure/**",
        },
        {
          type: "types",
          pattern: "src/proxy/types/**",
        },
        {
          type: "interfaces",
          pattern: "src/interfaces/**",
        },
        {
          type: "container",
          pattern: "src/container/**",
        },
        {
          type: "shared-infrastructure",
          pattern: "src/infrastructure/**",
        },
      ],
    },
    rules: {
      "no-undef": "off", // TypeScript проверяет это
      "@typescript-eslint/no-explicit-any": "warn", // Предупреждение вместо ошибки для any
      "@typescript-eslint/no-require-imports": "off", // Разрешаем require для динамических импортов
      // Bundler moduleResolution - TypeScript и bundler сами проверяют импорты
      "import/no-unresolved": "off", // Не понимает bundler-style resolution (path aliases, .js для .ts и т.д.)
      "import/extensions": "off", // TypeScript ESM требует .js в импортах, но это конфликтует с правилом
      "import/no-extraneous-dependencies": [
        "error",
        {
          devDependencies: [
            "**/*.test.ts",
            "**/*.spec.ts",
            "**/tests/**",
            "**/vitest.config.ts",
            "**/eslint.config.js",
          ],
        },
      ], // Проверяет, что зависимости используются правильно
      // Clean Architecture boundaries enforcement
      "boundaries/element-types": [
        "error",
        {
          default: "disallow",
          rules: [
            {
              from: ["domain"],
              allow: ["types", "interfaces", "shared-infrastructure"], // Domain can import types, interfaces, and shared infrastructure
            },
            {
              from: ["application"],
              allow: ["domain", "types", "interfaces", "shared-infrastructure"], // Application can import from domain, types, interfaces, and shared infrastructure
            },
            {
              from: ["infrastructure"],
              allow: [
                "domain",
                "application",
                "types",
                "interfaces",
                "shared-infrastructure",
              ], // Infrastructure can import from domain, application, types, interfaces, and shared infrastructure
            },
            {
              from: ["types"],
              allow: [], // Types are standalone
            },
            {
              from: ["interfaces"],
              allow: [], // Interfaces are standalone
            },
            {
              from: ["container"],
              allow: [
                "domain",
                "application",
                "infrastructure",
                "types",
                "interfaces",
                "shared-infrastructure",
              ], // Container can import from all layers
            },
            {
              from: ["shared-infrastructure"],
              allow: [], // Shared infrastructure is standalone
            },
          ],
        },
      ],
      "boundaries/no-unknown": "off", // Отключено: требует типизацию всех файлов, но многие директории (test helpers, environment и т.д.) не описаны
      // Block Express for all layers (external dependency, not a CA layer)
      "import/no-restricted-paths": [
        "error",
        {
          zones: [
            {
              target: "./src/proxy/**/*",
              from: ["express", "**/express/**"],
              message: "Express is not allowed in Clean Architecture",
            },
          ],
        },
      ],
    },
  },
  {
    // Config files should be linted
    files: ["*.config.js", "*.config.ts"],
    languageOptions: {
      parserOptions: {
        ecmaVersion: 2022,
        sourceType: "module",
      },
      globals: {
        ...globals.node,
      },
    },
    rules: {
      "@typescript-eslint/no-require-imports": "off",
    },
  },
  {
    ignores: ["dist/", "node_modules/", "src/**/*.js"], // Игнорируем только JS в src, конфиги в корне должны линтиться
  },
];
