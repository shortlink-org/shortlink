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
 * - ✅ SearchForm (desktop) + Drawer + SearchForm (mobile); DataTable keeps pagination & column filters
 *
 * Old version backed up in git history
 */

import { useDeferredValue, useMemo, useState } from 'react'
import { useQueryClient } from '@tanstack/react-query'
import { Button, Drawer, SearchForm } from '@shortlink-org/ui-kit'

import Statistic from '@/components/Dashboard/stats'
import withAuthSync from '@/components/Private'
import Header from '@/components/Page/Header'
import PageSection from '@/components/Page/Section'
import UserLinksTable from '@/components/Page/user/linksTable'
import { LinksTableSkeleton } from '@/components/Skeleton'
import { LinksErrorBoundary } from '@/components/error'
import { queryKeys, useLinksListQuery } from '@/lib/datalayer'
import { filterLinksBySearch } from '@/lib/filterLinksTable'
import { protoTimestampToIsoString } from '@/lib/time'

/**
 * Component that reads links data via TanStack Query
 * Loading state handled locally
 */
function LinksData() {
  const queryClient = useQueryClient()
  const { data, isLoading, error } = useLinksListQuery()
  const links = (data ?? []) as any[]
  const [search, setSearch] = useState('')
  const deferredSearch = useDeferredValue(search)
  const [drawerOpen, setDrawerOpen] = useState(false)

  if (error) {
    throw error
  }

  if (isLoading) {
    return <LinksTableSkeleton />
  }

  const allRows = links.map((link: any) => ({
    url: link.url || '',
    hash: link.hash || '',
    describe: link.describe,
    created_at: protoTimestampToIsoString(link.created_at),
    updated_at: protoTimestampToIsoString(link.updated_at),
  }))

  const tableData = useMemo(
    () => filterLinksBySearch(allRows, deferredSearch),
    [allRows, deferredSearch],
  )

  const handleRefresh = () => {
    void queryClient.invalidateQueries({ queryKey: queryKeys.linksList() })
  }

  const searchForm = (
    <SearchForm
      label="Search links"
      placeholder="Search by URL, hash, description…"
      value={search}
      onValueChange={setSearch}
      onSearch={setSearch}
      fullWidth
    />
  )

  return (
    <>
      <Statistic count={links.length} />
      <div className="mb-4 flex flex-col gap-3 sm:flex-row sm:items-start sm:justify-between">
        <div className="hidden min-w-0 flex-1 md:block md:max-w-xl">{searchForm}</div>
        <Button
          type="button"
          variant="secondary"
          className="w-full md:hidden"
          onClick={() => setDrawerOpen(true)}
        >
          Search links
        </Button>
      </div>
      {search.trim() ? (
        <p className="mb-3 text-sm text-[var(--color-muted-foreground)]">
          Showing {tableData.length} of {links.length} links
        </p>
      ) : null}
      <Drawer
        open={drawerOpen}
        onClose={setDrawerOpen}
        position="bottom"
        title="Search links"
      >
        {searchForm}
        <p className="mt-4 text-sm text-gray-600 dark:text-gray-400">
          Use column headers in the table for per-column filters.
        </p>
      </Drawer>
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
        <PageSection className="pb-10">
          <LinksData />
        </PageSection>
      </LinksErrorBoundary>
    </>
  )
}

export default withAuthSync(LinkTable)
