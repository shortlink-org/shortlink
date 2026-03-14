'use client'

// @ts-ignore
import { AppHeader } from '@shortlink-org/ui-kit'
import Link from 'next/link'
import { usePathname } from 'next/navigation'
import { ThemeToggle } from '@/components/ThemeToggle'
import Notification from './notification'
import Profile from './profile'
import TravelExploreIcon from '@mui/icons-material/TravelExplore'
import MailOutlineIcon from '@mui/icons-material/MailOutline'
import BarChartIcon from '@mui/icons-material/BarChart'

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
      notifications={{
        render: () => <Notification />,
      }}
      showProfile={hasSession}
      profile={{
        render: () => <Profile />,
      }}
      showLogin={!hasSession && !isSessionLoading}
      loginButton={{
        href: '/auth/login',
        label: 'Log in',
      }}
    />
  )
}
