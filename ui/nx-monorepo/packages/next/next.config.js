/* eslint-disable */
// const withPWA = require('@ducanh2912/next-pwa').default({
//   dest: 'public',
//   maximumFileSizeToCacheInBytes: 4000000,
//   swcMinify: true,
//   cacheOnFrontendNav: true,
//   aggressiveFrontEndNavCaching: true,
// })
const { composePlugins } = require('@nx/next')
const { withSentryConfig } = require('@sentry/nextjs')
const path = require('path')

// PLUGINS =============================================================================================================
// const plugins = [withPWA]
const plugins = []

// ENVIRONMENT VARIABLE ================================================================================================
const isProd = process.env.NODE_ENV === 'production'
const isEnableSentry = process.env.SENTRY_ENABLE === 'true'
const API_URI = process.env.API_URI || 'http://127.0.0.1:7070'
const AUTH_URI = process.env.AUTH_URI || 'http://127.0.0.1:4433'

console.info('API_URI', API_URI)
console.info('NODE_ENV', process.env.NODE_ENV)

/** @type {import('@nx/next/plugins/with-nx').WithNxOptions} * */
let NEXT_CONFIG = {
  basePath: '/next',
  generateEtags: true,
  env: {
    // ShortLink API
    NEXT_PUBLIC_SERVICE_NAME: 'shortlink-next',
    NEXT_PUBLIC_API_URI: process.env.API_URI,

    // Sentry
    NEXT_PUBLIC_SENTRY_DSN: process.env.SENTRY_DSN,

    // Firebase
    NEXT_PUBLIC_FIREBASE_VAPID_KEY: process.env.FIREBASE_VAPID_KEY,
    NEXT_PUBLIC_FIREBASE_API_KEY: process.env.FIREBASE_API_KEY,
    NEXT_PUBLIC_FIREBASE_AUTH_DOMAIN: process.env.FIREBASE_AUTH_DOMAIN,
    NEXT_PUBLIC_FIREBASE_PROJECT_ID: process.env.FIREBASE_PROJECT_ID,
    NEXT_PUBLIC_FIREBASE_STORAGE_BUCKET: process.env.FIREBASE_STORAGE_BUCKET,
    NEXT_PUBLIC_FIREBASE_MESSAGING_SENDER_ID:
      process.env.FIREBASE_MESSAGING_SENDER_ID,
    NEXT_PUBLIC_FIREBASE_APP_ID: process.env.FIREBASE_APP_ID,
    NEXT_PUBLIC_FIREBASE_MEASUREMENT_ID: process.env.FIREBASE_MEASUREMENT_ID,
  },
  swcMinify: true,
  productionBrowserSourceMaps: true,
  transpilePackages: ['@shortlink-org/ui-kit'],
  // images: {
  //   loader: 'custom',
  //   domains: ['images.unsplash.com'],
  //   formats: ['image/avif', 'image/webp'],
  //   remotePatterns: [
  //     {
  //       // The `src` property hostname must end with `.example.com`,
  //       // otherwise the API will respond with 400 Bad Request.
  //       protocol: 'https',
  //       hostname: 'images.unsplash.com',
  //     },
  //   ],
  //   dangerouslyAllowSVG: false,
  //   contentDispositionType: 'attachment',
  //   contentSecurityPolicy: "default-src 'self'; script-src 'none'; sandbox;",
  // },
  trailingSlash: false,
  webpack: (config, { isServer, buildId }) => {
    config.module.rules.push({
      test: /\.svg$/i,
      issuer: /\.[jt]sx?$/,
      use: ['@svgr/webpack'],
    })

    // This fixes the invalid hook React error which
    // will occur when multiple versions of React are detected
    // This can happen since a common project is also using Next (which is using React)
    const reactPaths = {
      react: path.join(__dirname, '../../node_modules/react'),
      'react-dom': path.join(__dirname, '../../node_modules/react-dom'),
    }
    config.resolve = {
      ...config.resolve,
      alias: {
        ...config.resolve.alias,
        ...reactPaths,
      },
    }

    return config
  },
  logging: {
    fetches: {
      fullUrl: true,
    },
  },
  experimental: {
    forceSwcTransforms: true,
    swcTraceProfiling: true,
    instrumentationHook: true,
    turbo: {},
  },
}

if (isProd) {
  NEXT_CONFIG.output = 'export'
  NEXT_CONFIG.compress = true
  NEXT_CONFIG.productionBrowserSourceMaps = true
}

if (!isProd) {
  NEXT_CONFIG.httpAgentOptions = {
    keepAlive: true,
  }

  NEXT_CONFIG.rewrites = async function () {
    return {
      beforeFiles: [
        // we need to define a no-op rewrite to trigger checking
        // all pages/static files before we attempt proxying
        {
          source: `/api/auth/:uri*`,
          destination: `${AUTH_URI}/:uri*`,
          basePath: false,
        },
        {
          source: `/api/:uri*`,
          destination: `${API_URI}/api/:uri*`,
          basePath: false,
        },
        {
          source: `/s`,
          destination: `${API_URI}/s`,
          basePath: false,
        },
        {
          source: `/s/:uri`,
          destination: `${API_URI}/s/:uri`,
          basePath: false,
        },
      ],
    }
  }
}

// Make sure adding Sentry options is the last code to run before exporting, to
// ensure that your source maps include changes from all other Webpack plugins
if (isEnableSentry) {
  const SentryWebpackPluginOptions = {
    // For all available options, see:
    // https://github.com/getsentry/sentry-webpack-plugin#options

    // Suppresses source map uploading logs during build
    silent: true,

    org: 'batazor',
    project: 'shortlink-next',
  }

  NEXT_CONFIG = withSentryConfig(NEXT_CONFIG, SentryWebpackPluginOptions)
}

module.exports = composePlugins(...plugins)(NEXT_CONFIG)
