'use client'

// @ts-ignore
import { AppHeader } from '@shortlink-org/ui-kit'
import { AxiosError } from 'axios'
import Link from 'next/link'
import { usePathname } from 'next/navigation'
import { useCallback, useEffect, useMemo, useState } from 'react'
import { ThemeToggle } from '@/components/ThemeToggle'
import TravelExploreIcon from '@mui/icons-material/TravelExplore'
import MailOutlineIcon from '@mui/icons-material/MailOutline'
import BarChartIcon from '@mui/icons-material/BarChart'
import { useSession } from '@/contexts/SessionContext'
import ory from '@/pkg/sdk'

interface HeaderProps {
  hasSession: boolean
  isSessionLoading?: boolean
  setOpen: () => void
}

const secondMenuItems = [
  {
    name: 'Pricing',
    description: 'Measure actions your users take',
    href: '/pricing',
    icon: TravelExploreIcon,
  },
  {
    name: 'Contacts',
    description: 'Send us an email',
    href: '/contact',
    icon: MailOutlineIcon,
  },
  {
    name: 'Reports',
    description: 'Keep track of your growth',
    href: '/user/reports',
    icon: BarChartIcon,
  },
]

export default function Header({ hasSession, isSessionLoading = false, setOpen }: HeaderProps) {
  const pathname = usePathname()
  const [logoutToken, setLogoutToken] = useState('')
  const { session } = useSession()

  useEffect(() => {
    if (!hasSession) return

    ory
      .createBrowserLogoutFlow()
      .then(({ data }) => {
        setLogoutToken(data.logout_token)
      })
      .catch((err: AxiosError) => {
        if (err.response?.status === 401) return
        return Promise.reject(err)
      })
  }, [hasSession])

  const traits = (session?.identity?.traits as Record<string, any> | undefined) ?? {}
  const firstName = traits?.name?.first ?? ''
  const lastName = traits?.name?.last ?? ''
  const email = traits?.email ?? ''
  const displayName = `${firstName} ${lastName}`.trim() || email || 'User'
  const avatar = traits?.avatar_url

  const notifications = useMemo(
    () => ({
      count: 4,
      items: [
        {
          id: 'sara-reply',
          title: 'Sara Salah',
          message: 'Replied on the Upload Image article.',
          time: '2m',
          avatar:
            'https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=334&q=80',
        },
        {
          id: 'slick-follow',
          title: 'Slick Net',
          message: 'Started following you.',
          time: '45m',
          avatar:
            'https://images.unsplash.com/photo-1531427186611-ecfd6d936c79?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=634&q=80',
        },
        {
          id: 'jane-like',
          title: 'Jane Doe',
          message: 'Liked your reply on Test with TDD.',
          time: '1h',
          avatar:
            'https://images.unsplash.com/photo-1450297350677-623de575f31c?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=334&q=80',
        },
        {
          id: 'abigail-follow',
          title: 'Abigail Bennett',
          message: 'Started following you.',
          time: '3h',
          avatar:
            'https://images.unsplash.com/photo-1580489944761-15a19d654956?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=398&q=80',
        },
      ],
    }),
    [],
  )

  const handleLogout = useCallback(async () => {
    try {
      const token = logoutToken || (await ory.createBrowserLogoutFlow()).data.logout_token
      await ory.updateLogoutFlow({ token })
      window.location.assign('/auth/login')
    } catch (err) {
      console.error('Logout failed', err)
    }
  }, [logoutToken])

  const profile = useMemo(
    () => ({
      avatar,
      name: displayName,
      email,
      menuItems: [
        {
          name: 'Your Profile',
          href: '/profile',
        },
        {
          name: 'Sign out',
          onClick: handleLogout,
          confirmDialog: {
            title: 'Sign out?',
            description: 'You will be redirected to the login page.',
            confirmText: 'Sign out',
            variant: 'danger',
          },
        },
      ],
    }),
    [avatar, displayName, email, handleLogout],
  )

  return (
    <AppHeader
      currentPath={pathname}
      LinkComponent={Link}
      showMenuButton={true}
      onMenuClick={setOpen}
      menuButtonDisabled={!hasSession}
      showThemeToggle={true}
      themeToggleComponent={<ThemeToggle />}
      showSecondMenu={true}
      secondMenuLabel="Solutions"
      secondMenuItems={secondMenuItems}
      showSearch={true}
      showNotifications={hasSession}
      notifications={notifications}
      showProfile={hasSession}
      profile={profile}
      showLogin={!hasSession && !isSessionLoading}
      loginButton={{
        href: '/auth/login',
        label: 'Log in',
      }}
    />
  )
}
