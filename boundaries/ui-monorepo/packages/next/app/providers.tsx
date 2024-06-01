'use client'

import React, { useRef, Suspense } from 'react'
import CssBaseline from '@mui/material/CssBaseline'
import { AppRouterCacheProvider } from '@mui/material-nextjs/v14-appRouter'
import { Experimental_CssVarsProvider as CssVarsProvider, getInitColorSchemeScript } from '@mui/material/styles'
import { theme } from '@shortlink-org/ui-kit/src/theme/theme'
import { LocalizationProvider } from '@mui/x-date-pickers'
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'
import { Provider as BalancerProvider } from 'react-wrap-balancer'
import { Provider } from 'react-redux'
import { ThemeProvider as NextThemeProvider } from 'next-themes'

import 'react-toastify/dist/ReactToastify.css'
import '@shortlink-org/ui-kit/dist/cjs/index.css'

import { Layout } from 'components'
import { makeStore, AppStore } from 'store/store'

// TODO: faro has old peer dependencies, so we need to fix it before enabling it
//
// import { TracingInstrumentation } from '@grafana/faro-web-tracing'
// import { getWebInstrumentations, initializeFaro } from '@grafana/faro-web-sdk'
//
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
  const storeRef = useRef<AppStore>()
  if (!storeRef.current) {
    // Create the store instance the first time this renders
    // @ts-ignore
    storeRef.current = makeStore()
  }

  return (
    <AppRouterCacheProvider>
      <NextThemeProvider enableSystem attribute="class" defaultTheme={'light'}>
        <CssVarsProvider theme={theme} defaultMode="light">
          {getInitColorSchemeScript()}

          {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
          <CssBaseline />

          <LocalizationProvider dateAdapter={AdapterDayjs}>
            <Layout>
              <div className="text-black dark:bg-gray-800 dark:text-white">
                <BalancerProvider>
                  <Suspense fallback={<div className={'h-full justify-center'}>Loading...</div>}>
                    <Provider store={storeRef.current}>{children}</Provider>
                  </Suspense>
                </BalancerProvider>
              </div>
            </Layout>
          </LocalizationProvider>
        </CssVarsProvider>
      </NextThemeProvider>
    </AppRouterCacheProvider>
  )
}

export default Providers
