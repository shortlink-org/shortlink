import React, { useState } from 'react'
import Head from 'next/head'
import { AppProps, NextWebVitalsMetric } from 'next/app'
import { ThemeProvider } from '@mui/material/styles'
import CssBaseline from '@mui/material/CssBaseline'
import { CacheProvider, EmotionCache } from '@emotion/react'
import { ThemeProvider as NextThemeProvider } from 'next-themes'
import { Roboto_Mono } from '@next/font/google' // eslint-disable-line
import { DefaultSeo, LogoJsonLd, SiteLinksSearchBoxJsonLd } from 'next-seo'
import '../public/assets/styles.css'

import {
  ColorModeContext,
  createEmotionCache,
  darkTheme,
  lightTheme, // @ts-ignore
} from '@shortlink-org/ui-kit'

// Client-side cache, shared for the whole session of the user in the browser.
const clientSideEmotionCache = createEmotionCache()

interface MyAppProps extends AppProps {
  emotionCache?: EmotionCache
}

const fonts = Roboto_Mono({
  subsets: ['latin'],
  variable: '--font-inter',
})

const MyApp = (props: MyAppProps) => {
  const { Component, emotionCache = clientSideEmotionCache, pageProps } = props

  const [darkMode, setDarkMode] = useState(false)
  const theme = darkMode ? darkTheme : lightTheme

  return (
    <CacheProvider value={emotionCache}>
      <Head>
        <meta name="viewport" content="initial-scale=1, width=device-width" />
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
            target: 'android-app://com.shortlink/https/shortlink.best/search?q',
            queryInput: 'search_term_string',
          },
        ]}
      />

      <LogoJsonLd
        logo="https://shortlink.best/images/logo.png"
        url="https://shortlink.best/"
      />

      <ThemeProvider theme={theme}>
        {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
        <CssBaseline />
        <ColorModeContext.Provider value={{ darkMode, setDarkMode }}>
          <NextThemeProvider enableSystem attribute="class">
            <main className={`${fonts.variable} font-sans`}>
              <Component {...pageProps} />
            </main>
          </NextThemeProvider>
        </ColorModeContext.Provider>
      </ThemeProvider>
    </CacheProvider>
  )
}

export function reportWebVitals(metric: NextWebVitalsMetric) {
  console.log(metric)
}

export default MyApp
