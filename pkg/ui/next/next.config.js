const webpack = require('webpack')
const withSourceMaps = require('@zeit/next-source-maps')
const {
  PHASE_DEVELOPMENT_SERVER,
  PHASE_PRODUCTION_BUILD,
} = require('next/constants')

const NEXT_CONFIG = {
  basePath: '/next',
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
}

if (PHASE_DEVELOPMENT_SERVER) {
  NEXT_CONFIG.rewrites = async function() {
    return [
      // we need to define a no-op rewrite to trigger checking
      // all pages/static files before we attempt proxying
      {
        source: `/api/:uri`,
        destination: `http://localhost:7070/api/:uri`,
        basePath: false,
      },
    ]
  }
}

module.exports = withSourceMaps(NEXT_CONFIG)
