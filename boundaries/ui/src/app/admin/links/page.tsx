'use client'

/**
 * Admin Links Page - Migrated to React 19
 * 
 * Changes:
 * - ✅ Replaced Redux + useEffect with TanStack Query
 * - ✅ Added ErrorBoundary for error handling
 * - ✅ Added skeleton loader instead of spinner
 * - ✅ Data cached automatically (1 minute TTL)
 * - ✅ No manual loading state management
 * 
 * Old version backed up in git history
 */


import Statistic from '@/components/Dashboard/stats'
import withAuthSync from '@/components/Private'
import Header from '@/components/Page/Header'
import AdminUserLinksTable from '@/components/Page/admin/user/linksTable'
import { LinksTableSkeleton } from '@/components/Skeleton'
import { LinksErrorBoundary } from '@/components/error'
import { useLinksListQuery } from '@/lib/datalayer'
import { protoTimestampToIsoString } from '@/lib/time'

/**
 * Component that reads links data via TanStack Query
 * Loading state handled locally
 */
function AdminLinksData() {
  const { data, isLoading, error } = useLinksListQuery()
  const links = (data ?? []) as any[]

  if (error) {
    throw error
  }

  if (isLoading) {
    return <LinksTableSkeleton />
  }

  // Transform data for table
  const tableData = links.map((link: any) => ({
    url: link.url || '',
    hash: link.hash || '',
    describe: link.describe,
    created_at: protoTimestampToIsoString(link.created_at),
    updated_at: protoTimestampToIsoString(link.updated_at),
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
        <AdminLinksData />
      </LinksErrorBoundary>
    </>
  )
}

export default withAuthSync(AdminLinkTable)
