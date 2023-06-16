import createCache from '@emotion/cache'

// prepend: true moves MUI styles to the top of the <head> so they're loaded first.
// It allows developers to easily override MUI styles with other styling solutions, like CSS modules.
export const createEmotionCache = () =>
  createCache({ key: 'css', prepend: true })
