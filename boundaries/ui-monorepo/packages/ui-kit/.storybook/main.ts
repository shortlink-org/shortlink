import type { StorybookConfig } from '@storybook/react-vite'

const config: StorybookConfig = {
  stories: ['../src/**/*.stories.mdx', '../src/**/*.stories.@(js|jsx|ts|tsx)'],
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
  framework: '@storybook/react-vite',
  features: {},
  typescript: {},
}

export default config
