'use client'

import { useTransition, FormEvent, ReactNode } from 'react'
import { Box, LinearProgress } from '@mui/material'

interface AsyncFormProps {
  children: ReactNode
  /** Async action to execute on submit */
  onSubmit: (formData: FormData) => Promise<void>
  /** Show progress bar when pending */
  showProgress?: boolean
  /** Additional form props */
  [key: string]: any
}

/**
 * AsyncForm - форма с автоматическим управлением pending состоянием
 * 
 * Использует useTransition для оборачивания submit action
 * 
 * @example
 * ```tsx
 * <AsyncForm onSubmit={async (formData) => await saveProfile(formData)}>
 *   <input name="name" />
 *   <AsyncButton type="submit">Save</AsyncButton>
 * </AsyncForm>
 * ```
 */
export function AsyncForm({
  children,
  onSubmit,
  showProgress = true,
  ...formProps
}: AsyncFormProps) {
  const [isPending, startTransition] = useTransition()

  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    
    const formData = new FormData(e.currentTarget)
    
    startTransition(async () => {
      try {
        await onSubmit(formData)
      } catch (error) {
        // Ошибки обрабатываются ErrorBoundary
        throw error
      }
    })
  }

  return (
    <Box sx={{ position: 'relative' }}>
      {showProgress && isPending && (
        <LinearProgress 
          sx={{ 
            position: 'absolute', 
            top: 0, 
            left: 0, 
            right: 0,
            zIndex: 1
          }} 
        />
      )}
      <form
        {...formProps}
        onSubmit={handleSubmit}
        style={{
          opacity: isPending ? 0.7 : 1,
          transition: 'opacity 0.2s',
          ...formProps.style
        }}
      >
        <fieldset disabled={isPending} style={{ border: 'none', padding: 0, margin: 0 }}>
          {children}
        </fieldset>
      </form>
    </Box>
  )
}

export default AsyncForm

