'use client'

import { Box, Card, CardContent, Skeleton, Stack } from '@mui/material'

/**
 * Skeleton loader for Profile page
 * 
 * Shows while profile data is loading
 */
export function ProfileSkeleton() {
  return (
    <Box sx={{ p: 3, maxWidth: 800, mx: 'auto' }}>
      {/* Header */}
      <Box sx={{ mb: 4 }}>
        <Skeleton variant="text" width={200} height={40} sx={{ mb: 1 }} />
        <Skeleton variant="text" width={300} height={24} />
      </Box>

      {/* Profile Info Card */}
      <Card variant="outlined" sx={{ mb: 3 }}>
        <CardContent>
          <Stack spacing={3}>
            {/* Avatar + Name */}
            <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
              <Skeleton variant="circular" width={80} height={80} />
              <Box sx={{ flex: 1 }}>
                <Skeleton variant="text" width="60%" height={32} />
                <Skeleton variant="text" width="40%" height={24} />
              </Box>
            </Box>

            {/* Divider */}
            <Box sx={{ height: 1, bgcolor: 'divider' }} />

            {/* Info rows */}
            {Array.from({ length: 4 }).map((_, i) => (
              <Box key={i} sx={{ display: 'flex', gap: 2 }}>
                <Skeleton variant="text" width="30%" height={24} />
                <Skeleton variant="text" width="70%" height={24} />
              </Box>
            ))}
          </Stack>
        </CardContent>
      </Card>

      {/* Personal Info Form Card */}
      <Card variant="outlined" sx={{ mb: 3 }}>
        <CardContent>
          <Skeleton variant="text" width={150} height={32} sx={{ mb: 3 }} />
          
          <Stack spacing={2.5}>
            {/* Form fields */}
            {Array.from({ length: 5 }).map((_, i) => (
              <Box key={i}>
                <Skeleton variant="text" width={120} height={20} sx={{ mb: 0.5 }} />
                <Skeleton variant="rectangular" width="100%" height={56} sx={{ borderRadius: 1 }} />
              </Box>
            ))}

            {/* Buttons */}
            <Box sx={{ display: 'flex', gap: 2, justifyContent: 'flex-end', mt: 2 }}>
              <Skeleton variant="rectangular" width={100} height={36} sx={{ borderRadius: 1 }} />
              <Skeleton variant="rectangular" width={100} height={36} sx={{ borderRadius: 1 }} />
            </Box>
          </Stack>
        </CardContent>
      </Card>

      {/* Additional sections */}
      <Card variant="outlined">
        <CardContent>
          <Skeleton variant="text" width={180} height={32} sx={{ mb: 2 }} />
          <Stack spacing={2}>
            <Skeleton variant="rectangular" width="100%" height={60} sx={{ borderRadius: 1 }} />
            <Skeleton variant="rectangular" width="100%" height={60} sx={{ borderRadius: 1 }} />
          </Stack>
        </CardContent>
      </Card>
    </Box>
  )
}

export default ProfileSkeleton

