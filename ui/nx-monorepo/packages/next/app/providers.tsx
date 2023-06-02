'use client'

import React, { useState } from 'react'
import createCache from '@emotion/cache'
import { CacheProvider } from '@emotion/react'
import { ThemeProvider as NextThemeProvider } from 'next-themes'
import { StyledEngineProvider, ThemeProvider } from '@mui/material/styles'
import CssBaseline from '@mui/material/CssBaseline'
import { getInitColorSchemeScript } from '@mui/material/styles'

import {
  ColorModeContext,
  darkTheme,
  lightTheme, // @ts-ignore
} from '@shortlink-org/ui-kit'
// @ts-ignore
import { ReduxProvider } from '../store/provider'
import Fab from '@mui/material/Fab'
import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp'
import ScrollTop from '../components/ScrollTop'

const cache = createCache({ key: 'next' })

// @ts-ignore
export function Providers({ Component, ...rest }) {
  const CLOUDFLARE_SITE_KEY = process.env.NEXT_PUBLIC_CLOUDFLARE_SITE_KEY

  const [darkMode, setDarkMode] = useState(false)
  const theme = darkMode ? darkTheme : lightTheme

  return (
    <NextThemeProvider enableSystem attribute="class">
      <ReduxProvider>
        <CacheProvider value={cache}>
          <StyledEngineProvider injectFirst>
            <ThemeProvider theme={theme}>
              {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
              <CssBaseline />
              <ColorModeContext.Provider value={{ darkMode, setDarkMode }}>
                {getInitColorSchemeScript()}
                {rest.children}

                {/* TODO: improve page up */}
                <ScrollTop>
                  <Fab
                    color="secondary"
                    size="small"
                    aria-label="scroll back to top"
                    className="bg-red-600 hover:bg-red-700"
                  >
                    <KeyboardArrowUpIcon />
                  </Fab>
                </ScrollTop>
              </ColorModeContext.Provider>
            </ThemeProvider>
          </StyledEngineProvider>
        </CacheProvider>
      </ReduxProvider>
    </NextThemeProvider>
  )
}
