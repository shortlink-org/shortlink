'use client'

import { useState, useDeferredValue, Suspense, use, ReactNode } from 'react'
import { TextField, Box, CircularProgress } from '@mui/material'
import type { TextFieldProps } from '@mui/material/TextField'

interface DeferredSearchProps extends Omit<TextFieldProps, 'value' | 'onChange'> {
  /** Initial search query */
  initialQuery?: string
  /** Render function that receives deferred query */
  children: (query: string) => ReactNode
  /** Placeholder for loading state */
  loadingFallback?: ReactNode
  /** Debounce delay (default: 300ms) */
  debounceMs?: number
}

/**
 * DeferredSearch - поисковое поле с useDeferredValue
 * 
 * Input всегда responsive, результаты обновляются с задержкой
 * 
 * @example
 * ```tsx
 * <DeferredSearch>
 *   {(query) => <SearchResults query={query} />}
 * </DeferredSearch>
 * ```
 */
export function DeferredSearch({
  initialQuery = '',
  children,
  loadingFallback = <CircularProgress size={20} />,
  debounceMs = 300,
  ...textFieldProps
}: DeferredSearchProps) {
  const [query, setQuery] = useState(initialQuery)
  const deferredQuery = useDeferredValue(query)
  const isStale = query !== deferredQuery

  return (
    <Box>
      <TextField
        {...textFieldProps}
        value={query}
        onChange={(e) => setQuery(e.target.value)}
        slotProps={{
          input: {
            endAdornment: isStale ? loadingFallback : textFieldProps.slotProps?.input?.endAdornment
          }
        }}
      />
      
      <Box
        sx={{
          mt: 2,
          opacity: isStale ? 0.5 : 1,
          transition: 'opacity 0.2s'
        }}
      >
        <Suspense fallback={loadingFallback}>
          {children(deferredQuery)}
        </Suspense>
      </Box>
    </Box>
  )
}

export default DeferredSearch

