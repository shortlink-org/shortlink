'use client'

/**
 * Security Page - Simplified with centralized session
 * 
 * Changes:
 * - ✅ Removed duplicate useEffect + useState
 * - ✅ Uses centralized SessionContext
 * - ✅ No session check duplication
 */

import React from 'react'
import Header from '@/components/Page/Header'
import PageSection from '@/components/Page/Section'

import { Layout } from '@/components'
import withAuthSync from '@/components/Private'
import Security from '@/components/Profile/Security'
import { useSession } from '@/contexts/SessionContext'

// <NextSeo
// title="Security"
// description="Security page for your account."
// openGraph={{
//   title: 'Security',
//     description: 'Security page for your account.',
//     type: 'website',
// }}
// />

function SecurityContent() {
  // Get session from centralized context (no duplication!)
  const { session, hasSession } = useSession()

  return (
    <Layout>
      <Header title="Security" description="Review your authentication settings and keep your account protected." />

      <PageSection className="pb-10">
        <Security />
      </PageSection>
    </Layout>
  )
}

export default withAuthSync(SecurityContent)
