'use client'

/**
 * SessionWrapper - OPTIONAL session provider for public pages
 * 
 * Provides empty session context so useSession() hook works everywhere
 * WITHOUT actually fetching session (для публичных страниц)
 */

import React, { ReactNode } from 'react'
import { SessionProvider } from '@/contexts/SessionContext'

/**
 * SessionWrapper - Provides empty session for public pages
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
  // Provide empty session context
  // Protected pages will handle their own session check
  return (
    <SessionProvider session={null}>
      {children}
    </SessionProvider>
  )
}

export default SessionWrapper

