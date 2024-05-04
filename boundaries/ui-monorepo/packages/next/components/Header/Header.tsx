'use client'

// @ts-ignore
import { ToggleDarkMode } from '@shortlink-org/ui-kit'
import SearchForm from '@shortlink-org/ui-kit/src/ui/SearchForm/SearchForm'
import Button from '@mui/material/Button'
import IconButton from '@mui/material/IconButton'
import { useTheme as useNextTheme } from 'next-themes'
import Typography from '@mui/material/Typography'
import Link from 'next/link'

// Importing icons
import MenuIcon from '@mui/icons-material/Menu'

import Notification from './notification'
import Profile from './profile'
import secondMenu from './secondMenu'

// @ts-ignore
export default function Header({ hasSession, setOpen }) {
  const { setTheme } = useNextTheme()

  const onChangeTheme = (theme: string) => {
    setTheme(theme)
  }

  return (
    <nav className={'bg-indigo-500 text-white grid grid-cols-[auto_1fr_auto] z-100 p-2 justify-center items-center dark:bg-slate-800'}>
      <div className={'flex flex-row mx-2'}>
        <IconButton color="inherit" aria-label="menu" onClick={setOpen} edge="start" disabled={!hasSession}>
          <MenuIcon />
        </IconButton>

        <Button href="/" component={Link} color="secondary">
          <Typography className={'mx-5'} component="h1" variant="h6" color="inherit" noWrap>
            Shortlink
          </Typography>
        </Button>
      </div>

      <div />

      <div className={'flex flex-row justify-center items-center'}>
        <ToggleDarkMode id="toggleDarkMode" onChange={onChangeTheme} />

        {secondMenu()}

        <SearchForm />

        {hasSession ? (
          <div className={'flex flex-row mx-2'}>
            <Profile />

            <Notification />
          </div>
        ) : (
          <Button component={Link} href="/auth/login" type="submit" variant="outlined" color="secondary">
            Log in
          </Button>
        )}
      </div>
    </nav>
  )
}
