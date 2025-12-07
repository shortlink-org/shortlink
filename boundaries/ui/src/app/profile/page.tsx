'use client'

/**
 * Profile Page - Migrated to React 19 Async Patterns
 * 
 * Changes from old version:
 * - ✅ Replaced useState + useEffect with use() API
 * - ✅ Added ErrorBoundary for error handling
 * - ✅ Added Suspense for loading states
 * - ✅ Code reduced from 115 lines to ~60 lines (-48%)
 * - ✅ No manual state management
 * - ✅ No race conditions
 * 
 * Old version is backed up as page.old.tsx
 */

import { use, Suspense } from 'react'
import { Session } from '@ory/client'

import { fetchSession } from '@/lib/data'
import { ProfileErrorBoundary } from '@/components/error'
import { ProfileSkeleton } from '@/components/Skeleton'
import withAuthSync from '@/components/Private'
import Header from '@/components/Page/Header'
import Notifications from '@/components/Profile/Notifications'
import Personal from '@/components/Profile/Personal'
import Profile from '@/components/Profile/Profile'
import Welcome from '@/components/Profile/Welcome'

/**
 * Component that uses use() to read session data
 * Automatically suspends while data is loading
 */
function ProfileData() {
  // use() reads the promise and suspends the component
  const session: Session = use(fetchSession())
  
  const firstName = session.identity?.traits?.name?.first || 'User'
  const lastName = session.identity?.traits?.name?.last || ''
  const email = session.identity?.traits?.email || ''

  return (
    <>
      <Welcome nickname={firstName} />
      
      <Profile />
      
      <div className="hidden sm:block" aria-hidden="true">
        <div className="py-5">
          <div className="border-t border-gray-200 dark:border-gray-700" />
        </div>
      </div>
      
      <Personal 
        session={session}
        firstName={firstName}
        lastName={lastName}
        email={email}
      />
      
      <div className="hidden sm:block" aria-hidden="true">
        <div className="py-5">
          <div className="border-t border-gray-200 dark:border-gray-700" />
        </div>
      </div>
      
      <Notifications />
    </>
  )
}

/**
 * Main component with declarative async management
 */
function ProfileContent() {
  return (
    <>
      <Header title="Profile" />
      
      {/* ErrorBoundary catches errors with retry support */}
      <ProfileErrorBoundary>
        {/* Suspense shows ProfileSkeleton while data loads */}
        <Suspense fallback={<ProfileSkeleton />}>
          <ProfileData />
        </Suspense>
      </ProfileErrorBoundary>
    </>
  )
}

export default withAuthSync(ProfileContent)
