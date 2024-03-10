import CssBaseline from '@mui/material/CssBaseline'
import {
  getInitColorSchemeScript,
  Experimental_CssVarsProvider as CssVarsProvider,
} from '@mui/material/styles'
// @ts-ignore
import { AppCacheProvider } from '@mui/material-nextjs/v14-pagesRouter'
import { LocalizationProvider } from '@mui/x-date-pickers'
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'
import { theme } from '@shortlink-org/ui-kit/src/theme/theme'
import Head from 'next/head'
import { DefaultSeo, LogoJsonLd, SiteLinksSearchBoxJsonLd } from 'next-seo'
import { ThemeProvider as NextThemeProvider } from 'next-themes'
import { StrictMode, useState } from 'react'
import 'react-toastify/dist/ReactToastify.css'
import { Provider } from 'react-redux'
import { Provider as BalancerProvider } from 'react-wrap-balancer'
import { useTheme } from '@mui/material'

import { wrapper } from 'store/store'
import '@shortlink-org/ui-kit/dist/cjs/index.css'

import '../public/assets/styles.css'

const defaultSeo = (theme: any) => (
  <>
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
    ,{/* @ts-ignore */}
    <SiteLinksSearchBoxJsonLd
      url="https://shortlink.best/"
      potentialActions={[
        {
          target: 'https://shortlink.best/search?q',
          queryInput: 'search_term_string',
        },
        {
          target: 'android-app://com.shortlink/https/shortlink.best/search?q',
          queryInput: 'search_term_string',
        },
      ]}
    />
    ,
    <LogoJsonLd
      logo="https://shortlink.best/images/logo.png"
      url="https://shortlink.best/"
    />
  </>
)

// @ts-ignore
function MyApp({ Component, ...rest }) {
  const { store, props } = wrapper.useWrappedStore(rest)
  const currentTheme = useTheme()

  return (
    <StrictMode>
      <LocalizationProvider dateAdapter={AdapterDayjs}>
        <AppCacheProvider options={{ key: 'css' }}>
          <NextThemeProvider enableSystem attribute="class">
            <Provider store={store}>
              <Head>
                <meta
                  name="viewport"
                  content="initial-scale=1, width=device-width"
                />
              </Head>

              <CssVarsProvider theme={theme}>
                {defaultSeo(currentTheme)}

                {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
                <CssBaseline />
                {getInitColorSchemeScript()}

                <BalancerProvider>
                  <Component {...props} />
                </BalancerProvider>
              </CssVarsProvider>
            </Provider>
          </NextThemeProvider>
        </AppCacheProvider>
      </LocalizationProvider>
    </StrictMode>
  )
}

export default MyApp
