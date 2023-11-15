'use client'

import createCache from '@emotion/cache'
import { CacheProvider } from '@emotion/react'
import {
  DEFAULT_ONLOAD_NAME,
  DEFAULT_SCRIPT_ID,
  SCRIPT_URL,
} from '@marsidev/react-turnstile'
import { Turnstile } from '@marsidev/react-turnstile'
import CssBaseline from '@mui/material/CssBaseline'
import { ThemeProvider, getInitColorSchemeScript } from '@mui/material/styles'
import {
  ColorModeContext,
  darkTheme,
  lightTheme, // @ts-ignore
} from '@shortlink-org/ui-kit'
import Script from 'next/script'
import { ThemeProvider as NextThemeProvider } from 'next-themes'
import React, { useState } from 'react'

const cache = createCache({ key: 'next' })

// TODO: research problem with faro
// initializeFaro({
//   url: process.env.NEXT_PUBLIC_FARO_URI,
//   app: {
//     name: process.env.NEXT_PUBLIC_SERVICE_NAME,
//     version: process.env.NEXT_PUBLIC_GIT_TAG,
//     environment: 'production',
//   },
//   instrumentations: [
//     // Mandatory, overwriting the instrumentations array would cause the default instrumentations to be omitted
//     ...getWebInstrumentations(),
//
//     // Initialization of the tracing package.
//     // This package is optional because it increases the bundle size noticeably. Only add it if you want tracing data.
//     new TracingInstrumentation(),
//   ],
// })

// @ts-ignore
export function Providers({ children }) {
  const CLOUDFLARE_SITE_KEY = process.env.NEXT_PUBLIC_CLOUDFLARE_SITE_KEY

  const [darkMode, setDarkMode] = useState(false)
  const theme = darkMode ? darkTheme : lightTheme

  const [isCaptcha, setIsCaptcha] = useState(false)

  return (
    <ThemeProvider theme={theme}>
      <CacheProvider value={cache}>
        <NextThemeProvider enableSystem attribute="class">
          <ColorModeContext.Provider value={{ darkMode, setDarkMode }}>
            <Script
              id={DEFAULT_SCRIPT_ID}
              src={`${SCRIPT_URL}?onload=${DEFAULT_ONLOAD_NAME}`}
              strategy="afterInteractive"
            />
            <div className="text-black dark:bg-gray-800 dark:text-white">
              {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
              <CssBaseline />
              {getInitColorSchemeScript()}

              {isCaptcha && children}

              <Turnstile
                // @ts-ignore
                siteKey={CLOUDFLARE_SITE_KEY}
                injectScript={false}
                className="captcha"
                onSuccess={() => setIsCaptcha(true)}
                onError={() => setIsCaptcha(false)}
                onAbort={() => setIsCaptcha(false)}
              />
            </div>
          </ColorModeContext.Provider>
        </NextThemeProvider>
      </CacheProvider>
    </ThemeProvider>
  )
}
