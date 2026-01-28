'use client'

/**
 * Links Search Page with useDeferredValue
 * 
 * Features:
 * - ✅ Responsive search input
 * - ✅ Deferred results update
 * - ✅ Loading state handled locally
 * - ✅ ErrorBoundary for errors
 * - ✅ Old results stay visible while loading new ones
 */

import { Box, Typography, Card, CardContent, Skeleton, Alert } from '@mui/material'
import Link from 'next/link'

import { DeferredSearch } from '@/components/async'
import { LinksErrorBoundary } from '@/components/error'
import { useSearchLinksQuery } from '@/lib/datalayer'
import withAuthSync from '@/components/Private'
import PageHeader from '@/components/Page/Header'

/**
 * Component that reads search results via TanStack Query
 */
function SearchResults({ query }: { query: string }) {
  const { data, isLoading, error } = useSearchLinksQuery(query)
  const links = (data ?? []) as any[]

  if (!query.trim()) {
    return (
      <Box sx={{ py: 8, textAlign: 'center' }}>
        <Typography variant="h6" color="text.secondary">
          Start typing to search for links...
        </Typography>
        <Typography variant="body2" color="text.secondary" sx={{ mt: 1 }}>
          Try searching by URL, title, or description
        </Typography>
      </Box>
    )
  }
  
  if (error) {
    throw error
  }

  if (isLoading) {
    return <SearchSkeleton />
  }

  if (links.length === 0) {
    return (
      <Alert severity="info" sx={{ mt: 2 }}>
        No links found for <strong>"{query}"</strong>
      </Alert>
    )
  }
  
  return (
    <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2, mt: 2 }}>
      <Typography variant="body2" color="text.secondary">
        Found {links.length} link{links.length !== 1 ? 's' : ''} for "{query}"
      </Typography>
      
      {links.map((link: any) => (
        <Card key={link.id} variant="outlined" sx={{ '&:hover': { boxShadow: 2 } }}>
          <CardContent>
            <Typography variant="h6" gutterBottom>
              {link.title || 'Untitled Link'}
            </Typography>
            
            <Link 
              href={link.url} 
              target="_blank" 
              rel="noopener noreferrer"
              style={{ 
                color: '#3f51b5', 
                textDecoration: 'none',
                display: 'block',
                marginBottom: '8px',
                wordBreak: 'break-all'
              }}
            >
              {link.url}
            </Link>
            
            {link.description && (
              <Typography variant="body2" color="text.secondary">
                {link.description}
              </Typography>
            )}
            
            {link.shortUrl && (
              <Box sx={{ mt: 2, pt: 2, borderTop: 1, borderColor: 'divider' }}>
                <Typography variant="caption" color="text.secondary">
                  Short URL: {link.shortUrl}
                </Typography>
              </Box>
            )}
          </CardContent>
        </Card>
      ))}
    </Box>
  )
}

/**
 * Loading skeleton
 */
function SearchSkeleton() {
  return (
    <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2, mt: 2 }}>
      {[1, 2, 3].map((i) => (
        <Card key={i} variant="outlined">
          <CardContent>
            <Skeleton variant="text" width="60%" height={32} />
            <Skeleton variant="text" width="100%" />
            <Skeleton variant="text" width="80%" />
          </CardContent>
        </Card>
      ))}
    </Box>
  )
}

/**
 * Main search page component
 */
function LinksSearchPage() {
  return (
    <>
      <PageHeader title="Search Links" />
      
      <Box sx={{ p: 3, maxWidth: 1200, mx: 'auto' }}>
        <Typography variant="h4" gutterBottom>
          Search Links
        </Typography>
        
        <Typography variant="body1" color="text.secondary" sx={{ mb: 3 }}>
          Search through your shortened links by URL, title, or description
        </Typography>
        
        {/* DeferredSearch wraps the search input and results */}
        <LinksErrorBoundary>
          <DeferredSearch
            label="Search"
            placeholder="Type to search links..."
            fullWidth
            variant="outlined"
            loadingFallback={<Typography variant="caption">Searching...</Typography>}
          >
            {(query) => <SearchResults query={query} />}
          </DeferredSearch>
        </LinksErrorBoundary>
      </Box>
    </>
  )
}

export default withAuthSync(LinksSearchPage)
