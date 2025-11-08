'use client'

import React, { Suspense, useEffect, useState } from 'react'
import CssBaseline from '@mui/material/CssBaseline'
import { AppRouterCacheProvider } from '@mui/material-nextjs/v15-appRouter'
import { ThemeProvider, createTheme } from '@mui/material/styles'
import InitColorSchemeScript from '@mui/material/InitColorSchemeScript'
import { LocalizationProvider } from '@mui/x-date-pickers'
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'
import { Provider as BalancerProvider } from 'react-wrap-balancer'
import { Provider } from 'react-redux'
import { ThemeProvider as NextThemeProvider } from 'next-themes'

import 'react-toastify/dist/ReactToastify.css'

import { Layout } from '@/components'
import { makeStore, AppStore } from '@/store/store'

// Create a basic Material-UI theme
const theme = createTheme({
  palette: {
    mode: 'light',
  },
})

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

const loadingFallback = <div className="h-full justify-center">Loading...</div>

export function Providers({ children, ..._props }: { children: React.ReactNode; [key: string]: any }) {
  const [store, setStore] = useState<AppStore | null>(null)

  useEffect(() => {
    if (!store) {
      setStore(makeStore())
    }
  }, [store])

  return (
    <AppRouterCacheProvider options={{ enableCssLayer: true }}>
      <NextThemeProvider enableSystem attribute="class" defaultTheme={'light'}>
        <ThemeProvider theme={theme}>
          <InitColorSchemeScript />

          {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
          <CssBaseline />

          <LocalizationProvider dateAdapter={AdapterDayjs}>
            <Layout>
              <div className="text-black dark:bg-gray-800 dark:text-white">
                <BalancerProvider>
                  <Suspense fallback={loadingFallback}>
                    {store ? <Provider store={store}>{children}</Provider> : loadingFallback}
                  </Suspense>
                </BalancerProvider>
              </div>
            </Layout>
          </LocalizationProvider>
        </ThemeProvider>
      </NextThemeProvider>
    </AppRouterCacheProvider>
  )
}

export default Providers
