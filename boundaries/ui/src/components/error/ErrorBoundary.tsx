'use client'

import React, { ReactNode, useRef, useState } from 'react'
import { Alert, AlertTitle, Button, Box } from '@mui/material'
import { ErrorBoundary as UiKitErrorBoundary } from '@shortlink-org/ui-kit'

interface Props {
  children: ReactNode
  fallback?: ReactNode
  onError?: (error: Error, errorInfo: React.ErrorInfo) => void
  maxRetries?: number
  onRetry?: () => void
}

type FallbackRenderArgs = {
  error: Error
  resetErrorBoundary: () => void
}

/**
 * App error boundary built on top of the shared ui-kit runtime boundary.
 *
 * Usage:
 * ```tsx
 * <ErrorBoundary fallback={<div>Something went wrong</div>}>
 *   <MyComponent />
 * </ErrorBoundary>
 * ```
 */
export function ErrorBoundary({ children, fallback, onError, maxRetries = 3, onRetry }: Props) {
  const [retryCount, setRetryCount] = useState(0)
  const latestRetryCount = useRef(0)

  const handleRetry = (resetErrorBoundary: () => void) => {
    if (latestRetryCount.current >= maxRetries) return

    latestRetryCount.current += 1
    setRetryCount(latestRetryCount.current)
    onRetry?.()
    resetErrorBoundary()
  }

  const handleReload = () => {
    window.location.reload()
  }

  const renderDefaultFallback = ({ error, resetErrorBoundary }: FallbackRenderArgs) => {
    if (fallback) {
      return fallback
    }

    const canRetry = retryCount < maxRetries

    return (
      <Box sx={{ p: 3 }}>
        <Alert severity="error">
          <AlertTitle>Oops! Something went wrong</AlertTitle>
          {error.message || 'An unexpected error occurred'}

          {retryCount > 0 && (
            <Box sx={{ mt: 1, fontSize: '0.875rem', opacity: 0.8 }}>
              Retry attempt: {retryCount} / {maxRetries}
            </Box>
          )}

          <Box sx={{ mt: 2, display: 'flex', gap: 1 }}>
            {canRetry && (
              <Button variant="outlined" size="small" onClick={() => handleRetry(resetErrorBoundary)}>
                Try Again
              </Button>
            )}

            <Button variant="contained" size="small" onClick={handleReload}>
              Reload Page
            </Button>
          </Box>
        </Alert>
      </Box>
    )
  }

  return (
    <UiKitErrorBoundary
      onError={onError}
      onReset={() => {
        latestRetryCount.current = retryCount
      }}
      fallbackRender={renderDefaultFallback}
    >
      {children}
    </UiKitErrorBoundary>
  )
}

/**
 * Specialized Error Boundaries
 */

export function ProfileErrorBoundary({ children }: { children: ReactNode }) {
  return (
    <ErrorBoundary
      fallback={
        <Alert severity="error">
          <AlertTitle>Failed to load profile</AlertTitle>
          Please try refreshing the page or contact support if the problem persists.
        </Alert>
      }
    >
      {children}
    </ErrorBoundary>
  )
}

export function LinksErrorBoundary({ children }: { children: ReactNode }) {
  return (
    <ErrorBoundary
      fallback={
        <Alert severity="error">
          <AlertTitle>Failed to load links</AlertTitle>
          Unable to fetch your links. Please try again later.
        </Alert>
      }
    >
      {children}
    </ErrorBoundary>
  )
}

export function FormErrorBoundary({ children }: { children: ReactNode }) {
  return (
    <ErrorBoundary
      fallback={
        <Alert severity="error">
          <AlertTitle>Form Error</AlertTitle>
          Something went wrong while processing your form. Please try again.
        </Alert>
      }
    >
      {children}
    </ErrorBoundary>
  )
}

export default ErrorBoundary

