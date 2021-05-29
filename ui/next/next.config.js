const webpack = require('webpack')
const withSourceMaps = require('@zeit/next-source-maps')
const isProd = process.env.NODE_ENV === 'production'
const API_URI_PORT = process.env.API_URI_PORT || 7070
const API_URI = process.env.API_URI ? `${process.env.API_URI}:${API_URI_PORT}` : `http://localhost:7070:${API_URI_PORT}`

console.info("API_URI", API_URI)

const NEXT_CONFIG = {
  basePath: '/next',
  env: {},
  webpack: (config, { isServer, buildId }) => {
    config.plugins.push(
      new webpack.DefinePlugin({})
    )

    return config
  },
  future: {
    webpack5: true,
  },
}

if (!isProd) {
  NEXT_CONFIG.rewrites = async function() {
    return {
      fallback: [
        // we need to define a no-op rewrite to trigger checking
        // all pages/static files before we attempt proxying
        {
          source: `/api`,
          destination: `${API_URI}/api`,
          basePath: false,
        },
        {
          source: `/api/:uri`,
          destination: `${API_URI}/api/:uri`,
          basePath: false,
        },
      ],
    }
  }
}

module.exports = withSourceMaps(NEXT_CONFIG)
