'use client'

import CssBaseline from '@mui/material/CssBaseline'
import { AppRouterCacheProvider } from '@mui/material-nextjs/v14-appRouter'
import {
  Experimental_CssVarsProvider as CssVarsProvider,
  getInitColorSchemeScript,
} from '@mui/material/styles'
import { theme } from '@shortlink-org/ui-kit/src/theme/theme'
import { ThemeProvider as NextThemeProvider } from 'next-themes'
import React from 'react'
import { LocalizationProvider } from '@mui/x-date-pickers'
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'
import { Provider as BalancerProvider } from 'react-wrap-balancer'

import 'react-toastify/dist/ReactToastify.css'
import '@shortlink-org/ui-kit/dist/cjs/index.css'

import { Layout } from 'components'

import { wrapper } from 'store/store'

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

// @ts-ignore
function Providers({ children, ...props }) {
  // const { store, props } = wrapper.useWrappedStore(rest)

  return (
    <AppRouterCacheProvider options={{ key: 'css' }}>
      <CssVarsProvider theme={theme} defaultMode="light">
        <LocalizationProvider dateAdapter={AdapterDayjs}>
          <NextThemeProvider
            enableSystem
            attribute="class"
            defaultTheme="light"
          >
            {getInitColorSchemeScript()}

            <Layout>
              <div className="text-black dark:bg-gray-800 dark:text-white">
                {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
                <CssBaseline />

                <BalancerProvider>{children}</BalancerProvider>
              </div>
            </Layout>
          </NextThemeProvider>
        </LocalizationProvider>
      </CssVarsProvider>
    </AppRouterCacheProvider>
  )
}

export default Providers
