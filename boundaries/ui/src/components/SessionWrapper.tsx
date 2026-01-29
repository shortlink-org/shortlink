'use client'

/**
 * SessionWrapper - OPTIONAL session provider for public pages
 * 
 * Provides session context for public pages
 * Uses optional session query to avoid hard auth requirements
 */

import React, { ReactNode } from 'react'
import { SessionProvider } from '@/contexts/SessionContext'
import { useOptionalSessionQuery } from '@/lib/datalayer'

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
  const { data: session, isLoading } = useOptionalSessionQuery()

  return (
    <SessionProvider session={session ?? null} isLoading={isLoading}>
      {children}
    </SessionProvider>
  )
}

export default SessionWrapper
