import { Box, CircularProgress, CircularProgressProps } from '@mui/material'
import React from 'react'

interface LoadingSpinnerProps {
  size?: CircularProgressProps['size']
  minHeight?: string | number
  fullScreen?: boolean
}

export default function LoadingSpinner({ 
  size = 40, 
  minHeight = '200px',
  fullScreen = false 
}: LoadingSpinnerProps) {
  const containerStyle = fullScreen
    ? {
        position: 'fixed' as const,
        top: 0,
        left: 0,
        right: 0,
        bottom: 0,
        zIndex: 9999,
      }
    : {
        minHeight,
      }

  return (
    <Box
      display="flex"
      justifyContent="center"
      alignItems="center"
      sx={containerStyle}
    >
      <CircularProgress size={size} />
    </Box>
  )
}

