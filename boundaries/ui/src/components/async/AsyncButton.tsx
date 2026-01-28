'use client'

import { useTransition, ReactNode } from 'react'
import { Button, CircularProgress } from '@mui/material'
import type { ButtonProps } from '@mui/material/Button'

interface AsyncButtonProps extends Omit<ButtonProps, 'onClick' | 'action'> {
  children: ReactNode
  /** Async action to execute */
  action?: () => Promise<void>
  /** Alternative: sync onClick handler wrapped in transition */
  onClick?: () => void
  /** Show loading text instead of children when pending */
  loadingText?: string
  /** Disable button when pending (default: true) */
  disableOnPending?: boolean
}

/**
 * AsyncButton - кнопка с автоматическим управлением pending состоянием
 * 
 * Использует useTransition для оборачивания async операций
 * 
 * @example
 * ```tsx
 * <AsyncButton 
 *   action={async () => await saveData()} 
 *   loadingText="Saving..."
 * >
 *   Save
 * </AsyncButton>
 * ```
 */
export function AsyncButton({
  children,
  action,
  onClick,
  loadingText,
  disableOnPending = true,
  disabled,
  ...buttonProps
}: AsyncButtonProps) {
  const [isPending, startTransition] = useTransition()

  const handleClick = () => {
    startTransition(async () => {
      try {
        if (action) {
          await action()
        } else if (onClick) {
          onClick()
        }
      } catch (error) {
        // Ошибки обрабатываются ErrorBoundary
        throw error
      }
    })
  }

  return (
    <Button
      {...buttonProps}
      onClick={handleClick}
      disabled={disabled || (disableOnPending && isPending)}
      startIcon={isPending ? <CircularProgress size={16} /> : buttonProps.startIcon}
      sx={{
        opacity: isPending ? 0.7 : 1,
        transition: 'opacity 0.2s',
        ...buttonProps.sx
      }}
    >
      {isPending && loadingText ? loadingText : children}
    </Button>
  )
}

export default AsyncButton
