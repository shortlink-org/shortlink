import type { StorybookConfig } from '@storybook/experimental-nextjs-vite'

const config: StorybookConfig = {
  stories: ['../src/**/*.stories.mdx', '../src/**/*.stories.@(js|jsx|mjs|ts|tsx)'],

  addons: [
    '@storybook/addon-links',
    '@storybook/addon-essentials',
    '@storybook/addon-themes',
    '@storybook/addon-interactions',
    '@storybook/addon-controls',
    '@chromatic-com/storybook',
    '@storybook/addon-a11y',
    '@storybook/addon-coverage',
    '@storybook/addon-jest',
  ],

  core: {
    builder: '@storybook/builder-vite',
  },

  framework: {
    name: '@storybook/experimental-nextjs-vite',
    options: {}
  },

  features: {},
  typescript: {},

  docs: {}
}

export default config
