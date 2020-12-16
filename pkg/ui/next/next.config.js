const webpack = require('webpack')
const withSourceMaps = require('@zeit/next-source-maps')

module.exports = withSourceMaps({
  env: {
    SENTRY_DSN: process.env.SENTRY_DSN,
    API_URL_HTTP: "http://localhost:7070"
  },
  webpack: (config, { isServer, buildId }) => {
    config.plugins.push(
      new webpack.DefinePlugin({
        'process.env.SENTRY_RELEASE': JSON.stringify(buildId),
      })
    )

    if (!isServer) {
      config.resolve.alias['@sentry/node'] = '@sentry/browser'
    }

    return config
  },
  async rewrites() {
    return [
      // we need to define a no-op rewrite to trigger checking
      // all pages/static files before we attempt proxying
      {
        source: `/api/:uri`,
        destination: `http://localhost:7070/api/:uri`,
      },
    ]
  },
})
