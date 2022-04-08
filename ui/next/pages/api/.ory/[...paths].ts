// @ory/integrations offers a package for integrating with NextJS.
import { config, createApiHandler } from '@ory/integrations/next-edge'

// We need to export the config.
export { config }

// And create the Ory Cloud API "bridge".
export default createApiHandler({
  fallbackToPlayground: true,
  // Because vercel.app is a public suffix and setting cookies for
  // vercel.app is not possible.
  dontUseTldForCookieDomain: true
})
