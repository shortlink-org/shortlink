/**
 * Layout - Updated with NavigationProvider and SessionContext
 * 
 * Changes:
 * - ✅ Added NavigationProvider for global progress bar
 * - ✅ Content fades slightly during navigation
 * - ✅ Progress bar shows at top during page transitions
 * - ✅ Removed duplicate session check (uses SessionContext)
 */

import CssBaseline from '@mui/material/CssBaseline'
import React, { useState } from 'react'
// @ts-ignore
import { ScrollToTopButton, Sidebar } from '@shortlink-org/ui-kit'

// import PushNotificationLayout from '@/components/PushNotificationLayout'
import Footer from '@/components/Footer'
import Header from '@/components/Header'
import { NavigationProvider } from '@/components/Navigation'

// @ts-ignore
export function Layout({ children }: { children: React.ReactNode }) {
  const [open, setOpen] = useState(false)
  
  // Try to get session from context (may be null for public pages)
  let hasSession = false
  try {
    const { useSession } = require('@/contexts/SessionContext')
    const session = useSession()
    hasSession = session?.hasSession || false
  } catch {
    // Context not available - public page
    hasSession = false
  }

  return (
    <NavigationProvider>
      <CssBaseline />

      <div className="grid grid-rows-[auto_1fr] h-screen overflow-hidden">
        <Header hasSession={hasSession} setOpen={() => setOpen(!open)} />

        <main className={'grid grid-cols-[auto_1fr] min-h-0'}>
          <div className={'h-full overflow-auto'}>{hasSession && <Sidebar mode={open ? 'full' : 'mini'} />}</div>

          <div className="overflow-auto h-full">
            {children}

            <Footer />

            <ScrollToTopButton />
          </div>
        </main>
      </div>
    </NavigationProvider>
  )
}
