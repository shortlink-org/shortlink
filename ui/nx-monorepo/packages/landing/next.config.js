/* eslint-disable */

const withPWA = require('next-pwa')({
  dest: 'public',
})

/** @type {import('next').NextConfig} */
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
    styledComponents: true,
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
  },
  trailingSlash: false,
  experimental: {
    forceSwcTransforms: true,
    swcTraceProfiling: true,
    instrumentationHook: true,
    appDir: true,
    // webVitalsAttribution: ["CLS", "FCP", "FID", "INP", "LCP", "TTFB"],
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
      },
    },
  },
}

module.exports = withPWA(nextConfig)
