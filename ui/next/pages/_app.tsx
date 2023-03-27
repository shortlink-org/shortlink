// @ts-nocheck

import React, { useState } from 'react'
import Head from 'next/head'
import { wrapper } from 'store/store'
import { Provider } from 'react-redux'
import Fab from '@mui/material/Fab'
import { AppInitialProps, NextWebVitalsMetric } from 'next/app'
import { StyledEngineProvider, ThemeProvider } from '@mui/material/styles'
import { ThemeProvider as NextThemeProvider } from 'next-themes'
import CssBaseline from '@mui/material/CssBaseline'
import { CacheProvider, EmotionCache } from '@emotion/react'
import 'public/assets/styles.css'
import ScrollTop from 'components/ScrollTop'
import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp'
import 'react-toastify/dist/ReactToastify.css'

// @ts-ignore
import {
  ColorModeContext,
  createEmotionCache,
  darkTheme,
  lightTheme,
} from '@shortlink-org/ui-kit'
import { DefaultSeo, LogoJsonLd, SiteLinksSearchBoxJsonLd } from 'next-seo'

// Client-side cache, shared for the whole session of the user in the browser.
const clientSideEmotionCache = createEmotionCache()

interface MyAppProps extends AppInitialProps {
  emotionCache?: EmotionCache
}

const MyApp: FC<AppProps> = ({ Component, ...rest }) => {
  const { store, props } = wrapper.useWrappedStore(rest)
  const { emotionCache = clientSideEmotionCache, pageProps } = props

  const [darkMode, setDarkMode] = useState(false)
  const theme = darkMode ? darkTheme : lightTheme

  return (
    <React.StrictMode>
      <Provider store={store}>
        <CacheProvider value={emotionCache}>
          <Head>
            <meta
              name="viewport"
              content="initial-scale=1, width=device-width"
            />
          </Head>
          <DefaultSeo
            openGraph={{
              type: 'website',
              locale: 'en_IE',
              url: 'https://shortlink.best/',
              site_name: 'Shortlink',
              images: [
                {
                  url: 'https://shortlink.best/images/logo.png',
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
            titleTemplate="Shortlink | %s"
            defaultTitle="Shortlink"
            themeColor={theme.palette.primary.main}
          />

          {/* @ts-ignore */}
          <SiteLinksSearchBoxJsonLd
            url="https://shortlink.best/"
            potentialActions={[
              {
                target: 'https://shortlink.best/search?q',
                queryInput: 'search_term_string',
              },
              {
                target:
                  'android-app://com.shortlink/https/shortlink.best/search?q',
                queryInput: 'search_term_string',
              },
            ]}
          />

          <LogoJsonLd
            logo="https://shortlink.best/images/logo.png"
            url="https://shortlink.best/"
          />

          <StyledEngineProvider injectFirst>
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
          </StyledEngineProvider>
        </CacheProvider>
      </Provider>
    </React.StrictMode>
  )
}

export default MyApp

export function reportWebVitals(metric: NextWebVitalsMetric) {
  console.log(metric)
}
