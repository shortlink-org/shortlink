'use client'

import { Box, Card, CardContent, Skeleton } from '@mui/material'

/**
 * Skeleton loader for links table
 * 
 * Shows while links data is loading
 */
export function LinksTableSkeleton({ rows = 5 }: { rows?: number }) {
  return (
    <Box sx={{ p: 3 }}>
      {/* Statistics skeleton */}
      <Box sx={{ mb: 3, display: 'flex', gap: 2 }}>
        <Skeleton variant="rectangular" width={200} height={80} sx={{ borderRadius: 2 }} />
        <Skeleton variant="rectangular" width={200} height={80} sx={{ borderRadius: 2 }} />
      </Box>

      {/* Table header skeleton */}
      <Card variant="outlined" sx={{ mb: 2 }}>
        <CardContent>
          <Box sx={{ display: 'flex', gap: 2, mb: 2 }}>
            <Skeleton variant="text" width="25%" height={24} />
            <Skeleton variant="text" width="15%" height={24} />
            <Skeleton variant="text" width="30%" height={24} />
            <Skeleton variant="text" width="15%" height={24} />
            <Skeleton variant="text" width="15%" height={24} />
          </Box>

          {/* Table rows skeleton */}
          {Array.from({ length: rows }).map((_, i) => (
            <Box key={i} sx={{ display: 'flex', gap: 2, mb: 1.5 }}>
              <Skeleton variant="text" width="25%" />
              <Skeleton variant="text" width="15%" />
              <Skeleton variant="text" width="30%" />
              <Skeleton variant="text" width="15%" />
              <Skeleton variant="text" width="15%" />
            </Box>
          ))}
        </CardContent>
      </Card>

      {/* Actions skeleton */}
      <Box sx={{ display: 'flex', gap: 2, justifyContent: 'flex-end' }}>
        <Skeleton variant="rectangular" width={100} height={36} sx={{ borderRadius: 1 }} />
        <Skeleton variant="rectangular" width={100} height={36} sx={{ borderRadius: 1 }} />
      </Box>
    </Box>
  )
}

export default LinksTableSkeleton

