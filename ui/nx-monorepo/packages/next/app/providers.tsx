'use client'

import React, { useState } from 'react'
import { CacheProvider } from '@emotion/react'
import createCache from '@emotion/cache'
import { ThemeProvider as NextThemeProvider } from 'next-themes'
import { StyledEngineProvider, ThemeProvider } from '@mui/material/styles'
import CssBaseline from '@mui/material/CssBaseline'
import { getInitColorSchemeScript } from '@mui/material/styles'
// @ts-ignore
import { Layout } from 'components'

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

// You can consider sharing the same emotion cache between all the SSR requests to speed up performance.
// However, be aware that it can have global side effects.
const cache = createCache({ key: 'next' })

// @ts-ignore
export function Providers({ children }) {
  const [darkMode, setDarkMode] = useState(false)
  const theme = darkMode ? darkTheme : lightTheme

  return (
    <ThemeProvider theme={theme}>
      <CacheProvider value={cache}>
        <NextThemeProvider enableSystem attribute="class">
          <ColorModeContext.Provider value={{ darkMode, setDarkMode }}>
            <ReduxProvider>
              <StyledEngineProvider injectFirst>
                <div className="text-black dark:bg-gray-800 dark:text-white">
                  {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
                  <CssBaseline />
                  {getInitColorSchemeScript()}

                  <Layout>{children}</Layout>

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
                </div>
              </StyledEngineProvider>
            </ReduxProvider>
          </ColorModeContext.Provider>
        </NextThemeProvider>
      </CacheProvider>
    </ThemeProvider>
  )
}
