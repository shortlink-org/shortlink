'use client'

/**
 * Admin Links Page - Migrated to React 19
 * 
 * Changes:
 * - ✅ Replaced Redux + useEffect with use() + Suspense
 * - ✅ Added ErrorBoundary for error handling
 * - ✅ Added skeleton loader instead of spinner
 * - ✅ Data cached automatically (1 minute TTL)
 * - ✅ No manual loading state management
 * 
 * Old version backed up in git history
 */

import { use, Suspense } from 'react'

import Statistic from '@/components/Dashboard/stats'
import withAuthSync from '@/components/Private'
import Header from '@/components/Page/Header'
import AdminUserLinksTable from '@/components/Page/admin/user/linksTable'
import { LinksTableSkeleton } from '@/components/Skeleton'
import { LinksErrorBoundary } from '@/components/error'
import { fetchLinksList, useInvalidateCache } from '@/lib/data'

/**
 * Component that reads links data using use()
 * Automatically suspends while data is loading
 */
function AdminLinksData() {
  // use() reads the promise and suspends the component
  const links = use(fetchLinksList())
  const { invalidate } = useInvalidateCache()

  // Transform data for table
  const tableData = links.map((link: any) => ({
    url: link.url || '',
    hash: link.hash || '',
    describe: link.describe,
    created_at: link.created_at
      ? new Date((link.created_at.seconds || 0) * 1000 + (link.created_at.nanos || 0) / 1000000).toISOString()
      : '',
    updated_at: link.updated_at
      ? new Date((link.updated_at.seconds || 0) * 1000 + (link.updated_at.nanos || 0) / 1000000).toISOString()
      : '',
  }))

  return (
    <>
      <Statistic count={links.length} />
      <AdminUserLinksTable data={tableData} />
    </>
  )
}

/**
 * Main component with declarative async management
 */
function AdminLinkTable() {
  return (
    <>
      <Header title="Admin Links" />

      {/* ErrorBoundary catches errors */}
      <LinksErrorBoundary>
        {/* Suspense shows skeleton while data loads */}
        <Suspense fallback={<LinksTableSkeleton />}>
          <AdminLinksData />
        </Suspense>
      </LinksErrorBoundary>
    </>
  )
}

export default withAuthSync(AdminLinkTable)
