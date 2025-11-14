import js from "@eslint/js";
import tseslint from "@typescript-eslint/eslint-plugin";
import tsparser from "@typescript-eslint/parser";
import importPlugin from "eslint-plugin-import";
import globals from "globals";

export default [
  js.configs.recommended,
  {
    files: ["src/**/*.ts"],
    languageOptions: {
      parser: tsparser,
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
      "@typescript-eslint": tseslint,
      import: importPlugin,
    },
    rules: {
      ...tseslint.configs.recommended.rules,
      "no-undef": "off", // TypeScript проверяет это
      "@typescript-eslint/no-explicit-any": "warn", // Предупреждение вместо ошибки для any
      "@typescript-eslint/no-require-imports": "off", // Разрешаем require для динамических импортов
      "import/no-restricted-paths": [
        "error",
        {
          zones: [
            // Domain слой - не может импортировать инфраструктурные зависимости
            {
              target: "./src/proxy/domain/**/*",
              from: "./src/proxy/infrastructure/**/*",
              message: "Domain layer cannot import from infrastructure layer",
            },
            {
              target: "./src/proxy/domain/**/*",
              from: "./src/proxy/interfaces/**/*",
              message: "Domain layer cannot import from interfaces layer",
            },
            {
              target: "./src/proxy/domain/**/*",
              from: "./src/proxy/application/**/*",
              message: "Domain layer cannot import from application layer",
            },
            {
              target: "./src/proxy/domain/**/*",
              from: "./**/express",
              message: "Domain layer cannot import Express",
            },
            // Application слой - не может импортировать interfaces и Express
            {
              target: "./src/proxy/application/**/*",
              from: "./src/proxy/interfaces/**/*",
              message: "Application layer cannot import from interfaces layer",
            },
            {
              target: "./src/proxy/application/**/*",
              from: "./src/interfaces/**/*",
              message: "Application layer cannot import from interfaces layer",
            },
            {
              target: "./src/proxy/application/**/*",
              from: "./**/express",
              message: "Application layer cannot import Express",
            },
            // Infrastructure может импортировать Domain (это разрешено)
            // Но Domain не может импортировать Infrastructure
          ],
        },
      ],
    },
  },
  {
    ignores: ["dist/", "node_modules/", "*.js", "!*.config.js"],
  },
];

