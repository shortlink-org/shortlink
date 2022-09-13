// @ts-nocheck

import React, { useState } from 'react'
import Head from 'next/head'
import { wrapper } from 'store/store'
import Fab from '@mui/material/Fab'
import { AppInitialProps } from 'next/app'
import { ThemeProvider } from '@mui/material/styles'
import { ThemeProvider as NextThemeProvider } from 'next-themes'
import CssBaseline from '@mui/material/CssBaseline'
import { CacheProvider, EmotionCache } from '@emotion/react'
import { StyledEngineProvider } from '@mui/material/styles'
import 'public/assets/styles.css'
import ScrollTop from 'components/ScrollTop'
import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp'
import reportWebVitals from '../pkg/reportWebVitals'
// @ts-ignore
import { darkTheme, lightTheme, ColorModeContext, createEmotionCache } from '@shortlink-org/ui-kit'
import { DefaultSeo, SiteLinksSearchBoxJsonLd } from 'next-seo'

// Client-side cache, shared for the whole session of the user in the browser.
const clientSideEmotionCache = createEmotionCache()

interface MyAppProps extends AppInitialProps {
  emotionCache?: EmotionCache
}

const MyApp = (props: MyAppProps) => {
  const { Component, emotionCache = clientSideEmotionCache, pageProps } = props

  const [darkMode, setDarkMode] = useState(false)

  return (
    <React.StrictMode>
      <CacheProvider value={emotionCache}>
        <Head>
          <meta name="viewport" content="initial-scale=1, width=device-width" />
        </Head>
        <DefaultSeo
          openGraph={{
            type: 'website',
            locale: 'en_IE',
            url: 'https://architecture.ddns.net/',
            site_name: 'Shortlink',
            images: [
              {
                url: 'https://architecture.ddns.net/images/logo.png',
                width: 600,
                height: 600,
                alt: 'Shortlink service',
              },
            ],
          }}
          twitter={{
            handle: '@shortlink',
            site: '@shortlink',
            cardType: 'summary_large_image',
          }}
          titleTemplate={'Shortlink | %s'}
          defaultTitle={'Shortlink'}
        />

        {/* @ts-ignore */}
        <SiteLinksSearchBoxJsonLd
          url="https://architecture.ddns.net/"
          potentialActions={[
            {
              target: 'https://architecture.ddns.net/search?q',
              queryInput: 'search_term_string',
            },
            {
              target: 'android-app://com.shortlink/https/architecture.ddns.net/search?q',
              queryInput: 'search_term_string',
            },
          ]}
        />

        <StyledEngineProvider injectFirst>
          <ThemeProvider theme={darkMode ? darkTheme : lightTheme}>
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
        </StyledEngineProvider>
      </CacheProvider>
    </React.StrictMode>
  )
}

export default wrapper.withRedux(MyApp)

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals()
