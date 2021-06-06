import React from 'react';
import { wrapper } from 'store/store';
import App from 'next/app';
import Head from 'next/head';
import CssBaseline from '@material-ui/core/CssBaseline';
import { ThemeProvider } from '@material-ui/core/styles';
import Fab from '@material-ui/core/Fab';
import KeyboardArrowUpIcon from '@material-ui/icons/KeyboardArrowUp';
import 'tailwindcss/tailwind.css';
import theme from 'theme/theme';
import ScrollTop from 'components/ScrollTop';
import 'assets/styles.css';

class MyApp extends App {
  constructor() {
    // @ts-ignore
    super(...arguments);
    this.state = {
      hasError: false,
      errorEventId: undefined,
    };
  }

  // @ts-ignore
  static getDerivedStateFromProps(props, state) {
    // If there was an error generated within getInitialProps, and we haven't
    // yet seen an error, we add it to this.state here
    return {
      hasError: props.hasError || state.hasError || false,
      errorEventId: props.errorEventId || state.errorEventId || undefined,
    };
  }

  static getDerivedStateFromError() {
    // React Error Boundary here allows us to set state flagging the error (and
    // later render a fallback UI).
    return { hasError: true };
  }

  render() {
    // @ts-ignore
    return this.state.hasError ? (
      <section>
        <h1>There was an error!</h1>
      </section>
    ) : (
      // Render the normal Next.js page
      <>
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
          {super.render()}
          <ScrollTop {...this.props}>
            <Fab color="secondary" size="small" aria-label="scroll back to top">
              <KeyboardArrowUpIcon />
            </Fab>
          </ScrollTop>
        </ThemeProvider>
      </>
    );
  }
}

export default wrapper.withRedux(MyApp);
