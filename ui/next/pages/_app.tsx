import React, {useEffect} from 'react'
import { wrapper } from 'store/store'
import App, { AppInitialProps, AppContext } from "next/app"
import Head from 'next/head'
import CssBaseline from '@material-ui/core/CssBaseline'
import { ThemeProvider } from '@material-ui/core/styles'
import Fab from '@material-ui/core/Fab'
import KeyboardArrowUpIcon from '@material-ui/icons/KeyboardArrowUp'
import 'tailwindcss/tailwind.css'
import theme from 'theme/theme'
import ScrollTop from 'components/ScrollTop'
import 'assets/styles.css'

class MyApp extends App<AppInitialProps> {
  public static getInitialProps = wrapper.getInitialAppProps(store => async ({Component, ctx}) => {
    // Keep in mind that this will be called twice on server, one for page and second for error page
    store.dispatch({ type: "APP", payload: "was set in _app" })

    return {
      pageProps: {
        // Call page-level getInitialProps
        // DON'T FORGET TO PROVIDE STORE TO PAGE
        ...(Component.getInitialProps ? await Component.getInitialProps({...ctx, store}) : {}),
        // Some custom thing for all pages
        pathname: ctx.pathname,
      }
    }
  })

  render() {
    const { Component, pageProps } = this.props

    // @ts-ignore
    return (
      // Render the normal Next.js page
      <React.Fragment>
        <Head>
          <title>Shortlink</title>
          <meta charSet="utf-8"/>
          <meta
            name="viewport"
            content="minimum-scale=1, initial-scale=1, width=device-width"
          />
        </Head>
        <ThemeProvider theme={theme}>
          {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
          <CssBaseline/>
          <Component {...pageProps} />
          <ScrollTop {...this.props}>
            <Fab color="secondary" size="small" aria-label="scroll back to top">
              <KeyboardArrowUpIcon/>
            </Fab>
          </ScrollTop>
        </ThemeProvider>
      </React.Fragment>
    )
  }
}

export default wrapper.withRedux(MyApp)
