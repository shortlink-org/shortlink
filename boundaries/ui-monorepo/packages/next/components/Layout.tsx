import Box from '@mui/material/Box'
import CssBaseline from '@mui/material/CssBaseline'
import * as React from 'react'
// @ts-ignore
import { ScrollToTopButton } from '@shortlink-org/ui-kit'

import PushNotificationLayout from 'components/PushNotificationLayout'
import Footer from 'components/Footer'
import Header from 'components/Header'

// @ts-ignore
export function Layout({ children }) {
  return (
    <PushNotificationLayout>
      <CssBaseline />

      <div className="flex flex-row h-full mt-8">
        <Header />

        <main className={'flex-auto'}>
          <div>{children}</div>

          <Footer />

          <ScrollToTopButton />
        </main>
      </div>
    </PushNotificationLayout>
  )
}
