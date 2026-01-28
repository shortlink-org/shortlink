/**
 * withAuthSync HOC - Fetches and provides session to protected pages
 * 
 * Changes:
 * - ✅ Uses TanStack Query + fetchSession from data layer
 * - ✅ Uses explicit loading state
 * - ✅ Uses ErrorBoundary for error handling
 * - ✅ Provides session via SessionProvider
 * - ✅ No manual useState/useEffect
 */

'use client'

import React, { ReactNode } from 'react'
import { useRouter } from 'next/navigation'
import { SessionProvider } from '@/contexts/SessionContext'
import { useSessionQuery } from '@/lib/datalayer'

/**
 * Component that fetches session for protected page
 */
function ProtectedPageData<P extends object>({ 
  Component, 
  props 
}: { 
  Component: React.ComponentType<P>
  props: P 
}) {
  const { data: session, isLoading, error } = useSessionQuery()

  if (error) {
    throw error
  }

  if (isLoading) {
    return <div>Loading...</div>
  }

  if (!session) {
    return null
  }
  
  // Session exists - render protected page
  return (
    <SessionProvider session={session}>
      {React.createElement(Component, props)}
    </SessionProvider>
  )
}

/**
 * ErrorBoundary for auth errors
 */
class AuthErrorBoundary extends React.Component<
  { children: ReactNode; router: any },
  { hasError: boolean; error: Error | null }
> {
  constructor(props: { children: ReactNode; router: any }) {
    super(props)
    this.state = { hasError: false, error: null }
  }

  static getDerivedStateFromError(error: Error) {
    return { hasError: true, error }
  }

  componentDidCatch(error: Error) {
    console.log('Auth error:', error.message)
    
    // Redirect to login
    this.props.router.push('/auth/login')
  }

  render() {
    if (this.state.hasError) {
      // While redirecting, show nothing
      return <div>Redirecting to login...</div>
    }

    return this.props.children
  }
}

export default function withAuthSync<P extends object>(Child: React.ComponentType<P>) {
  return function WrappedComponent(props: P) {
    const router = useRouter()
    
    return (
      <AuthErrorBoundary router={router}>
        <ProtectedPageData Component={Child} props={props} />
      </AuthErrorBoundary>
    )
  }
}
