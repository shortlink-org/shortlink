module.exports = {
  stories: [
    "../src/**/*.stories.mdx",
    "../src/**/*.stories.@(js|jsx|ts|tsx)"
  ],
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
        toolbars: true,
      },
    },
    "@storybook/addon-postcss",
    "@storybook/addon-interactions"
  ],
  framework: '@storybook/react',
  core: {
    builder: 'webpack5',
  },
  features: {
    interactionsDebugger: true,
  },
}
