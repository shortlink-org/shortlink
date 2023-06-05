// @ts-ignore
import React, { useState } from 'react'
import Head from 'next/head'
// @ts-ignore
import { wrapper } from 'store/store'
import { Provider } from 'react-redux'
import Fab from '@mui/material/Fab'
import {
  getInitColorSchemeScript,
  StyledEngineProvider,
  ThemeProvider,
} from '@mui/material/styles'
import { ThemeProvider as NextThemeProvider } from 'next-themes'
import CssBaseline from '@mui/material/CssBaseline'
import { CacheProvider } from '@emotion/react'
import '../public/assets/styles.css'
// @ts-ignore
import ScrollTop from 'components/ScrollTop'
import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp'
import 'react-toastify/dist/ReactToastify.css'

import {
  ColorModeContext,
  createEmotionCache,
  darkTheme,
  lightTheme, // @ts-ignore
} from '@shortlink-org/ui-kit'
import { DefaultSeo, LogoJsonLd, SiteLinksSearchBoxJsonLd } from 'next-seo'

// Client-side cache, shared for the whole session of the user in the browser.
const clientSideEmotionCache = createEmotionCache()

// @ts-ignore
const MyApp = ({ Component, ...rest }) => {
  const { store, props } = wrapper.useWrappedStore(rest)
  const { emotionCache = clientSideEmotionCache, pageProps } = props

  const [darkMode, setDarkMode] = useState(false)
  const theme = darkMode ? darkTheme : lightTheme

  return (
    <React.StrictMode>
      <NextThemeProvider enableSystem attribute="class">
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
                siteName: 'Shortlink',
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
                <ColorModeContext.Provider value={{ darkMode, setDarkMode }}>
                  {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
                  <CssBaseline />
                  {getInitColorSchemeScript()}

                  <Component {...pageProps} />

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
      </NextThemeProvider>
    </React.StrictMode>
  )
}

export default MyApp
