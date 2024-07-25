import { useTheme } from 'next-themes'
import React, { useEffect } from 'react'
import { useColorScheme } from '@mui/material/styles'

import './styles.css'

type ToggleDarkModeProps = {
  id: string
}

export const ToggleDarkMode: React.FC<ToggleDarkModeProps> = ({
  id,
}) => {
  const { mode, setMode } = useColorScheme()
  const { setTheme } = useTheme()
  const [mounted, setMounted] = React.useState(false)

  useEffect(() => {
    setMounted(true)
  }, [])

  if (!mounted) {
    // for server-side rendering
    // learn more at https://github.com/pacocoursey/next-themes#avoid-hydration-mismatch
    return null
  }

  const onToggle = () => {
    setMode(mode === 'light' ? 'dark' : 'light')
    setTheme(mode === 'light' ? 'dark' : 'light')
  }

  const labelText =
    mode === 'dark' ? 'Switch to light mode' : 'Switch to dark mode'

  return (
    <div id={id} className="toggleWrapper">
      <input
        type="checkbox"
        className="dn"
        id="dn"
        onChange={onToggle}
        checked={mode === 'light'}
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
