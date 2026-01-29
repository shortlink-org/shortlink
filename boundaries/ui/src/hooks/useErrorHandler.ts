'use client'

import { useCallback, useRef } from 'react'
import { toast } from 'sonner'
import { LinkDomainError, LinkAction } from '@/domain/link/link.types'
import { useRouter } from 'next/navigation'

type ErrorHandlerOptions = {
  /** Max retry attempts for SERVICE_UNAVAILABLE errors */
  maxRetries?: number
  /** Delay between retries in ms */
  retryDelay?: number
  /** Callback when retry limit exceeded */
  onRetryExhausted?: () => void
  /** Custom action handlers */
  onLogin?: () => void
}

type RetryableFunction<T> = () => Promise<T>

/**
 * Hook for handling errors with toast notifications and auto-retry
 */
export function useErrorHandler(options: ErrorHandlerOptions = {}) {
  const { maxRetries = 3, retryDelay = 2000, onRetryExhausted, onLogin } = options

  const router = useRouter()
  const retryCountRef = useRef<Map<string, number>>(new Map())

  /**
   * Show error as toast notification
   */
  const showError = useCallback((error: LinkDomainError | string) => {
    if (typeof error === 'string') {
      toast.error(error, {
        duration: 5000,
      })
      return
    }

    const toastId = `error-${error.code}`

    toast.error(error.title, {
      id: toastId,
      description: error.detail,
      duration: error.action === 'RETRY' ? 8000 : 5000,
      action: getActionButton(error.action),
    })
  }, [])

  /**
   * Show success toast
   */
  const showSuccess = useCallback((message: string) => {
    toast.success(message, {
      duration: 3000,
    })
  }, [])

  /**
   * Get action button based on error action type
   */
  const getActionButton = useCallback(
    (action: LinkAction) => {
      switch (action) {
        case 'LOGIN':
          return {
            label: 'Sign in',
            onClick: () => {
              if (onLogin) {
                onLogin()
              } else {
                router.push('/auth/login')
              }
            },
          }
        case 'RETRY':
          return {
            label: 'Try again',
            onClick: () => {
              // This will be handled by the retry logic
              toast.dismiss()
            },
          }
        default:
          return undefined
      }
    },
    [router, onLogin],
  )

  /**
   * Handle error with appropriate action
   */
  const handleError = useCallback(
    (error: LinkDomainError) => {
      showError(error)

      // Handle LOGIN action
      if (error.action === 'LOGIN') {
        if (onLogin) {
          onLogin()
        } else {
          // Delay redirect to allow user to see the toast
          setTimeout(() => {
            router.push('/auth/login')
          }, 2000)
        }
      }
    },
    [showError, router, onLogin],
  )

  /**
   * Execute function with auto-retry for SERVICE_UNAVAILABLE errors
   */
  const withRetry = useCallback(
    async <T>(fn: RetryableFunction<T>, operationId: string = 'default'): Promise<T> => {
      const currentRetries = retryCountRef.current.get(operationId) || 0

      try {
        const result = await fn()
        // Reset retry count on success
        retryCountRef.current.delete(operationId)
        return result
      } catch (error: any) {
        const domainError = error as LinkDomainError

        // Check if error is retryable
        if (domainError.code === 'SERVICE_UNAVAILABLE' || domainError.code === 'NETWORK_ERROR') {
          if (currentRetries < maxRetries) {
            retryCountRef.current.set(operationId, currentRetries + 1)

            toast.loading(`Connection issue. Retrying... (${currentRetries + 1}/${maxRetries})`, {
              id: `retry-${operationId}`,
              duration: retryDelay,
            })

            await new Promise((resolve) => setTimeout(resolve, retryDelay))

            return withRetry(fn, operationId)
          } else {
            // Retry limit exceeded
            retryCountRef.current.delete(operationId)
            toast.dismiss(`retry-${operationId}`)

            if (onRetryExhausted) {
              onRetryExhausted()
            }

            toast.error('Service unavailable', {
              description: 'Please try again later. If the problem persists, contact support.',
              duration: 8000,
            })
          }
        }

        throw error
      }
    },
    [maxRetries, retryDelay, onRetryExhausted],
  )

  /**
   * Reset retry count for an operation
   */
  const resetRetries = useCallback((operationId: string = 'default') => {
    retryCountRef.current.delete(operationId)
  }, [])

  return {
    showError,
    showSuccess,
    handleError,
    withRetry,
    resetRetries,
  }
}

/**
 * Standalone function to show error toast (for use outside React components)
 */
export function showErrorToast(error: LinkDomainError | string) {
  if (typeof error === 'string') {
    toast.error(error, { duration: 5000 })
    return
  }

  toast.error(error.title, {
    description: error.detail,
    duration: 5000,
  })
}

/**
 * Standalone function to show success toast
 */
export function showSuccessToast(message: string) {
  toast.success(message, { duration: 3000 })
}

/**
 * Show loading toast that can be updated
 */
export function showLoadingToast(message: string, id?: string) {
  return toast.loading(message, { id })
}

/**
 * Update existing toast to success
 */
export function updateToastSuccess(id: string | number, message: string) {
  toast.success(message, { id, duration: 3000 })
}

/**
 * Update existing toast to error
 */
export function updateToastError(id: string | number, message: string) {
  toast.error(message, { id, duration: 5000 })
}

/**
 * Dismiss toast by id
 */
export function dismissToast(id?: string | number) {
  toast.dismiss(id)
}
