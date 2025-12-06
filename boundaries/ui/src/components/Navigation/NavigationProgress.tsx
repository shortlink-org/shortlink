'use client'

import { useTransition, createContext, useContext, ReactNode } from 'react'
import { LinearProgress } from '@mui/material'

/**
 * Context для отслеживания глобального pending состояния навигации
 */
const NavigationContext = createContext<{
  isPending: boolean
  startTransition: (callback: () => void) => void
}>({
  isPending: false,
  startTransition: () => {}
})

export function useNavigation() {
  return useContext(NavigationContext)
}

interface NavigationProviderProps {
  children: ReactNode
}

/**
 * Provider для глобального управления transitions навигации
 * 
 * Показывает progress bar вверху страницы при навигации
 * 
 * @example
 * ```tsx
 * <NavigationProvider>
 *   <Layout>
 *     {children}
 *   </Layout>
 * </NavigationProvider>
 * ```
 */
export function NavigationProvider({ children }: NavigationProviderProps) {
  const [isPending, startTransition] = useTransition()

  return (
    <NavigationContext.Provider value={{ isPending, startTransition }}>
      {/* Progress bar вверху страницы */}
      {isPending && (
        <LinearProgress 
          sx={{ 
            position: 'fixed', 
            top: 0, 
            left: 0, 
            right: 0, 
            zIndex: 9999 
          }} 
        />
      )}
      
      {/* Весь контент с небольшой прозрачностью при pending */}
      <div style={{ 
        opacity: isPending ? 0.7 : 1,
        transition: 'opacity 0.2s ease-in-out'
      }}>
        {children}
      </div>
    </NavigationContext.Provider>
  )
}

export default NavigationProvider

