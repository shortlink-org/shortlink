import React from 'react';
import App from 'next/app'
import Head from 'next/head';
import CssBaseline from '@material-ui/core/CssBaseline';
import { ThemeProvider } from '@material-ui/core/styles';
import Fab from '@material-ui/core/Fab';
import KeyboardArrowUpIcon from '@material-ui/icons/KeyboardArrowUp';
import sentry from '../utils/sentry'
import theme from '../theme/theme';
import ScrollTop from '../components/ScrollTop';
import "../assets/styles.css"

const { Sentry, captureException } = sentry()

export default class MyApp extends App {
  constructor() {
    super(...arguments)
    this.state = {
      hasError: false,
      errorEventId: undefined,
    }
  }

  static async getInitialProps({ Component, ctx }) {
    try {
      let pageProps = {}

      if (Component.getInitialProps) {
        pageProps = await Component.getInitialProps(ctx)
      }

      return { pageProps }
    } catch (error) {
      // Capture errors that happen during a page's getInitialProps.
      // This will work on both client and server sides.
      const errorEventId = captureException(error, ctx)
      return {
        hasError: true,
        errorEventId,
      }
    }
  }

  static getDerivedStateFromProps(props, state) {
    // If there was an error generated within getInitialProps, and we haven't
    // yet seen an error, we add it to this.state here
    return {
      hasError: props.hasError || state.hasError || false,
      errorEventId: props.errorEventId || state.errorEventId || undefined,
    }
  }

  static getDerivedStateFromError() {
    // React Error Boundary here allows us to set state flagging the error (and
    // later render a fallback UI).
    return { hasError: true }
  }

  componentDidCatch(error, errorInfo) {
    const errorEventId = captureException(error, { errorInfo })

    // Store the event id at this point as we don't have access to it within
    // `getDerivedStateFromError`.
    this.setState({ errorEventId })
  }

  render() {
    return this.state.hasError ? (
      <section>
        <h1>There was an error!</h1>
        <p>
          <a
            href="#"
            onClick={() =>
              Sentry.showReportDialog({ eventId: this.state.errorEventId })
            }
          >
            📣 Report this error
          </a>
        </p>
        <p>
          <a
            href="#"
            onClick={() => {
              window.location.reload(true)
            }}
          >
            Or, try reloading the page
          </a>
        </p>
      </section>
    ) : (
      // Render the normal Next.js page
      <React.Fragment>
        <Head>
          <title>Shortlink</title>
          <meta name="viewport" content="minimum-scale=1, initial-scale=1, width=device-width" />
        </Head>
        <ThemeProvider theme={theme}>
          {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
          <CssBaseline />
          { super.render() }
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
