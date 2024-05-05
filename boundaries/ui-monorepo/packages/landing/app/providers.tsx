'use client'

import React, { useState } from 'react'
import { DEFAULT_ONLOAD_NAME, DEFAULT_SCRIPT_ID, SCRIPT_URL } from '@marsidev/react-turnstile'
import { Turnstile } from '@marsidev/react-turnstile'
import CssBaseline from '@mui/material/CssBaseline'
import { AppRouterCacheProvider } from '@mui/material-nextjs/v14-appRouter'
import { Experimental_CssVarsProvider as CssVarsProvider, getInitColorSchemeScript } from '@mui/material/styles'
import { theme } from '@shortlink-org/ui-kit/src/theme/theme'
import Script from 'next/script'

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
export function Providers({ children, ...props }) {
  const [isCaptcha, setIsCaptcha] = useState(false)

  return (
    <AppRouterCacheProvider>
      <CssVarsProvider theme={theme} defaultMode="dark">
        <Script id={DEFAULT_SCRIPT_ID} src={`${SCRIPT_URL}?onload=${DEFAULT_ONLOAD_NAME}`} strategy="afterInteractive" />
        {getInitColorSchemeScript()}

        <div className="flex m-auto text-black dark:bg-gray-800 dark:text-white flex-col">
          {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
          <CssBaseline />

          {isCaptcha && children}

          <Turnstile
            siteKey={process.env.NEXT_PUBLIC_CLOUDFLARE_SITE_KEY}
            injectScript={false}
            className="captcha"
            onSuccess={() => setIsCaptcha(true)}
            onError={() => setIsCaptcha(false)}
            onAbort={() => setIsCaptcha(false)}
          />
        </div>
      </CssVarsProvider>
    </AppRouterCacheProvider>
  )
}
