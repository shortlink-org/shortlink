const { composePlugins } = require('@nx/next')

// ENVIRONMENT VARIABLE ================================================================================================
const isProd = process.env.NODE_ENV === 'production'
const isEnablePWA = process.env.PWA_ENABLE === 'true'
const isEnableSentry = process.env.SENTRY_ENABLE === 'true'

const PROXY_URI = process.env.PROXY_URI || 'http://127.0.0.1:3000'
const AUTH_URI = process.env.AUTH_URI || 'http://127.0.0.1:4433'
const API_URI = process.env.API_URI || 'http://127.0.0.1:7070'

console.info('NODE_ENV', process.env.NODE_ENV)
console.info('PROXY_URI', PROXY_URI)
console.info('AUTH_URI', AUTH_URI)
console.info('API_URI', API_URI)

// PLUGINS =============================================================================================================
const plugins = []

if (isEnablePWA) {
  const withPWA = require('@ducanh2912/next-pwa').default({
    dest: 'public',
    maximumFileSizeToCacheInBytes: 10 * 1024 * 1024, // Set the limit to 10 MB
    swcMinify: isProd,
    cacheOnFrontendNav: true,
    aggressiveFrontEndNavCaching: true,
    publicExcludes: ['!robots.txt', '!sitemap.xml'],
    extendDefaultRuntimeCaching: false,
    workboxOptions: {
      dontCacheBustURLsMatching: /\/api/,
    },
  })

  plugins.push(withPWA)
}

// Make sure adding Sentry options is the last code to run before exporting, to
// ensure that your source maps include changes from all other Webpack plugins
if (isEnableSentry) {
  const { withSentryConfig } = require('@sentry/nextjs')

  const config = {
    // For all available options, see:
    // https://github.com/getsentry/sentry-webpack-plugin#options

    // Suppresses source map uploading logs during build
    silent: true,

    org: 'batazor',
    project: 'shortlink-next',
  }

  plugins.push(() => withSentryConfig(config))
}

/** @type {import('@nx/next/plugins/with-nx').WithNxOptions} * */
const NEXT_CONFIG = {
  basePath: '/next',
  output: 'export',
  compress: isProd,
  productionBrowserSourceMaps: isProd,
  reactStrictMode: true,
  generateEtags: false,
  env: {
    // ShortLink API
    NEXT_PUBLIC_SERVICE_NAME: 'shortlink-next',
    NEXT_PUBLIC_API_URI: process.env.PROXY_URI,
    NEXT_PUBLIC_GIT_TAG: process.env.CI_COMMIT_TAG,

    // Sentry
    NEXT_PUBLIC_SENTRY_DSN: process.env.SENTRY_DSN,

    // Firebase
    NEXT_PUBLIC_FIREBASE_VAPID_KEY: process.env.FIREBASE_VAPID_KEY,
    NEXT_PUBLIC_FIREBASE_API_KEY: process.env.FIREBASE_API_KEY,
    NEXT_PUBLIC_FIREBASE_AUTH_DOMAIN: process.env.FIREBASE_AUTH_DOMAIN,
    NEXT_PUBLIC_FIREBASE_PROJECT_ID: process.env.FIREBASE_PROJECT_ID,
    NEXT_PUBLIC_FIREBASE_STORAGE_BUCKET: process.env.FIREBASE_STORAGE_BUCKET,
    NEXT_PUBLIC_FIREBASE_MESSAGING_SENDER_ID: process.env.FIREBASE_MESSAGING_SENDER_ID,
    NEXT_PUBLIC_FIREBASE_APP_ID: process.env.FIREBASE_APP_ID,
    NEXT_PUBLIC_FIREBASE_MEASUREMENT_ID: process.env.FIREBASE_MEASUREMENT_ID,

    // Faro
    NEXT_PUBLIC_FARO_URI: process.env.FARO_URI || 'http://localhost:3030',
  },
  transpilePackages: ['@shortlink-org/ui-kit'],
  compiler: {
    // ssr and displayName are configured by default
    emotion: true,
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
    dangerouslyAllowSVG: false,
    contentDispositionType: 'attachment',
    contentSecurityPolicy: "default-src 'self'; script-src 'none'; sandbox;",
  },
  trailingSlash: false,
  logging: {
    fetches: {
      fullUrl: true,
    },
  },
  webpack: (config, { isServer, buildId }) => {
    config.module.rules.push({
      test: /\.svg$/i,
      issuer: /\.[jt]sx?$/,
      use: ['@svgr/webpack'],
    })

    return config
  },
  experimental: {
    forceSwcTransforms: true,
    swcTraceProfiling: true,
    instrumentationHook: false,
    webVitalsAttribution: ['CLS', 'FCP', 'FID', 'INP', 'LCP', 'TTFB'],
    turbo: {},
    // typedRoutes: true,
  },
}

if (!isProd) {
  NEXT_CONFIG.rewrites = async () => ({
    beforeFiles: [
      // we need to define a no-op rewrite to trigger checking
      // all pages/static files before we attempt proxying
      {
        source: `/api/auth/:uri*`,
        destination: `${AUTH_URI}/:uri*`,
        basePath: false,
      },
      {
        source: `/api/links`,
        destination: `${API_URI}/api/links`,
        basePath: false,
      },
      {
        source: `/api/links/:uri*`,
        destination: `${API_URI}/api/links/:uri*`,
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
  })
}

module.exports = composePlugins(...plugins)(NEXT_CONFIG)
