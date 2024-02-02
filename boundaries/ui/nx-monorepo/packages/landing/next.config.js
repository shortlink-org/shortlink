const withPWA = require('@ducanh2912/next-pwa').default({
  dest: 'public',
  swcMinify: true,
  cacheOnFrontendNav: true,
  aggressiveFrontEndNavCaching: true,
})
const { composePlugins } = require('@nx/next')

// PLUGINS =============================================================================================================
const plugins = [withPWA]

// ENVIRONMENT VARIABLE ================================================================================================
const isProd = process.env.NODE_ENV === 'production'

/** @type {import('@nx/next/plugins/with-nx').WithNxOptions} * */
const nextConfig = {
  reactStrictMode: true,
  generateEtags: true,
  output: 'export',
  env: {
    // ShortLink API
    NEXT_PUBLIC_SERVICE_NAME: 'shortlink-landing',
    NEXT_PUBLIC_GIT_TAG: process.env.CI_COMMIT_TAG,

    // Sentry
    NEXT_PUBLIC_SENTRY_DSN: process.env.SENTRY_DSN,

    // Faro
    NEXT_PUBLIC_FARO_URI: process.env.FARO_URI,
  },
  swcMinify: true,
  productionBrowserSourceMaps: true,
  transpilePackages: ['@shortlink-org/ui-kit'],
  compiler: {
    // ssr and displayName are configured by default
    emotion: false,
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
    // webVitalsAttribution: ["CLS", "FCP", "FID", "INP", "LCP", "TTFB"],
    turbo: {},
    // for Vercel deployment
    useDeploymentId: true,
    useDeploymentIdServerActions: true,
  },
}

if (isProd) {
  nextConfig.compress = true
  nextConfig.productionBrowserSourceMaps = true
}

module.exports = composePlugins(...plugins)(nextConfig)
