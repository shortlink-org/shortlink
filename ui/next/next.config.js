/* eslint-disable */

const webpack = require('webpack')
const { withImageLoader } = require('next-image-loader')
const withExportImages = require('next-export-optimize-images')

// ENVIRONMENT VARIABLE ================================================================================================
const isProd = process.env.NODE_ENV === 'production'
const isEnableSentry = process.env.SENTRY_ENABLE === 'true'
const API_URI = process.env.API_URI || 'http://localhost:7070'

console.info('API_URI', API_URI)

// You can choose which headers to add to the list
// after learning more below.
const securityHeaders = [
  {
    key: 'X-DNS-Prefetch-Control',
    value: 'on',
  },
  {
    key: 'X-XSS-Protection',
    value: '1; mode=block',
  },
  {
    key: 'X-Frame-Options',
    value: 'SAMEORIGIN',
  },
  {
    key: 'Permissions-Policy',
    value: 'camera=(), microphone=(), geolocation=(), interest-cohort=()',
  },
  {
    key: 'X-Content-Type-Options',
    value: 'nosniff',
  },
  {
    key: 'Referrer-Policy',
    value: 'origin-when-cross-origin',
  },
]

if (!isProd) {
  securityHeaders.push({
    key: 'Strict-Transport-Security',
    value: 'max-age=63072000; includeSubDomains; preload',
  })
}

const NEXT_CONFIG = {
  basePath: '/next',
  generateEtags: true,
  env: {
    NEXT_PUBLIC_API_URI: process.env.API_URI,
  },
  swcMinify: true,
  productionBrowserSourceMaps: true,
  transpilePackages: ['@shortlink-org/ui-kit'],
  compiler: {
    // ssr and displayName are configured by default
    styledComponents: true,
  },
  images: {
    loader: 'custom',
    domains: ['images.unsplash.com'],
    formats: ['image/avif', 'image/webp'],
    remotePatterns: [
      {
        // The `src` property hostname must end with `.example.com`,
        // otherwise the API will respond with 400 Bad Request.
        protocol: 'https',
        hostname: 'images.unsplash.com',
      },
    ],
    dangerouslyAllowSVG: true,
    contentSecurityPolicy: "default-src 'self'; script-src 'none'; sandbox;",
  },
  trailingSlash: false,
  headers: () => {
    return [
      {
        // Apply these headers to all routes in your application.
        source: '/:path*',
        headers: securityHeaders,
      },
    ]
  },
  experimental: {
    forceSwcTransforms: true,
    turbo: {
      loaders: {
        '.md': [
          {
            loader: '@mdx-js/loader',
            options: {
              format: 'md',
            },
          },
        ],
        // Option-less format
        '.mdx': '@mdx-js/loader',
        '.svg': '@svgr/webpack',
      }
    }
  },
}

if (!isProd) {
  NEXT_CONFIG.rewrites = async function () {
    return {
      fallback: [
        // we need to define a no-op rewrite to trigger checking
        // all pages/static files before we attempt proxying
        {
          source: `/api/auth/:uri*`,
          destination: `http://127.0.0.1:4433/:uri*`,
          basePath: false,
        },
        {
          source: `/api/:uri*`,
          destination: `http://127.0.0.1:7070/api/:uri*`,
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

let EXPORT_CONFIG = withImageLoader(NEXT_CONFIG)

// Make sure adding Sentry options is the last code to run before exporting, to
// ensure that your source maps include changes from all other Webpack plugins
if (isEnableSentry) {
  // This file sets a custom webpack configuration to use your Next.js app
  // with Sentry.
  // https://nextjs.org/docs/api-reference/next.config.js/introduction
  // https://docs.sentry.io/platforms/javascript/guides/nextjs/
  const { withSentryConfig } = require('@sentry/nextjs')

  const SentryWebpackPluginOptions = {
    // Additional config options for the Sentry Webpack plugin. Keep in mind that
    // the following options are set automatically, and overriding them is not
    // recommended:
    //   release, url, org, project, authToken, configFile, stripPrefix,
    //   urlPrefix, include, ignore
    // For all available options, see:
    // https://github.com/getsentry/sentry-webpack-plugin#options.
  }

  EXPORT_CONFIG = withExportImages(
    withSentryConfig(EXPORT_CONFIG, SentryWebpackPluginOptions),
  )
}

module.exports = EXPORT_CONFIG
