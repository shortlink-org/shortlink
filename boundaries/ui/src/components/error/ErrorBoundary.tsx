'use client'

import React, { Component, ReactNode } from 'react'
import { Alert, AlertTitle, Button, Box } from '@mui/material'

interface Props {
  children: ReactNode
  fallback?: ReactNode
  onError?: (error: Error, errorInfo: React.ErrorInfo) => void
  maxRetries?: number
  onRetry?: () => void
}

interface State {
  hasError: boolean
  error?: Error
  retryCount: number
}

/**
 * ErrorBoundary component для обработки ошибок в React 19
 * 
 * Использование:
 * ```tsx
 * <ErrorBoundary fallback={<div>Something went wrong</div>}>
 *   <MyComponent />
 * </ErrorBoundary>
 * ```
 */
export class ErrorBoundary extends Component<Props, State> {
  constructor(props: Props) {
    super(props)
    this.state = { hasError: false, retryCount: 0 }
  }

  static getDerivedStateFromError(error: Error): Partial<State> {
    return { hasError: true, error }
  }

  componentDidCatch(error: Error, errorInfo: React.ErrorInfo) {
    // Логируем ошибку
    console.error('ErrorBoundary caught an error:', error, errorInfo)
    
    // Вызываем callback если передан
    this.props.onError?.(error, errorInfo)
    
    // Можно отправить в Sentry/LogRocket
    // if (typeof window !== 'undefined' && window.Sentry) {
    //   window.Sentry.captureException(error, { contexts: { react: errorInfo } })
    // }
  }

  handleReset = () => {
    const maxRetries = this.props.maxRetries ?? 3
    const canRetry = this.state.retryCount < maxRetries
    
    if (canRetry) {
      this.setState(prev => ({ 
        hasError: false, 
        error: undefined,
        retryCount: prev.retryCount + 1
      }))
      this.props.onRetry?.()
    }
  }
  
  handleReload = () => {
    window.location.reload()
  }

  render() {
    if (this.state.hasError) {
      // Если передан custom fallback, используем его
      if (this.props.fallback) {
        return this.props.fallback
      }

      // Дефолтный fallback
      const maxRetries = this.props.maxRetries ?? 3
      const canRetry = this.state.retryCount < maxRetries
      
      return (
        <Box sx={{ p: 3 }}>
          <Alert severity="error">
            <AlertTitle>Oops! Something went wrong</AlertTitle>
            {this.state.error?.message || 'An unexpected error occurred'}
            
            {this.state.retryCount > 0 && (
              <Box sx={{ mt: 1, fontSize: '0.875rem', opacity: 0.8 }}>
                Retry attempt: {this.state.retryCount} / {maxRetries}
              </Box>
            )}
            
            <Box sx={{ mt: 2, display: 'flex', gap: 1 }}>
              {canRetry && (
                <Button 
                  variant="outlined" 
                  size="small" 
                  onClick={this.handleReset}
                >
                  Try Again
                </Button>
              )}
              
              <Button 
                variant="contained" 
                size="small" 
                onClick={this.handleReload}
              >
                Reload Page
              </Button>
            </Box>
          </Alert>
        </Box>
      )
    }

    return this.props.children
  }
}

/**
 * Специализированные Error Boundaries
 */

// Для данных профиля
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

// Для списка ссылок
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

// Для форм
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

