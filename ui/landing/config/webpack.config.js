const path = require('path');

const config = require('./site.config');
const loaders = require('./webpack.loaders');
const plugins = require('./webpack.plugins');

const CONFIG = {
  context: path.join(config.root, config.paths.src),
  entry: [
    path.join(config.root, config.paths.src, 'javascripts/scripts.js'),
    path.join(config.root, config.paths.src, 'stylesheets/styles.scss'),
  ],
  output: {
    path: path.join(config.root, config.paths.dist),
    filename: '[name].[fullhash].js',
  },
  mode: ['production', 'development'].includes(config.env)
    ? config.env
    : 'development',
  devtool: config.env === 'production'
    ? 'hidden-source-map'
    : 'eval-source-map',
  devServer: {
    contentBase: path.join(config.root, config.paths.src),
    watchContentBase: true,
    hot: true,
    overlay: true,
    open: true,
    port: config.port,
    host: config.dev_host,
  },
  module: {
    rules: loaders,
  },
  plugins,
  experiments: {
    asset: true
  },
};

if (process.env.NODE_ENV !== "development") {
  CONFIG.output.publicPath = '/landing/'
}

module.exports = CONFIG
