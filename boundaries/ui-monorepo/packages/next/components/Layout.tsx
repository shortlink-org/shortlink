import CssBaseline from '@mui/material/CssBaseline'
import React, { useEffect, useState } from 'react'
// @ts-ignore
import { ScrollToTopButton, Sidebar } from '@shortlink-org/ui-kit'

import PushNotificationLayout from 'components/PushNotificationLayout'
import Footer from 'components/Footer'
import Header from 'components/Header'
import ory from '../pkg/sdk'
import { AxiosError } from 'axios'

// @ts-ignore
export function Layout({ children }) {
  const [session, setSession] = useState<string>('No valid Ory Session was found.\nPlease sign in to receive one.')
  const [hasSession, setHasSession] = useState<boolean>(false)
  const [open, setOpen] = useState(false)

  useEffect(() => {
    ory
      .toSession()
      .then(({ data }) => {
        setSession(JSON.stringify(data, null, 2))
        setHasSession(true)
      })
      .catch((err: AxiosError) =>
        // Something else happened!
        Promise.reject(err),
      )
  }, [])

  return (
    <PushNotificationLayout>
      <CssBaseline />

      <div className="grid grid-rows-[auto_1fr] h-screen overflow-hidden">
        <Header hasSession={hasSession} setOpen={() => setOpen(!open)} />

        <main className={'grid grid-cols-[auto_1fr] min-h-0'}>
          <div className={'h-full overflow-auto'}>
            {hasSession && <Sidebar mode={open ? 'full' : 'mini'} />}
          </div>

          <div className="overflow-auto h-full">
            {children}

            <Footer />

            <ScrollToTopButton />
          </div>
        </main>
      </div>
    </PushNotificationLayout>
  )
}
