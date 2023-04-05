/* eslint-disable */

const withExportImages = require('next-export-optimize-images')

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

/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  generateEtags: true,
  env: {
    // ShortLink API
    NEXT_PUBLIC_SERVICE_NAME: "shortlink-landing",
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
      },
    },
  },
}

module.exports = withExportImages(nextConfig)
