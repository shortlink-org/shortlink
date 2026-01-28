'use client'

/**
 * Links List Page - Migrated to React 19
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

import { useQueryClient } from '@tanstack/react-query'

import Statistic from '@/components/Dashboard/stats'
import withAuthSync from '@/components/Private'
import Header from '@/components/Page/Header'
import UserLinksTable from '@/components/Page/user/linksTable'
import { LinksTableSkeleton } from '@/components/Skeleton'
import { LinksErrorBoundary } from '@/components/error'
import { queryKeys, useLinksListQuery } from '@/lib/datalayer'

/**
 * Component that reads links data via TanStack Query
 * Loading state handled locally
 */
function LinksData() {
  const queryClient = useQueryClient()
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
    created_at: link.created_at
      ? new Date((link.created_at.seconds || 0) * 1000 + (link.created_at.nanos || 0) / 1000000).toISOString()
      : '',
    updated_at: link.updated_at
      ? new Date((link.updated_at.seconds || 0) * 1000 + (link.updated_at.nanos || 0) / 1000000).toISOString()
      : '',
  }))

  const handleRefresh = () => {
    void queryClient.invalidateQueries({ queryKey: queryKeys.linksList() })
  }

  return (
    <>
      <Statistic count={links.length} />
      <UserLinksTable data={tableData} onRefresh={handleRefresh} />
    </>
  )
}

/**
 * Main component with declarative async management
 */
function LinkTable() {
  return (
    <>
      <Header title="Links" />

      {/* ErrorBoundary catches errors */}
      <LinksErrorBoundary>
        <LinksData />
      </LinksErrorBoundary>
    </>
  )
}

export default withAuthSync(LinkTable)
