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
})
