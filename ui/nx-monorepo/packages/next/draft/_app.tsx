import React, { useState } from 'react'
import Fab from '@mui/material/Fab'
import { ThemeProvider as NextThemeProvider } from 'next-themes'
import CssBaseline from '@mui/material/CssBaseline'
import { CacheProvider, EmotionCache } from '@emotion/react'
import ScrollTop from 'components/ScrollTop'
import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp'

import {
  ColorModeContext,
  createEmotionCache,
  darkTheme,
  lightTheme, // @ts-ignore
} from '@shortlink-org/ui-kit'

// @ts-ignore
const MyApp = ({ Component, ...rest }) => {
  const { store, props } = wrapper.useWrappedStore(rest)
  const { emotionCache = clientSideEmotionCache, pageProps } = props

  const [darkMode, setDarkMode] = useState(false)
  const theme = darkMode ? darkTheme : lightTheme

  return (
    <React.StrictMode>
      <ThemeProvider theme={theme}>
        {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
        <CssBaseline />
        <ColorModeContext.Provider value={{ darkMode, setDarkMode }}>
          <NextThemeProvider enableSystem attribute="class">
            <Component {...pageProps} />
          </NextThemeProvider>
          <ScrollTop {...props}>
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
    </React.StrictMode>
  )
}

export default MyApp
