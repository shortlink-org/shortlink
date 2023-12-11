// @ts-ignore
import { lightTheme } from '@shortlink-org/ui-kit'
import Document, { Html, Head, Main, NextScript } from 'next/document'
import * as React from 'react'

import { renderStatic } from '../pkg/renderer'

// TODO: research problem with faro
// initializeFaro({
//   url: process.env.NEXT_PUBLIC_FARO_URI,
//   app: {
//     name: process.env.NEXT_PUBLIC_SERVICE_NAME,
//     version: process.env.NEXT_PUBLIC_GIT_TAG,
//     environment: 'production',
//   },
//   instrumentations: [
//     // Mandatory, overwriting the instrumentations array would cause the default instrumentations to be omitted
//     ...getWebInstrumentations(),
//
//     // Initialization of the tracing package.
//     // This package is optional because it increases the bundle size noticeably. Only add it if you want tracing data.
//     new TracingInstrumentation(),
//   ],
// })

class MyDocument extends Document {
  static async getInitialProps(ctx: any) {
    const page = await ctx.renderPage()
    const { css, ids } = await renderStatic(page.html)
    const initialProps = await Document.getInitialProps(ctx)
    return {
      ...initialProps,
      styles: (
        <>
          {initialProps.styles}
          <style
            data-emotion={`css ${ids.join(' ')}`}
            dangerouslySetInnerHTML={{ __html: css }} // eslint-disable-line react/no-danger
          />
        </>
      ),
    }
  }

  render() {
    return (
      <Html lang="en" suppressHydrationWarning>
        {/* @ts-ignore */}
        <Head>
          <meta charSet="utf-8" />
          {/* PWA primary color */}
          <meta name="theme-color" content={lightTheme.palette.primary.main} />
          <link rel="shortcut icon" href="/next/favicon.ico" />
          <link
            rel="stylesheet"
            href="https://fonts.googleapis.com/css?family=Roboto:300,400,500,700&display=swap"
          />
        </Head>
        <body className="bg-white text-black dark:bg-black dark:text-white">
          <Main />
          {/* @ts-ignore */}
          <NextScript />
        </body>
      </Html>
    )
  }
}

export default MyDocument
