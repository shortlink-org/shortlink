import React, { useState } from 'react'
import Head from 'next/head';
import { AppProps, NextWebVitalsMetric } from 'next/app';
import { ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import { CacheProvider, EmotionCache } from '@emotion/react';
import { ThemeProvider as NextThemeProvider } from 'next-themes'
import { DefaultSeo, SiteLinksSearchBoxJsonLd } from "next-seo";
import '../public/assets/styles.css'
// @ts-ignore
import { createEmotionCache, darkTheme, lightTheme, ColorModeContext } from '@shortlink-org/ui-kit'
// import your default seo configuration

// Client-side cache, shared for the whole session of the user in the browser.
const clientSideEmotionCache = createEmotionCache()

interface MyAppProps extends AppProps {
  emotionCache?: EmotionCache
}

const MyApp = (props: MyAppProps) => {
  const {
    Component,
    emotionCache = clientSideEmotionCache,
    pageProps,
  } = props

  const [darkMode, setDarkMode] = useState(false)

  return (
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

      <ThemeProvider theme={darkMode ? darkTheme : lightTheme}>
        {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
        <CssBaseline />
        <ColorModeContext.Provider value={{ darkMode, setDarkMode }}>
          <NextThemeProvider enableSystem attribute="class">
            <Component {...pageProps} />
          </NextThemeProvider>
        </ColorModeContext.Provider>
      </ThemeProvider>
    </CacheProvider>
  );
}

export function reportWebVitals(metric: NextWebVitalsMetric) {
  console.log(metric)
}

export default MyApp
