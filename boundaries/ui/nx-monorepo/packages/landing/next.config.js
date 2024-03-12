const { composePlugins } = require('@nx/next')

// ENVIRONMENT VARIABLE ================================================================================================
const isProd = process.env.NODE_ENV === 'production'
const isEnablePWA = process.env.PWA_ENABLE === 'true'

// PLUGINS =============================================================================================================
const plugins = []

if (isEnablePWA) {
  // eslint-disable-next-line global-require
  const withPWA = require('@ducanh2912/next-pwa').default({
    dest: 'public',
    swcMinify: process.env.NODE_ENV === 'production',
    cacheOnFrontendNav: true,
    aggressiveFrontEndNavCaching: true,
  })

  plugins.push(withPWA)
}

/** @type {import('@nx/next/plugins/with-nx').WithNxOptions} * */
const nextConfig = {
  reactStrictMode: true,
  generateEtags: false,
  output: 'export',
  swcMinify: isProd,
  compress: isProd,
  productionBrowserSourceMaps: isProd,
  env: {
    // ShortLink API
    NEXT_PUBLIC_SERVICE_NAME: 'shortlink-landing',
    NEXT_PUBLIC_GIT_TAG: process.env.CI_COMMIT_TAG,

    // Sentry
    NEXT_PUBLIC_SENTRY_DSN: process.env.SENTRY_DSN,

    // Faro
    NEXT_PUBLIC_FARO_URI: process.env.FARO_URI,
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
  experimental: {
    forceSwcTransforms: true,
    swcTraceProfiling: true,
    instrumentationHook: true,
    webVitalsAttribution: ['CLS', 'FCP', 'FID', 'INP', 'LCP', 'TTFB'],
    turbo: {},
    // typedRoutes: true,
    // for Vercel deployment
    useDeploymentId: true,
    useDeploymentIdServerActions: true,
    modularizeImports: {
      '@mui/material': {
        transform: '@mui/material/{{member}}',
      },
      '@mui/icons-material': {
        transform: '@mui/icons-material/{{member}}',
      },
    },
  },
}

module.exports = composePlugins(...plugins)(nextConfig)
