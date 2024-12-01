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

/** @type {import('next').NextConfig} */
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
  bundlePagesRouterDependencies: true,
  experimental: {
    forceSwcTransforms: true,
    swcTraceProfiling: true,
    webVitalsAttribution: ['CLS', 'FCP', 'FID', 'INP', 'LCP', 'TTFB'],
    turbo: {},
    // reactCompiler: true,
    // typedRoutes: true,
  },
}

module.exports = nextConfig

// module.exports = composePlugins(...plugins)(nextConfig)


// Injected content via Sentry wizard below

const { withSentryConfig } = require("@sentry/nextjs");

module.exports = withSentryConfig(
  module.exports,
  {
    // For all available options, see:
    // https://github.com/getsentry/sentry-webpack-plugin#options

    org: "batazor-cj",
    project: "shortlink-ui",

    // Only print logs for uploading source maps in CI
    silent: !process.env.CI,

    // For all available options, see:
    // https://docs.sentry.io/platforms/javascript/guides/nextjs/manual-setup/

    // Upload a larger set of source maps for prettier stack traces (increases build time)
    widenClientFileUpload: true,

    // Automatically annotate React components to show their full name in breadcrumbs and session replay
    reactComponentAnnotation: {
      enabled: true,
    },

    // Uncomment to route browser requests to Sentry through a Next.js rewrite to circumvent ad-blockers.
    // This can increase your server load as well as your hosting bill.
    // Note: Check that the configured route will not match with your Next.js middleware, otherwise reporting of client-
    // side errors will fail.
    // tunnelRoute: "/monitoring",

    // Hides source maps from generated client bundles
    hideSourceMaps: true,

    // Automatically tree-shake Sentry logger statements to reduce bundle size
    disableLogger: true,

    // Enables automatic instrumentation of Vercel Cron Monitors. (Does not yet work with App Router route handlers.)
    // See the following for more information:
    // https://docs.sentry.io/product/crons/
    // https://vercel.com/docs/cron-jobs
    automaticVercelMonitors: true,
  }
);
