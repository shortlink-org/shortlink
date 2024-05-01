'use client'

// @ts-ignore
import { ToggleDarkMode, SearchForm, Sidebar } from '@shortlink-org/ui-kit'
import { useEffect, useState, Fragment } from 'react'
import MuiAppBar from '@mui/material/AppBar'
import Button from '@mui/material/Button'
import IconButton from '@mui/material/IconButton'
import { useTheme as useNextTheme } from 'next-themes'
import { useTheme } from '@mui/material/styles'
import Toolbar from '@mui/material/Toolbar'
import Typography from '@mui/material/Typography'
import { AxiosError } from 'axios'
import Link from 'next/link'

// Importing icons
import MenuIcon from '@mui/icons-material/Menu'
import ory from 'pkg/sdk'

import Notification from './notification'
import Profile from './profile'
import secondMenu from './secondMenu'

export default function Header() {
  const [session, setSession] = useState<string>('No valid Ory Session was found.\nPlease sign in to receive one.')
  const [hasSession, setHasSession] = useState<boolean>(false)

  const { setTheme } = useNextTheme()

  useEffect(() => {
    ory
      .toSession()
      .then(({ data }) => {
        setSession(JSON.stringify(data, null, 2))
        setHasSession(true)
      })
      .catch((err: AxiosError) =>
        // Something else happened!
        Promise.reject(err),
      )
  }, [])

  const theme = useTheme()
  const [open, setOpen] = useState(false)

  const handleDrawerOpen = () => {
    setOpen(!open)
  }

  const onChangeTheme = (theme: string) => {
    setTheme(theme)
  }

  return (
    <div className={'flex-initial h-full'}>
      <MuiAppBar position="fixed" open={open}>
        <Toolbar>
          <IconButton
            color="inherit"
            aria-label="menu"
            onClick={handleDrawerOpen}
            edge="start"
            sx={{
              marginRight: 5,
            }}
            disabled={!hasSession}
          >
            <MenuIcon />
          </IconButton>

          <Button href="/" component={Link} color="secondary" sx={{ flexGrow: 1, display: { xs: 'none', sm: 'block' } }}>
            <Typography component="h1" variant="h6" color="inherit" noWrap>
              Shortlink
            </Typography>
          </Button>

          <ToggleDarkMode id="toggleDarkMode" onChange={onChangeTheme} />

          {secondMenu()}

          <SearchForm />

          {hasSession ? (
            <Fragment>
              <Profile />

              <Notification />
            </Fragment>
          ) : (
            <Button component={Link} href="/auth/login" type="submit" variant="outlined" color="secondary">
              Log in
            </Button>
          )}
        </Toolbar>
      </MuiAppBar>

      {hasSession && <Sidebar mode={open ? 'full' : 'mini'} />}
    </div>
  )
}
