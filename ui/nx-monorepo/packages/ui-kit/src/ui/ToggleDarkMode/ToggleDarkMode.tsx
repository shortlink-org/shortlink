import { useTheme } from 'next-themes'
import React, { useState, useEffect, useContext } from 'react'

import { ColorModeContext } from '../../theme/ColorModeContext'

// @ts-ignore
import './styles.css'

type ToggleDarkModeProps = {
  id: string
  onChange?: (theme: string) => void
}

export const ToggleDarkMode: React.FC<ToggleDarkModeProps> = ({
  id,
  onChange,
}) => {
  // @ts-ignore
  const { setTheme, resolvedTheme } = useTheme()
  const [mounted, setMounted] = useState(false)

  useEffect(() => {
    setMounted(true)
  }, [])

  // @ts-ignore
  const { darkMode, setDarkMode } = useContext(ColorModeContext)

  // @ts-ignore
  const onChangeTheme = () => {
    setDarkMode(!darkMode)
    const newTheme = darkMode ? 'light' : 'dark'
    setTheme(newTheme)
    onChange?.(newTheme)
  }

  if (!mounted) return null

  const labelText = darkMode ? 'Switch to light mode' : 'Switch to dark mode'

  return (
    <div id={id} className="toggleWrapper">
      <input
        type="checkbox"
        className="dn"
        id="dn"
        onChange={onChangeTheme}
        checked={!darkMode}
        aria-label={labelText}
      />
      <label htmlFor="dn" className="toggle">
        <span className="toggle__handler">
          <span className="crater crater--1" />
          <span className="crater crater--2" />
          <span className="crater crater--3" />
        </span>
        <span className="star star--1" />
        <span className="star star--2" />
        <span className="star star--3" />
        <span className="star star--4" />
        <span className="star star--5" />
        <span className="star star--6" />
      </label>
    </div>
  )
}

export default ToggleDarkMode
