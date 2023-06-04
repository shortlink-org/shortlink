'use client'

import React, { useState } from 'react'
import createCache from '@emotion/cache'
import { CacheProvider } from '@emotion/react'
import { ThemeProvider as NextThemeProvider } from 'next-themes'
import { ThemeProvider } from '@mui/material/styles'
import CssBaseline from '@mui/material/CssBaseline'
import { getInitColorSchemeScript } from '@mui/material/styles'

import {
  ColorModeContext,
  darkTheme,
  lightTheme, // @ts-ignore
} from '@shortlink-org/ui-kit'
import { Turnstile } from '@marsidev/react-turnstile'

const cache = createCache({ key: 'next' })

// @ts-ignore
export function Providers({ children }) {
  const CLOUDFLARE_SITE_KEY = process.env.NEXT_PUBLIC_CLOUDFLARE_SITE_KEY

  const [darkMode, setDarkMode] = useState(false)
  const theme = darkMode ? darkTheme : lightTheme

  return (
    <ThemeProvider theme={theme}>
      <CacheProvider value={cache}>
        <NextThemeProvider enableSystem attribute="class">
          <ColorModeContext.Provider value={{ darkMode, setDarkMode }}>
            <div className="text-black dark:bg-gray-800 dark:text-white">
              {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
              <CssBaseline />
              {getInitColorSchemeScript()}

              {children}

              <Turnstile
                // @ts-ignore
                siteKey={CLOUDFLARE_SITE_KEY}
                injectScript={false}
                className="captcha"
              />
            </div>
          </ColorModeContext.Provider>
        </NextThemeProvider>
      </CacheProvider>
    </ThemeProvider>
  )
}
