// ENVIRONMENT VARIABLE ================================================================================================
const isProd = process.env.NODE_ENV === 'production'

const AUTH_URI = process.env.AUTH_URI || 'http://127.0.0.1:4433'

import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  assetPrefix: '/auth',
};

if (!isProd) {
  // @ts-ignore
  nextConfig.rewrites = async () => ({
    beforeFiles: [
      {
        source: `/api/auth/:uri*`,
        destination: `${AUTH_URI}/:uri*`,
        basePath: false,
      },
    ],
  });
}

export default nextConfig;
