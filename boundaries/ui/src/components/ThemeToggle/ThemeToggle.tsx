'use client'

import { useTheme } from 'next-themes'
import { useEffect, useState, useCallback } from 'react'
// @ts-ignore
import { ToggleDarkMode } from '@shortlink-org/ui-kit'

export function ThemeToggle() {
  const { resolvedTheme, setTheme } = useTheme()
  const [mounted, setMounted] = useState(false)

  // Avoid hydration mismatch
  useEffect(() => {
    setMounted(true)
  }, [])

  const isDark = resolvedTheme === 'dark'

  const handleToggle = useCallback(() => {
    const newTheme = isDark ? 'light' : 'dark'
    setTheme(newTheme)
  }, [isDark, setTheme])

  const handleWrapperClick = useCallback(
    (e: React.MouseEvent) => {
      e.stopPropagation()
      handleToggle()
    },
    [handleToggle],
  )

  if (!mounted) {
    return null
  }

  // Use stable ID
  const toggleId = 'ThemeToggle'

  return (
    <div
      style={{
        position: 'relative',
        display: 'inline-block',
        zIndex: 10,
        cursor: 'pointer',
      }}
      onClick={handleWrapperClick}
    >
      <ToggleDarkMode
        key={`${toggleId}-${resolvedTheme}`} // Force re-render when theme changes
        id={toggleId}
        checked={isDark}
        onClick={handleToggle}
        ariaLabel={isDark ? 'Switch to light mode' : 'Switch to dark mode'}
      />
    </div>
  )
}
