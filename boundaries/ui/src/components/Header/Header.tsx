'use client'

// @ts-ignore
import { SearchForm } from '@shortlink-org/ui-kit'
import { Button } from '@mui/material'
import IconButton from '@mui/material/IconButton'
import Typography from '@mui/material/Typography'
import Link from 'next/link'
import { TransitionLink } from '@/components/Navigation'
// @ts-ignore
import { ToggleDarkMode } from '@shortlink-org/ui-kit'
import MenuIcon from '@mui/icons-material/Menu'
import Notification from './notification'
import Profile from './profile'
import secondMenu from './secondMenu'
import { useEffect, useState } from 'react'

interface HeaderProps {
  hasSession: boolean
  setOpen: () => void
}

export default function Header({ hasSession, setOpen }: HeaderProps) {
  const [origin, setOrigin] = useState('')

  useEffect(() => {
    setOrigin(window.location.origin)
  }, [])

  return (
    <nav className="bg-gradient-to-r from-indigo-600 via-purple-600 to-indigo-700 dark:from-indigo-800 dark:via-purple-800 dark:to-indigo-900 text-white shadow-lg border-b border-indigo-500/20 dark:border-purple-500/30 backdrop-blur-sm z-50">
      <div className="max-w-7xl mx-auto px-3 sm:px-4 lg:px-8">
        <div className="flex items-center justify-between h-14 sm:h-16">
          {/* Left: Menu + Brand */}
          <div className="flex items-center gap-2 sm:gap-4">
            <IconButton
              color="inherit"
              aria-label="Toggle navigation menu"
              onClick={setOpen}
              edge="start"
              disabled={!hasSession}
              className="hover:bg-white/10 active:bg-white/20 transition-all duration-200 rounded-lg p-2"
              sx={{
                '&:hover': {
                  transform: 'scale(1.05)',
                },
                '&:active': {
                  transform: 'scale(0.95)',
                },
              }}
            >
              <MenuIcon className="text-lg" />
            </IconButton>

            <TransitionLink href="/" className="group">
              <div className="flex items-center gap-1.5 sm:gap-2">
                <div className="w-7 h-7 sm:w-8 sm:h-8 bg-indigo-900/30 dark:bg-indigo-400/20 rounded-lg flex items-center justify-center group-hover:bg-indigo-900/40 dark:group-hover:bg-indigo-400/30 transition-all duration-200">
                  <span className="text-white font-bold text-xs sm:text-sm">S</span>
                </div>
                <Typography
                  component="h1"
                  variant="h6"
                  className="font-bold tracking-wide text-white group-hover:text-gray-100 transition-all duration-200 text-base sm:text-lg"
                  noWrap
                >
                  Shortlink
                </Typography>
              </div>
            </TransitionLink>
          </div>

          {/* Center: Spacer (or could add nav links) */}
          <div className="hidden md:flex items-center space-x-8">
            {/* Add navigation links here if needed */}
          </div>

          {/* Right: Controls */}
          <div className="flex items-center gap-1.5 sm:gap-3">
            <div className="hidden sm:block relative">
              <ToggleDarkMode id="ToggleDarkMode" />
            </div>

            <div className="hidden sm:block">
              {secondMenu()}
            </div>

            <div className="hidden md:block">
              <SearchForm />
            </div>

            {hasSession ? (
              <div className="flex items-center gap-3">
                <div className="hidden sm:block">
                  <Notification />
                </div>
                <div className="relative">
                  <Profile />
                </div>
              </div>
            ) : (
              <Button
                component={Link}
                href={`${origin}/auth/login`}
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
        </div>
      </div>
    </nav>
  )
}
