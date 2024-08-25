// ENVIRONMENT VARIABLE ================================================================================================
const isProd = process.env.NODE_ENV === 'production'

const API_URI = process.env.API_URI || 'http://127.0.0.1:7070'

console.info('API_URI', API_URI)

/** @type {import('next').NextConfig} */
module.exports = {
  reactStrictMode: true,
  env: {
    // ShortLink API
    NEXT_PUBLIC_SERVICE_NAME: 'shortlink-shop-ui',
    NEXT_PUBLIC_GIT_TAG: process.env.CI_COMMIT_TAG,
  },
  generateEtags: isProd,
  compiler: {},
  images: {
    remotePatterns: [
      {
        protocol: 'https',
        hostname: 'picsum.photos',
      },
    ]
  },
  trailingSlash: false,
  logging: {
    fetches: {
      fullUrl: true,
    },
  },
  experimental: {
    ppr: 'incremental',
    forceSwcTransforms: true,
    swcTraceProfiling: true,
    instrumentationHook: false,
    webVitalsAttribution: ['CLS', 'FCP', 'FID', 'INP', 'LCP', 'TTFB'],
    turbo: {},
    // typedRoutes: true,
  },
};
