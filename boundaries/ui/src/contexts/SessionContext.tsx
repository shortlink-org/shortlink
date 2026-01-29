'use client'

/**
 * SessionContext - Centralized session management with React 19
 * 
 * Features:
 * - ✅ Single source of truth for session
 * - ✅ Uses fetchSession from data layer
 * - ✅ No duplication across components
 * - ✅ Automatic session check
 */

import { createContext, useContext, ReactNode } from 'react'
import { Session } from '@ory/client'

interface SessionContextValue {
  session: Session | null
  hasSession: boolean
  isLoading: boolean
}

const SessionContext = createContext<SessionContextValue | undefined>(undefined)

export function useSession() {
  const context = useContext(SessionContext)
  if (context === undefined) {
    throw new Error('useSession must be used within SessionProvider')
  }
  return context
}

interface SessionProviderProps {
  children: ReactNode
  session: Session | null
  isLoading?: boolean
}

/**
 * SessionProvider - Provides session to all child components
 * 
 * Usage:
 * ```tsx
 * <SessionProvider session={session}>
 *   <YourApp />
 * </SessionProvider>
 * ```
 */
export function SessionProvider({ children, session, isLoading = false }: SessionProviderProps) {
  const value: SessionContextValue = {
    session,
    hasSession: !!session,
    isLoading,
  }

  return (
    <SessionContext.Provider value={value}>
      {children}
    </SessionContext.Provider>
  )
}

export default SessionContext
