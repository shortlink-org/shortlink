const { composePlugins } = require('@nx/next')

// ENVIRONMENT VARIABLE ================================================================================================
const isProd = process.env.NODE_ENV === 'production'
const isEnablePWA = process.env.PWA_ENABLE === 'true'

// PLUGINS =============================================================================================================
const plugins = []

if (isEnablePWA) {
  const withPWA = require('@ducanh2912/next-pwa').default({
    dest: 'public',
    swcMinify: isProd,
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
  compress: isProd,
  productionBrowserSourceMaps: isProd,
  env: {
    // ShortLink API
    NEXT_PUBLIC_SERVICE_NAME: 'shortlink-landing',

    // Captcha
    NEXT_PUBLIC_CLOUDFLARE_SITE_KEY: process.env.NEXT_PUBLIC_CLOUDFLARE_SITE_KEY,
    NEXT_PUBLIC_GOOGLE_ANALYTICS_ID: process.env.NEXT_PUBLIC_GOOGLE_ANALYTICS_ID,

    // Build info
    NEXT_PUBLIC_GIT_TAG: process.env.CI_COMMIT_TAG || process.env.CI_COMMIT_REF_NAME || 'local',
    NEXT_PUBLIC_PIPELINE_ID: process.env.CI_PIPELINE_ID || 'local',
    NEXT_PUBLIC_CI_PIPELINE_URL: process.env.CI_PIPELINE_URL || '#',

    // Sentry
    NEXT_PUBLIC_SENTRY_DSN: process.env.SENTRY_DSN,

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
  experimental: {
    forceSwcTransforms: true,
    swcTraceProfiling: true,
    instrumentationHook: false,
    webVitalsAttribution: ['CLS', 'FCP', 'FID', 'INP', 'LCP', 'TTFB'],
    turbo: {},
    // typedRoutes: true,
  },
}

module.exports = composePlugins(...plugins)(nextConfig)
