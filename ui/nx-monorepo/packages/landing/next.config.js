const withPWA = require('@ducanh2912/next-pwa').default({
  dest: 'public',
  swcMinify: true,
  cacheOnFrontendNav: true,
  aggressiveFrontEndNavCaching: true,
})
const { composePlugins } = require('@nx/next')

// PLUGINS =============================================================================================================
const plugins = [withPWA]

/** @type {import('@nx/next/plugins/with-nx').WithNxOptions} * */
const nextConfig = {
  reactStrictMode: true,
  generateEtags: true,
  output: 'export',
  env: {
    // ShortLink API
    NEXT_PUBLIC_SERVICE_NAME: 'shortlink-landing',
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
  experimental: {
    forceSwcTransforms: true,
    swcTraceProfiling: true,
    instrumentationHook: true,
    // webVitalsAttribution: ["CLS", "FCP", "FID", "INP", "LCP", "TTFB"],
    turbo: {},
    // for Vercel deployment
    useDeploymentId: true,
    // if you use with serverActions is desired
    serverActions: false,
    useDeploymentIdServerActions: true,
  },
}

module.exports = composePlugins(...plugins)(nextConfig)
