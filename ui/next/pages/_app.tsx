// @ts-nocheck

import * as React from 'react'
import Head from 'next/head'
import { wrapper } from 'store/store'
import Fab from '@mui/material/Fab'
import App, { AppProps } from 'next/app'
import { ThemeProvider } from '@mui/material/styles'
import CssBaseline from '@mui/material/CssBaseline'
import { CacheProvider, EmotionCache } from '@emotion/react'
import { StyledEngineProvider } from '@mui/material/styles'
import theme from '../theme/theme'
import 'public/assets/styles.css'
import ScrollTop from 'components/ScrollTop'
import createEmotionCache from '../theme/createEmotionCache'
import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp'
import reportWebVitals from '../pkg/reportWebVitals'

// Client-side cache, shared for the whole session of the user in the browser.
const clientSideEmotionCache = createEmotionCache()

interface MyAppProps extends AppProps {
  emotionCache?: EmotionCache
}

class MyApp extends App<MyAppProps> {
  render() {
    const {
      Component,
      emotionCache = clientSideEmotionCache,
      pageProps,
    } = this.props

    return (
      <React.StrictMode>
        <CacheProvider value={emotionCache}>
          <Head>
            <meta
              name="viewport"
              content="initial-scale=1, width=device-width"
            />
          </Head>
          <StyledEngineProvider injectFirst>
            <ThemeProvider theme={theme}>
              {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
              <CssBaseline />
              <Component {...pageProps} />
              <ScrollTop {...this.props}>
                <Fab
                  color="secondary"
                  size="small"
                  aria-label="scroll back to top"
                  className="bg-red-600 hover:bg-red-700"
                >
                  <KeyboardArrowUpIcon />
                </Fab>
              </ScrollTop>
            </ThemeProvider>
          </StyledEngineProvider>
        </CacheProvider>
      </React.StrictMode>
    )
  }
}

export default wrapper.withRedux(MyApp)

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals()
