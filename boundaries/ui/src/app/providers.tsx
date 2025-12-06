'use client'

import React, { useEffect, useState, useMemo } from 'react'
import CssBaseline from '@mui/material/CssBaseline'
import { AppRouterCacheProvider } from '@mui/material-nextjs/v15-appRouter'
import { ThemeProvider, createTheme, useColorScheme } from '@mui/material/styles'
import { LocalizationProvider } from '@mui/x-date-pickers'
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'
import { Provider as BalancerProvider } from 'react-wrap-balancer'
import { Provider } from 'react-redux'
import { ThemeProvider as NextThemeProvider, useTheme } from 'next-themes'

import 'react-toastify/dist/ReactToastify.css'

import { Layout } from '@/components'
import { SessionWrapper } from '@/components/SessionWrapper'
import { makeStore, AppStore } from '@/store/store'

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

/**
 * ThemeSync - Syncs next-themes with MUI ColorScheme
 */
function ThemeSync({ children }: { children: React.ReactNode }) {
  const { resolvedTheme } = useTheme()
  const { setMode, mode } = useColorScheme()

  useEffect(() => {
    if (resolvedTheme && resolvedTheme !== 'system') {
      const newMode = resolvedTheme as 'light' | 'dark'
      if (mode !== newMode) {
        setMode(newMode)
      }
    }
  }, [resolvedTheme, setMode, mode])

  return <>{children}</>
}

/**
 * MuiThemeProvider - Dynamic MUI theme that syncs with next-themes and supports useColorScheme
 * Must be inside NextThemeProvider to access useTheme()
 */
function MuiThemeProvider({ children }: { children: React.ReactNode }) {
  const { resolvedTheme } = useTheme()
  const [mounted, setMounted] = useState(false)

  useEffect(() => {
    setMounted(true)
  }, [])

  // Create theme with CSS Variables support
  const theme = useMemo(() => {
    const mode = mounted && resolvedTheme && resolvedTheme !== 'system' ? (resolvedTheme as 'light' | 'dark') : 'light'

    return createTheme({
      cssVariables: true,
      palette: {
        mode,
        ...(mode === 'dark' && {
          background: {
            default: '#0a0a0a',
            paper: '#1a1a1a',
          },
        }),
      },
    })
  }, [resolvedTheme, mounted])

  return (
    <ThemeProvider theme={theme}>
      <ThemeSync>
        <CssBaseline />
        {children}
      </ThemeSync>
    </ThemeProvider>
  )
}

export function Providers({ children, ..._props }: { children: React.ReactNode; [key: string]: any }) {
  const store = useMemo(() => makeStore(), [])

  return (
    <AppRouterCacheProvider options={{ enableCssLayer: true }}>
      <NextThemeProvider enableSystem attribute="class" defaultTheme="light">
        <MuiThemeProvider>
          <LocalizationProvider dateAdapter={AdapterDayjs}>
            <Layout>
              <BalancerProvider>
                <Provider store={store}>{children}</Provider>
              </BalancerProvider>
            </Layout>
          </LocalizationProvider>
        </MuiThemeProvider>
      </NextThemeProvider>
    </AppRouterCacheProvider>
  )
}

/**
 * ProvidersWithSession - Wraps Providers with SessionWrapper
 *
 * This ensures session is available to all components
 */
export function ProvidersWithSession({ children, ...props }: { children: React.ReactNode; [key: string]: any }) {
  return (
    <SessionWrapper>
      <Providers {...props}>{children}</Providers>
    </SessionWrapper>
  )
}

export default Providers
