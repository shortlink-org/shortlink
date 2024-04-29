'use client'

// @ts-ignore
import { ToggleDarkMode, SearchForm, Sidebar } from '@shortlink-org/ui-kit'
import { useEffect, useState, Fragment } from 'react'
import MuiAppBar, { AppBarProps as MuiAppBarProps } from '@mui/material/AppBar'
import Button from '@mui/material/Button'
import Divider from '@mui/material/Divider'
import MuiDrawer from '@mui/material/Drawer'
import IconButton from '@mui/material/IconButton'
import { useTheme as useNextTheme } from 'next-themes'
import { styled, useTheme, Theme, CSSObject } from '@mui/material/styles'
import Toolbar from '@mui/material/Toolbar'
import Typography from '@mui/material/Typography'
import { AxiosError } from 'axios'
import Link from 'next/link'

// Importing icons
import ChevronLeftIcon from '@mui/icons-material/ChevronLeft'
import ChevronRightIcon from '@mui/icons-material/ChevronRight'
import MenuIcon from '@mui/icons-material/Menu'
import ory from 'pkg/sdk'

import Notification from './notification'
import Profile from './profile'
import secondMenu from './secondMenu'

const drawerWidth = 290

interface AppBarProps extends MuiAppBarProps {
  open?: boolean
}

const Drawer = styled(MuiDrawer, {
  shouldForwardProp: (prop) => prop !== 'open',
})(({ theme, open }) => ({
  width: drawerWidth,
  flexShrink: 0,
  whiteSpace: 'nowrap',
  boxSizing: 'border-box',
  ...(open && {
    ...openedMixin(theme),
    '& .MuiDrawer-paper': openedMixin(theme),
  }),
  ...(!open && {
    ...closedMixin(theme),
    '& .MuiDrawer-paper': closedMixin(theme),
  }),
}))

const DrawerHeader = styled('div')(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'flex-end',
  padding: theme.spacing(0, 1),
  // necessary for content to be below app bar
  ...theme.mixins.toolbar,
}))

const openedMixin = (theme: Theme): CSSObject => ({
  width: drawerWidth,
  transition: theme.transitions.create('width', {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.enteringScreen,
  }),
  overflowX: 'hidden',
})

const closedMixin = (theme: Theme): CSSObject => ({
  transition: theme.transitions.create('width', {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.leavingScreen,
  }),
  overflowX: 'hidden',
  width: `calc(${theme.spacing(7)} + 1px)`,
  [theme.breakpoints.up('sm')]: {
    width: `calc(${theme.spacing(8)} + 1px)`,
  },
})

const AppBar = styled(MuiAppBar, {
  shouldForwardProp: (prop) => prop !== 'open',
})<AppBarProps>(({ theme, open }) => ({
  zIndex: theme.zIndex.drawer + 1,
  transition: theme.transitions.create(['width', 'margin'], {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.leavingScreen,
  }),
  ...(open && {
    marginLeft: drawerWidth,
    width: `calc(100% - ${drawerWidth}px)`,
    transition: theme.transitions.create(['width', 'margin'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  }),
}))

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
    setOpen(true)
  }

  const handleDrawerClose = () => {
    setOpen(false)
  }

  const onChangeTheme = (theme: string) => {
    setTheme(theme)
  }

  return (
    <Fragment>
      <AppBar key="appbar" position="fixed" open={open}>
        <Toolbar>
          <IconButton
            color="inherit"
            aria-label="menu"
            onClick={handleDrawerOpen}
            edge="start"
            sx={{
              marginRight: 5,
              ...(open && { display: 'none' }),
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
      </AppBar>

      <Fragment key="menu">
        {hasSession && (
          <Drawer key="drawer" variant="permanent" open={open}>
            <DrawerHeader>
              <IconButton onClick={handleDrawerClose}>{theme.direction === 'rtl' ? <ChevronRightIcon /> : <ChevronLeftIcon />}</IconButton>
            </DrawerHeader>
            <Divider flexItem />

            <Sidebar />
          </Drawer>
        )}
      </Fragment>
    </Fragment>
  )
}
