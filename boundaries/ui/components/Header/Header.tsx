'use client'

// @ts-ignore
import { ToggleDarkMode, SearchForm } from '@shortlink-org/ui-kit'
import Button from '@mui/material/Button'
import IconButton from '@mui/material/IconButton'
import Typography from '@mui/material/Typography'
import Link from 'next/link'
import MenuIcon from '@mui/icons-material/Menu'
import Notification from './notification'
import Profile from './profile'
import secondMenu from './secondMenu'
import React from 'react'

export default function Header({ hasSession, setOpen }) {
  return (
    <nav className="bg-indigo-600 dark:bg-slate-800 text-white grid grid-cols-[auto_1fr_auto] items-center p-2 shadow-md z-50">
      {/* Left: Menu + Brand */}
      <div className="flex items-center gap-3">
        <IconButton
          color="inherit"
          aria-label="menu"
          onClick={setOpen}
          edge="start"
          disabled={!hasSession}
          className="hover:bg-indigo-500 dark:hover:bg-slate-700 transition-colors"
        >
          <MenuIcon />
        </IconButton>

        <Link href="/" passHref>
          <Typography
            component="h1"
            variant="h6"
            className="font-bold tracking-wide text-white hover:text-gray-100 transition-colors"
            noWrap
          >
            Shortlink
          </Typography>
        </Link>
      </div>

      {/* Center: Spacer (or could add nav links) */}
      <div />

      {/* Right: Controls */}
      <div className="flex items-center gap-3">
        <ToggleDarkMode id="ToggleDarkMode" />

        {secondMenu()}

        <SearchForm />

        {hasSession ? (
          <div className="flex items-center gap-3">
            <Profile />
            <Notification />
          </div>
        ) : (
          <Button
            component={Link}
            href="/auth/login"
            variant="outlined"
            sx={{
              color: 'white',
              borderColor: 'white',
              textTransform: 'none',
              '&:hover': {
                backgroundColor: 'rgba(255,255,255,0.1)',
                borderColor: 'white',
              },
            }}
          >
            Log in
          </Button>
        )}
      </div>
    </nav>
  )
}
