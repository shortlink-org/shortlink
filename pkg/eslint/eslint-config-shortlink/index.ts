const js = require('@eslint/js')
const tsParser = require('@typescript-eslint/parser')
const eslintPluginPrettierRecommended = require('eslint-plugin-prettier/recommended')
const react = require('eslint-plugin-react')
const shortlink = require('eslint-plugin-shortlink')
const globals = require('globals')
import stylisticJs from '@stylistic/eslint-plugin-js'

const config = [
  js.configs.recommended,
  eslintPluginPrettierRecommended,
  {
    files: ['**/*.{js,jsx,mjs,cjs,ts,tsx}'],
    languageOptions: {
      parser: tsParser,
      ecmaVersion: 'latest',
      sourceType: 'module',
      parserOptions: {
        ecmaFeatures: {
          jsx: true,
        },
      },
      globals: {
        ...globals.serviceworker,
        ...globals.browser,
      },
    },
    plugins: {
      react,
      '@stylistic/js': stylisticJs,
      shortlink,
    },
    rules: {
      'no-undef': 'off',
      'no-unused-vars': 'off',
      'no-fallthrough': 'off',
      semi: ['off', 'never'],
      'global-require': 'off',
      'no-console': ['warn', { allow: ['info', 'warn', 'error'] }],
      'no-shadow': 'off',
      camelcase: 'off',
      'react/jsx-pascal-case': 'off',
      'react-hooks/exhaustive-deps': 'off',
      'import/no-default-export': 'off',
    },
    ignores: ['.storybook/'],
  },
]

module.exports = config
