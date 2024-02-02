module.exports = {
  stories: ['../src/**/*.stories.mdx', '../src/**/*.stories.@(js|jsx|ts|tsx)'],
  addons: [
    '@storybook/addon-links',
    '@storybook/addon-themes',
    {
      name: '@storybook/addon-essentials',
      options: {
        actions: true,
        backgrounds: true,
        controls: true,
        docs: true, // https://github.com/hipstersmoothie/react-docgen-typescript-plugin/issues/83
        viewport: true,
        toolbars: true,
      },
    },
    {
      name: '@storybook/addon-styling-webpack',
      options: {
        postCss: {
          implementation: require.resolve('postcss'),
        },
      },
    },
    '@storybook/addon-interactions',
    '@storybook/addon-controls',
    {
      name: '@storybook/addon-styling-webpack',
      options: {
        rules: [
          {
            test: /\.css$/,
            sideEffects: true,
            use: [
              require.resolve('style-loader'),
              {
                loader: require.resolve('css-loader'),
                options: {
                  // Want to add more CSS Modules options? Read more here: https://github.com/webpack-contrib/css-loader#modules
                  modules: {
                    auto: true,
                  },
                  importLoaders: 1,
                },
              },
              {
                loader: require.resolve('postcss-loader'),
                options: {
                  implementation: require.resolve('postcss'),
                },
              },
            ],
          },
        ],
      },
    },
  ],
  framework: {
    name: '@storybook/react-webpack5',
    options: {
      fsCache: true,
      lazyCompilation: true,
      builder: {
        useSWC: true,
      }
    },
  },
  features: {
    interactionsDebugger: true,
  },
  docs: {},
  typescript: {
    reactDocgen: 'react-docgen',
  },
}
