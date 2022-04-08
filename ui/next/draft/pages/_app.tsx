// @ts-nocheck

import { appWithTranslation } from 'next-i18next'

import * as React from 'react';
import Head from 'next/head';
import { wrapper } from 'store/store'
import Fab from '@mui/material/Fab'
import App, { AppProps } from 'next/app';
import {
  ThemeProvider,
} from '@mui/material/styles'
import CssBaseline from '@mui/material/CssBaseline';
import { CacheProvider, EmotionCache } from '@emotion/react';
import theme from '../theme/theme';
import 'tailwindcss/tailwind.css'
import 'assets/styles.css'
import ScrollTop from 'components/ScrollTop'
import createEmotionCache from '../theme/createEmotionCache';
import { SESSION_FETCH_SUCCEEDED } from '../store/types/session'
import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp';

// Client-side cache, shared for the whole session of the user in the browser.
const clientSideEmotionCache = createEmotionCache();

interface MyAppProps extends AppProps {
  emotionCache?: EmotionCache;
}

class MyApp extends App<MyAppProps> {
  render() {
    const { Component, emotionCache = clientSideEmotionCache, pageProps } = this.props

    return (
      <CacheProvider value={emotionCache}>
        <Head>
          <meta name="viewport" content="initial-scale=1, width=device-width" />
        </Head>
        <ThemeProvider theme={theme}>
          {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
          <CssBaseline />
          <Component {...pageProps} />
          <ScrollTop {...this.props}>
            <Fab color="secondary" size="small" aria-label="scroll back to top">
              <KeyboardArrowUpIcon />
            </Fab>
          </ScrollTop>
        </ThemeProvider>
      </CacheProvider>
    )
  }
}

MyApp.getInitialProps = wrapper.getInitialAppProps(
  (store) =>
    async ({ Component, ctx }) => {
      // Init Kratos API
      const KRATOS_PUBLIC_API =
        process.env.KRATOS_PUBLIC_API || 'http://shortlink-api-kratos-public.shortlink:80'

      if (ctx.req?.headers) {
        // @ts-ignore
        const response = await fetch(`${KRATOS_PUBLIC_API}/sessions/whoami`, {
          headers: ctx.req?.headers,
        })

        // @ts-ignore
        const session = await response.json()

        // Save in store
        store.dispatch({ type: SESSION_FETCH_SUCCEEDED, payload: session })
      }

      return {
        pageProps: {
          // Call page-level getInitialProps
          // DON'T FORGET TO PROVIDE STORE TO PAGE
          ...(Component.getInitialProps
            ? await Component.getInitialProps({ ...ctx, store })
            : {}),
          // Some custom thing for all pages
          pathname: ctx.pathname,
        },
      }
    },
)

export default appWithTranslation(wrapper.withRedux(MyApp))
