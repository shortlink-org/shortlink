import ChevronLeftIcon from '@mui/icons-material/ChevronLeft'
import ChevronRightIcon from '@mui/icons-material/ChevronRight'
import MenuIcon from '@mui/icons-material/Menu'
import MuiAppBar, { AppBarProps as MuiAppBarProps } from '@mui/material/AppBar'
import Button from '@mui/material/Button'
import Divider from '@mui/material/Divider'
import MuiDrawer from '@mui/material/Drawer'
import IconButton from '@mui/material/IconButton'
import List from '@mui/material/List'
import { styled, useTheme, Theme, CSSObject } from '@mui/material/styles'
import Toolbar from '@mui/material/Toolbar'
import Typography from '@mui/material/Typography'
// @ts-ignore
import { ToggleDarkMode } from '@shortlink-org/ui-kit'
import { AxiosError } from 'axios'
import Link from 'next/link'
import { useTheme as useNextTheme } from 'next-themes'
import { useEffect, useState, Fragment } from 'react'

import { mainListItems, secondaryListItems, adminListItems } from './listItems'
import Notification from './notification'
import Profile from './profile'
// @ts-ignore
import secondMenu from './secondMenu'
import ory from '../../pkg/sdk'
import SearchForm from '../SearchForm'

const drawerWidth = 290

interface AppBarProps extends MuiAppBarProps {
  open?: boolean
}

/* eslint-disable */
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

const Header = () => {
  const [session, setSession] = useState<string>(
    'No valid Ory Session was found.\nPlease sign in to receive one.',
  )
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

  return [
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

        <Link href="/" legacyBehavior>
          <Button
            href="/"
            color="inherit"
            sx={{ flexGrow: 1, display: { xs: 'none', sm: 'block' } }}
          >
            <Typography component="h1" variant="h6" color="inherit" noWrap>
              Shortlink
            </Typography>
          </Button>
        </Link>

        <ToggleDarkMode onChange={onChangeTheme} />

        {secondMenu()}

        <SearchForm />

        {hasSession ? (
          <Fragment>
            <Profile />

            <Notification />
          </Fragment>
        ) : (
          <Link href="/auth/login" legacyBehavior>
            <Button type="submit" variant="outlined" color="inherit">
              Log in
            </Button>
          </Link>
        )}
      </Toolbar>
    </AppBar>,
    <Fragment>
      {hasSession && (
        <Drawer key="drawer" variant="permanent" open={open}>
          <DrawerHeader>
            <IconButton onClick={handleDrawerClose}>
              {theme.direction === 'rtl' ? (
                <ChevronRightIcon />
              ) : (
                <ChevronLeftIcon />
              )}
            </IconButton>
          </DrawerHeader>
          <Divider />

          <List>{mainListItems}</List>
          <Divider />

          <List>{secondaryListItems}</List>
          <Divider />

          <List className={'h-full'}>{adminListItems}</List>

          <div className="flex bg-blue-800 justify-start space-x-2 items-center py-4 px-3.5">
            <div>
              <img
                src="https://i.ibb.co/fxrbS6p/Ellipse-2-2.png"
                alt="avatar"
              />
            </div>
            {open && (
              <Fragment>
                <div className="flex flex-col justify-start items-start space-y-2">
                  <p className="cursor-pointer text-base leading-4 text-white">
                    Alexis Enache
                  </p>
                  <p className="cursor-pointer text-xs leading-3 text-gray-200">
                    alexis _enache@gmail.com
                  </p>
                </div>
                <Link href="/user/profile" legacyBehavior>
                  <button
                    aria-label="visit"
                    className=" focus:ring-2 focus:outline-none hover:bg-blue-900 p-2.5 bg-blue-600 rounded-full"
                  >
                    <svg
                      width={20}
                      height={20}
                      viewBox="0 0 20 20"
                      fill="none"
                      xmlns="http://www.w3.org/2000/svg"
                    >
                      <path
                        d="M4.16666 10H15.8333"
                        stroke="white"
                        strokeWidth="1.5"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                      />
                      <path
                        d="M10.8333 15L15.8333 10"
                        stroke="white"
                        strokeWidth="1.5"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                      />
                      <path
                        d="M10.8333 5L15.8333 10"
                        stroke="white"
                        strokeWidth="1.5"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                      />
                    </svg>
                  </button>
                </Link>
              </Fragment>
            )}
          </div>
        </Drawer>
      )}
    </Fragment>,
  ]
}

export default Header
