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
import { Turnstile } from "@marsidev/react-turnstile"

const cache = createCache({ key: 'next' })

// @ts-ignore
export function Providers({ children }) {
  const CLOUDFLARE_SITE_KEY = process.env.NEXT_PUBLIC_CLOUDFLARE_SITE_KEY

  const [darkMode, setDarkMode] = useState(false)
  const theme = darkMode ? darkTheme : lightTheme

  return (
    <NextThemeProvider>
      <CacheProvider value={cache}>
        <ThemeProvider theme={theme}>
          {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
          <CssBaseline />
          <ColorModeContext.Provider value={{ darkMode, setDarkMode }}>
            {getInitColorSchemeScript()}
            {children}

            {/* @ts-ignore */}
            <Turnstile siteKey={CLOUDFLARE_SITE_KEY} injectScript={false} className="captcha" />
          </ColorModeContext.Provider>
        </ThemeProvider>
      </CacheProvider>
    </NextThemeProvider>
  )
}
