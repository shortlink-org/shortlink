// @ts-ignore
import { CacheProvider } from '@emotion/react'
import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp'
import CssBaseline from '@mui/material/CssBaseline'
import Fab from '@mui/material/Fab'
import {
  getInitColorSchemeScript,
  StyledEngineProvider,
  ThemeProvider,
} from '@mui/material/styles'
import { LocalizationProvider } from '@mui/x-date-pickers'
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'
import {
  ColorModeContext,
  createEmotionCache,
  darkTheme,
  lightTheme, // @ts-ignore
} from '@shortlink-org/ui-kit'
import Head from 'next/head'
import { DefaultSeo, LogoJsonLd, SiteLinksSearchBoxJsonLd } from 'next-seo'
import { ThemeProvider as NextThemeProvider } from 'next-themes'
import { StrictMode, useState } from 'react'
import 'react-toastify/dist/ReactToastify.css'
import { Provider } from 'react-redux'
import { Provider as BalancerProvider } from 'react-wrap-balancer'

import ScrollTop from 'components/ScrollTop'
import { wrapper } from 'store/store'
import '@shortlink-org/ui-kit/dist/cjs/index.css'

import '../public/assets/styles.css'

// Client-side cache, shared for the whole session of the user in the browser.
const clientSideEmotionCache = createEmotionCache()

// @ts-ignore
function MyApp({ Component, ...rest }) {
  const { store, props } = wrapper.useWrappedStore(rest)
  const { emotionCache = clientSideEmotionCache, pageProps } = props

  const [darkMode, setDarkMode] = useState(false)
  const theme = darkMode ? darkTheme : lightTheme

  return (
    <StrictMode>
      <LocalizationProvider dateAdapter={AdapterDayjs}>
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

                    <BalancerProvider>
                      <Component {...pageProps} />
                    </BalancerProvider>

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
      </LocalizationProvider>
    </StrictMode>
  )
}

export default MyApp
