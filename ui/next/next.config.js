const webpack = require('webpack')
const withSourceMaps = require('@zeit/next-source-maps')
const isProd = process.env.NODE_ENV === 'production'

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
          destination: `http://localhost:7070/api`,
          basePath: false,
        },
        {
          source: `/api/:uri`,
          destination: `http://localhost:7070/api/:uri`,
          basePath: false,
        },
      ],
    }
  }
}

module.exports = withSourceMaps(NEXT_CONFIG)
