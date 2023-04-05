import remarkGfm from 'remark-gfm'

module.exports = {
  stories: ["../src/**/*.stories.mdx", "../src/**/*.stories.@(js|jsx|ts|tsx)"],
  addons: [
    "@storybook/addon-links",
    {
    name: '@storybook/addon-essentials',
      options: {
        actions: true,
        backgrounds: true,
        controls: false,
        docs: false, // https://github.com/hipstersmoothie/react-docgen-typescript-plugin/issues/83
        viewport: true,
        toolbars: true
      }
    },
    {
      name: '@storybook/addon-docs',
      options: {
        mdxPluginOptions: {
          mdxCompileOptions: {
            remarkPlugins: [remarkGfm],
          },
        },
      },
    },
    '@storybook/addon-styling',
    "@storybook/addon-interactions",
  ],
  framework: {
    name: "@storybook/react-webpack5",
    options: {}
  },
  features: {
    interactionsDebugger: true
  },
  docs: {
    autodocs: 'tag',
  }
};
