'use client'

/**
 * SessionWrapper - OPTIONAL session provider for public pages
 *
 * Provides session context for public pages
 * Uses optional session query to avoid hard auth requirements
 * Supports static export (deferred client-side rendering)
 */

import React, { ReactNode, useState, useEffect } from 'react'
import { SessionProvider } from '@/contexts/SessionContext'
import { useOptionalSessionQuery } from '@/lib/datalayer'

/**
 * Inner component that uses React Query (only rendered on client)
 */
function SessionContent({ children }: { children: ReactNode }) {
  const { data: session, isLoading } = useOptionalSessionQuery()

  return (
    <SessionProvider session={session ?? null} isLoading={isLoading}>
      {children}
    </SessionProvider>
  )
}

/**
 * SessionWrapper - Provides optional session for public pages
 *
 * Usage:
 * ```tsx
 * <SessionWrapper>
 *   <PublicPages />
 * </SessionWrapper>
 * ```
 *
 * Note: Protected pages will fetch session via withAuthSync HOC
 */
export function SessionWrapper({ children }: { children: ReactNode }) {
  const [mounted, setMounted] = useState(false)

  useEffect(() => {
    setMounted(true)
  }, [])

  // During SSR/static export, render children without session
  if (!mounted) {
    return (
      <SessionProvider session={null} isLoading={true}>
        {children}
      </SessionProvider>
    )
  }

  return <SessionContent>{children}</SessionContent>
}

export default SessionWrapper
