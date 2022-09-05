import React, { useState } from 'react'
import Head from 'next/head';
import { AppProps, NextWebVitalsMetric } from 'next/app';
import { ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import { CacheProvider, EmotionCache } from '@emotion/react';
import { ThemeProvider as NextThemeProvider } from 'next-themes'
import theme from '../../next/theme/theme';
import createEmotionCache from '../src/createEmotionCache';
import '../public/assets/styles.css'
import ColorModeContext from "../../next/theme/ColorModeContext";
import darkTheme from "../../next/theme/darkTheme";

// Client-side cache, shared for the whole session of the user in the browser.
const clientSideEmotionCache = createEmotionCache();

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
      <ThemeProvider theme={darkMode ? darkTheme : theme}>
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
