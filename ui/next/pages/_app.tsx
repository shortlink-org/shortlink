// @ts-nocheck

import { appWithTranslation } from 'next-i18next'

import React from 'react'
import { wrapper } from 'store/store'
import App, { AppInitialProps } from 'next/app'
import Head from 'next/head'
import CssBaseline from '@material-ui/core/CssBaseline'
import { ThemeProvider } from '@material-ui/core/styles'
import Fab from '@material-ui/core/Fab'
import KeyboardArrowUpIcon from '@material-ui/icons/KeyboardArrowUp'
import 'tailwindcss/tailwind.css'
import theme from 'theme/theme'
import ScrollTop from 'components/ScrollTop'
import 'assets/styles.css'
import { SESSION_FETCH_SUCCEEDED } from '../store/types/session'

class MyApp extends App<AppInitialProps> {
  render() {
    const { Component, pageProps } = this.props

    // @ts-ignore
    return (
      // Render the normal Next.js page
      <React.Fragment>
        <Head>
          <title>Shortlink</title>
          <meta charSet="utf-8" />
          <meta
            name="viewport"
            content="minimum-scale=1, initial-scale=1, width=device-width"
          />
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
      </React.Fragment>
    )
  }
}

MyApp.getInitialProps = wrapper.getInitialAppProps(
  (store) =>
    async ({ Component, ctx }) => {
      // Init Kratos API
      const KRATOS_PUBLIC_API =
        process.env.KRATOS_API || 'http://127.0.0.1:4433'

      if (ctx.req?.headers) {
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
